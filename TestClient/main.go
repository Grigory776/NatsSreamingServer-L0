package main

import (
	"io/ioutil"
	"log"

	stan "github.com/nats-io/stan.go"
)


func main(){
	sc, err := stan.Connect("test-cluster", "client-sender", stan.NatsURL("127.0.0.1:4223"))
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()
	inf, err := ioutil.ReadFile("model2.json")
	if err != nil {
		log.Println(err)
	}
	sc.Publish("foo", inf) 
}
