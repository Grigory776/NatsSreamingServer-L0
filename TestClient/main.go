package main

import (
	"fmt"
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
	for i:=1; i < 8;i++{
		inf, err := ioutil.ReadFile("models/model" + fmt.Sprint(i) + ".json")
		if err != nil {
			log.Println(err)
		}
		sc.Publish("foo", inf) 
	}
}
