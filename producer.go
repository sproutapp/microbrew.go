package microbrew

import (
  "log"
  "github.com/streadway/amqp"
  "encoding/json"
)


type Producer struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
  exchange   string
}

type MicrobrewProducer interface {
  Publish(routingKey string, payload *Payload) error
}

type Payload struct {
  Event string      `json:"event"`
  Data interface{}  `json:"data"`
}

func (p *Producer) Publish(routingKey string, payload *Payload) error {
  marshalled, _ := json.Marshal(payload)

  err := p.Channel.Publish(
    p.exchange,     // exchange
    routingKey,     // routing key
    false,          // mandatory
    false,          // immediate
    amqp.Publishing{
      Headers:         amqp.Table{},
        ContentType:     "text/plain",
        ContentEncoding: "",
        Body:            marshalled,
        DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
        Priority:        0,              // 0-9
    },
  )

  if err == nil {
    log.Printf("Produced \"sensor::received\" with payload \"%s\" to exchange \"%s\", with routing key \"%s\"", payload, p.exchange, routingKey)
  }

  return err
}

func NewProducer(uri, exchange, exchangeType string) *Producer {
  var err error
  var conn *amqp.Connection
  var channel *amqp.Channel

  conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
  FailOnError(err, "Failed to connect to RabbitMQ")

  channel, err = conn.Channel()
	FailOnError(err, "Failed to open a Channel")

  err = channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
  FailOnError(err, "Failed to setup an exchange")

  return &Producer { conn, channel, exchange }
}
