package microbrew

import (
  "encoding/json"
)

type Agent struct {
	producer *Producer
}

type Payload struct {
  Event string      `json:"event"`
  Data interface{}  `json:"data"`
}

func NewAgent(uri, exchange, exchangeType string) (*Agent, error) {
  var err error
  a := &Agent{
    producer: nil,
  }

  a.producer, err = NewProducer(uri, exchange, exchangeType)
  if err != nil {
    return nil, err
  }

  return a, nil
}

func (a Agent) Signal(event string, data interface{}) error {
  payload := &Payload{
    Event: event,
    Data: data,
  }
  marshalled, _ := json.Marshal(payload)

  err := Publish(*a.producer, "", marshalled)

  return err
}
