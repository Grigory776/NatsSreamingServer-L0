package main

import (
	"NatsStreamingServer/MainServer/nats"
)

const url = "postgres://gorge:Liza110909.@localhost:5432/gorge"

func main() {
	ex := nats.Nats {
		ClusterID: "test-cluster",
		ClientId: "boss",
		NameChanel: "foo",
		UrlBD: url,
	}
	ex.Start()
}
