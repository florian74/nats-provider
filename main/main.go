package main

import (
	"flag"
	"github.com/florian74/nats-provider/provider"
	"github.com/nats-io/nats.go"
)

func main() {
	url := flag.String("url", nats.DefaultURL, "url")
	queueName := flag.String("queue", "default", "queue target name")
	topic := flag.String("topic", "default", "topic target name")
	providerName := flag.String("name", "provider", "provider name")
	reply := flag.String("reply", "", "provider reply topic")
	jetStream := flag.Bool("stream", false, "use JetStream")
	flag.Parse()

	var w provider.Provider

	w = &provider.HelloProvider{}
	w.Connect(*url, *queueName, *topic, *providerName, *reply, *jetStream)
}
