package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	nats "github.com/nats-io/go-nats"
)

func main() {
	URL := nats.DefaultURL
	subs, _ := nats.Connect(URL)
	go func() {
		subs.Subscribe("pubsub", func(m *nats.Msg) {
			fmt.Println("partner: ", string(m.Data))
		})
	}()

	nc, err := nats.Connect(URL)
	log.Println("Connecting to", URL)
	if err != nil {
		log.Fatalln(err)
	}

	if nc.IsConnected() {
		log.Println("Success!")
	}

	var message string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message = scanner.Text()
		}
		if err = nc.Publish("subpub", []byte(message)); err != nil {
			log.Fatalln(err)
		}
	}

	nc.Close()
	subs.Close()
}
