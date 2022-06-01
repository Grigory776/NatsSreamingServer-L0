package main

import (
	"NatsStreamingServer/MainServer/nats"
)

const url = "postgres://gorge:..."

func main() {
	ex := nats.Nats {
		ClusterID: "test-cluster",
		ClientId: "boss",
		NameChanel: "foo",
		UrlBD: url,
	}
	ex.Start()
}
