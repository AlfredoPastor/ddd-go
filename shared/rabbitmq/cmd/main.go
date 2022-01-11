package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/AlfredoPastor/ddd-go/shared/rabbitmq"
)

type Config struct {
	*rabbitmq.AMQPEnv
}

func (c Config) GetAMQPEnv() *rabbitmq.AMQPEnv {
	return c.AMQPEnv
}

func main() {
	// now := time.Now()
	var item string
	flag.StringVar(&item, "mode", "pub", "mode publisher")
	flag.Parse()
	cfg := Config{
		AMQPEnv: &rabbitmq.AMQPEnv{
			User:     "guest",
			Password: "guest",
			URL:      "rabbitdev.alfredo.dev",
		},
	}
	cli := rabbitmq.NewRabbitClient(cfg)
	ctx := context.Background()
	err := cli.Run(ctx)
	if err != nil {
		panic(err)
	}
	ss := ExampleHanler{}
	da := make(chan bool)
	if item == "sub" {
		go func() {
			cli.Subscriber(ctx, "events.place.verified", ss.Handle)
		}()
		go func() {
			cli.Subscriber(ctx, "events.place.created", ss.Handle)
		}()
		<-da
	}

	for {
		go func() {
			ff := time.Now().String()
			err = cli.Publisher("events.place.verified", []byte(fmt.Sprintf("Helooooo %s", ff)))
			if err != nil {
				fmt.Println(err.Error())
			}
			err = cli.Publisher("events.place.created", []byte(fmt.Sprintf("Chaoooo %s", ff)))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Sent message ", ff)
		}()
		time.Sleep(2 * time.Second)
	}
	// fmt.Println("time: ", time.Since(now))
}

type ExampleHanler struct {
}

func (e ExampleHanler) Handle(message []byte) bool {
	fmt.Println("Received a message: ", string(message))
	return true
}
