package config

type GlobalConfig struct {
	DB   map[string]DBConfig `yaml:"db"`
	NSQ  NSQConfig           `yaml:"nsq"`
	HTTP HTTPConfig          `yaml:"http"`
}

type HTTPConfig struct {
	Port int `json:"port"`
}

type NSQConfig struct {
	Address string `yaml:"address"`
}

type DBConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
