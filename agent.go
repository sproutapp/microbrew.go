package microbrew

type MicrobrewAgent interface {
  Signal(event string, data interface{}) error
}

type Agent struct {
  producer MicrobrewProducer
}

func NewAgentForProducer(producer MicrobrewProducer) *Agent {
  return &Agent { producer }
}

func NewAgent(uri, exchange, exchangeType string) *Agent {
  producer := NewProducer(uri, exchange, exchangeType)
  return NewAgentForProducer(producer)
}

func (a *Agent) Signal(event string, data interface{}) error {
  payload := &Payload{
    Event: event,
    Data: data,
  }

  return a.producer.Publish("", payload)
}
