package services

import "github.com/nsqio/go-nsq"

type ProductDependencies struct {
	Publisher *nsq.Producer
	PR        ProductResource
}
