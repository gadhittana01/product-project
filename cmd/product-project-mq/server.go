package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gihub.com/gadhittana01/product-project/config"
	nsqd "gihub.com/gadhittana01/product-project/handler/nsq"
	"github.com/nsqio/go-nsq"
)

type MQServerDependency struct {
	c       *config.GlobalConfig
	handler nsqd.IHandler
}

func serverConsumers(mq MQServerDependency) error {
	consumers := mq.handler.BuildConsumers()
	var qList []*nsq.Consumer

	shuttedDown := 0
	go func() {
		var (
			wg  sync.WaitGroup
			mux sync.Mutex
		)
		total := len(qList)
		wg.Add(total)
		for _, cons := range qList {
			go func(c *nsq.Consumer) { // use goroutines to  stop all of them ASAP
				defer wg.Done()
				c.Stop()
				select {
				case <-c.StopChan:
					mux.Lock()
					shuttedDown++
					mux.Unlock()
				case <-time.After(1 * time.Minute): // wait for at max 1 minute
				}
			}(cons)
		}
	}()

	go func() {
		for _, consumer := range consumers {
			//The only valid way to instantiate the Config
			config := nsq.NewConfig()
			//Tweak several common setup in config
			// Maximum number of times this consumer will attempt to process a message before giving up
			config.MaxAttempts = consumer.Config.MaxAttempts
			// Maximum number of messages to allow in flight
			config.MaxInFlight = consumer.Config.MaxInFlight
			// Maximum duration when REQueueing
			config.MaxRequeueDelay = consumer.Config.MaxRequeueDelay
			config.DefaultRequeueDelay = consumer.Config.DefaultRequeueDelay
			//Init topic name and channel
			topic := consumer.Topic
			channel := consumer.Channel

			//Creating the consumer
			q, err := nsq.NewConsumer(topic, channel, config)
			if err != nil {
				log.Fatal(err)
			}

			// Set the Handler for messages received by this Consumer.
			q.AddHandler(consumer.Handler)

			//Use nsqlookupd to find nsqd instances
			q.ConnectToNSQDs([]string{mq.c.NSQ.Address})

			// wait for signal to exit
			// Gracefully stop the consumer.
			qList = append(qList, q)
		}

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	return nil
}
