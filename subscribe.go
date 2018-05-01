package main

import (
	"fmt"
	"log"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	log.Println("Connecting to", nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		nc.Subscribe("foo", func(m *nats.Msg) {
			log.Println(string(m.Data))
		})
	}()

	fmt.Scanln()
	nc.Close()
}
