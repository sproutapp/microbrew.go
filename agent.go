package microbrew

import (
  "encoding/json"
)

type Agent struct {
	producer *Producer
}

type MicrobrewAgent interface {
	Init(uri, exchange, exchangeType string) error
}

type Payload struct {
  Event string      `json:"event"`
  Data interface{}  `json:"data"`
}

func (a *Agent) Init(uri, exchange, exchangeType string) error {
	producer := &Producer{}
  err := producer.Init(uri, exchange, exchangeType)
	if err != nil {
		return err
	}

	a.producer = producer

  return nil
}

func (a *Agent) Signal(event string, data interface{}) error {
  payload := &Payload{
    Event: event,
    Data: data,
  }
  marshalled, _ := json.Marshal(payload)

  return a.producer.Publish("", marshalled)
}
