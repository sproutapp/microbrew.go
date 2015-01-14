package microbrew

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
)

type ProducerMock struct {
  mock.Mock
}

func (p *ProducerMock) Publish(routingKey string, payload *Payload) error {
  args := p.Called(routingKey, payload)
  return args.Error(0)
}

func TestNewAgentForProducer(t *testing.T) {
  producer := new(ProducerMock)
  agent := NewAgentForProducer(producer)

  assert.Implements(t, (*MicrobrewAgent)(nil), agent)
}

func TestSignal(t *testing.T) {
  producer := new(ProducerMock)
  agent    := NewAgentForProducer(producer)
  payload  := &Payload{Event: "MySignal", Data: "somedata"}

  producer.On("Publish", "", payload).Return(nil)

  agent.Signal("MySignal", "somedata")

  producer.AssertExpectations(t)
}
