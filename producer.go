package microbrew

import (
  "log"
  "github.com/streadway/amqp"
)

type Producer struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
  exchange   string
}

type MicrobrewProducer interface {
  Init(uri, exchange, exchangeType string) error
  Publish(routingKey string, payload []byte) error
}

func (p *Producer) Publish(routingKey string, payload []byte) error {
  err := p.Channel.Publish(
    p.exchange,     // exchange
    routingKey,     // routing key
    false,          // mandatory
    false,          // immediate
    amqp.Publishing{
      Headers:         amqp.Table{},
        ContentType:     "text/plain",
        ContentEncoding: "",
        Body:            payload,
        DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
        Priority:        0,              // 0-9
    },
  )

  if err == nil {
    log.Printf("Produced \"sensor::received\" with payload \"%s\" to exchange \"%s\", with routing key \"%s\"", payload, p.exchange, routingKey)
  }

  return err
}

func (p *Producer) Init(uri, exchange, exchangeType string) error {
  var err error

  p.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
  FailOnError(err, "Failed to connect to RabbitMQ")

  p.Channel, err = p.Conn.Channel()
	FailOnError(err, "Failed to open a Channel")

  err = p.Channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
  FailOnError(err, "Failed to setup an exchange")

  return err
}
