package provider

import "github.com/nats-io/nats.go"

type Provider interface {
	Connect(url string, queueName string, subject string, providerName string, reply string, jetStream bool)
	Provide() []byte
	OnReply(msg *nats.Msg)
}
