package config

import (
	"github.com/AlfredoPastor/ddd-go/shared/ginhttp"
	"github.com/AlfredoPastor/ddd-go/shared/mongoatlas"
	"github.com/AlfredoPastor/ddd-go/shared/rabbitmq"
	env "github.com/caarlos0/env/v6"
)

type Config struct {
	*mongoatlas.MongoAtlasEnv
	*rabbitmq.AMQPEnv
	*ginhttp.HTTPEnv
}

func (c Config) GetMongoAtlasEnv() *mongoatlas.MongoAtlasEnv {
	return c.MongoAtlasEnv
}

func (c Config) GetAMQPEnv() *rabbitmq.AMQPEnv {
	return c.AMQPEnv
}

func (c Config) GetHTTPEnv() *ginhttp.HTTPEnv {
	return c.HTTPEnv
}

func NewConfig() (Config, error) {
	cfg := Config{
		AMQPEnv:       &rabbitmq.AMQPEnv{},
		MongoAtlasEnv: &mongoatlas.MongoAtlasEnv{},
		HTTPEnv:       &ginhttp.HTTPEnv{},
	}
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
