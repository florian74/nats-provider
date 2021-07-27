package provider

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type HelloProvider struct {
	natsConn *nats.Conn
	name     string
	subject  string
	queue    string
	reply    string
}

func (provider *HelloProvider) Provide() []byte {
	return []byte(fmt.Sprintf("Helloworld from %s at %s", provider.name, time.Now().String()))
}

func (provider *HelloProvider) OnReply(msg *nats.Msg) {
	fmt.Printf("reply %s received for provider %s at %s\n", msg.Data, provider.name, time.Now().String())
}

func (provider *HelloProvider) Connect(
	url string,
	queueName string,
	subject string,
	providerName string,
	reply string,
	jetStream bool,
) {
	var err error

	fmt.Printf("nats url is %s\n", url)

	provider.name = providerName
	provider.queue = queueName
	provider.subject = subject
	provider.reply = reply
	provider.natsConn, err = nats.Connect(url)

	if err != nil {
		panic("could not start nats client " + err.Error())
	}

	if reply != "" {
		sub, err := provider.natsConn.Subscribe(reply, provider.OnReply)
		if err != nil {
			fmt.Printf("Could not read reply")
		}
		defer sub.Unsubscribe()
	}

	if !jetStream {

		for {
			err = provider.natsConn.PublishRequest(provider.subject, provider.reply, provider.Provide())
			if err != nil {
				fmt.Printf("could not public with reply message %v", err)
			}
			time.Sleep(5 * time.Second)
		}

	} else {

		js, err := provider.natsConn.JetStream(nats.PublishAsyncMaxPending(256))
		if err != nil {
			fmt.Printf("could not connect  %v", err)
		}

		for {
			_, err := js.Publish(provider.subject, provider.Provide())
			if err != nil {
				fmt.Printf("could not publish %v\n", err)
			}
			time.Sleep(5 * time.Second)
		}
	}
}
