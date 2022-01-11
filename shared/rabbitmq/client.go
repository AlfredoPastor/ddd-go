package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type AMQPEnv struct {
	User            string `env:"RABBIT_USER" envDefault:"guest"`
	Password        string `env:"RABBIT_PASSWORD" envDefault:"guest"`
	URL             string `env:"RABBIT_URL" envDefault:"localhost"`
	Vhost           string `env:"RABBIT_VHOST" envDefault:""`
	CredentialsFile string `env:"RABBIT_CREDENTIALS_FILE" envDefault:""`
	Protocol        string `env:"RABBIT_PROTOCOL" envDefault:"amqp"`
}

type Config interface {
	GetAMQPEnv() *AMQPEnv
}

type RabbitClient struct {
	*Connection
	url           string
	exchangeName  string
	AppName       string
	PublisherMode bool
}

func (r *RabbitClient) Run(ctx context.Context) error {
	return r.connect()
}

func (r *RabbitClient) connect() error {
	conn, err := Dial(r.url)
	if err != nil {
		return err
	}
	log.Println("connected to rabbitmq")
	r.Connection = conn
	return nil
}

func (r *RabbitClient) openChannel() *Channel {
	for {
		ch, err := r.Connection.Channel()
		if err == nil {
			return ch
		}
		debug(err)
		time.Sleep(2 * time.Second)
	}
}

func (r *RabbitClient) Close() error {
	if r.PublisherMode {
		return nil
	}
	err := r.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitClient) Publisher(topic string, body []byte) error {
	ch := r.openChannel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		r.exchangeName,     // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		debug(err)
		return err
	}
	err = ch.Publish(
		r.exchangeName, // exchange
		topic,          // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		debug(err)
		return err
	}
	// r.logger.InfoEvents(topic, "PUB", body)
	return nil
}

func (r *RabbitClient) Subscriber(ctx context.Context, topic string, function func([]byte) bool) {
	ch := r.openChannel()
	err := ch.ExchangeDeclare(r.exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// Ejemplo de como hacer que los mensajes tengan TTL en la cola
	// amqp.Table{
	// "x-message-ttl": int32(5000),
	// "x-expires": int32(12000),
	// }
	_, err = ch.QueueDeclare(r.AppName+"."+topic, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	if err := ch.QueueBind(r.AppName+"."+topic, topic, r.exchangeName, false, nil); err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(r.AppName+"."+topic, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func(forever chan bool) {
		for d := range msgs {
			if ok := function(d.Body); ok {
				d.Ack(ok)
				continue
			}
			d.Ack(false)
		}
	}(forever)
	reconnect := <-forever
	ch.Close()
	if reconnect {
		go r.Subscriber(ctx, topic, function)
	}
}

func (r *RabbitClient) addCredentialsFromFile(env *AMQPEnv) error {
	if env.CredentialsFile == "" {
		return nil
	}
	data, err := ioutil.ReadFile(env.CredentialsFile)
	if err != nil {
		return errors.New("the credentials file doesn't exist")
	}

	tmp := AMQPEnv{}

	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return errors.New("the credentials file can't deserialize")
	}

	env.User = tmp.User
	env.Password = tmp.Password
	return nil
}

func NewRabbitClient(cfg Config) *RabbitClient {
	env := cfg.GetAMQPEnv()
	appName := os.Args[0]
	if strings.Contains(appName, "/tmp/go-build") {
		appName = "test"
	}
	cli := &RabbitClient{
		exchangeName: "test-exchange",
		AppName:      appName,
	}
	cli.addCredentialsFromFile(env)
	cli.url = fmt.Sprintf("%s://%s:%s@%s/%s", env.Protocol, env.User, env.Password, env.URL, env.Vhost)
	return cli
}
