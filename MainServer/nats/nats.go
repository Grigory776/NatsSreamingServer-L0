package nats

import (
	"NatsStreamingServer/MainServer/bd"
	"encoding/json"
	"log"
	"net/http"

	stan "github.com/nats-io/stan.go"
)

type Nats struct {
	ClusterID string
	ClientId string
	NameChanel string
	UrlBD string
}

var (
	cache map[string]bd.Order
	uid string
)

func (n *Nats) Start(){
	sc, er := stan.Connect(n.ClusterID, n.ClientId,stan.NatsURL("127.0.0.1:4223"))
	if er != nil {
		log.Println(er)
	}
	defer sc.Close()
	con, er := bd.Connect(n.UrlBD)
	if er != nil {
		log.Println(er)
	}
	defer bd.Close(con)
	cache, er = bd.RestoreCache(con)
	if er != nil {
		log.Println("failed to restore cache")
	}
	_, er = sc.Subscribe(n.NameChanel, func(m *stan.Msg) { //handler msg
		ord, fl := unmarshal([]byte(m.Data))
		if !fl {
			return
		}
		bd.AddEntry(con, ord)
		cache[ord.OrderUid] = *ord
		log.Println("Added data")
	})
	if er != nil {
		log.Println(er)
	}

	mux := http.NewServeMux() 
	mux.HandleFunc("/",Home)
	mux.HandleFunc("/show_order",ShowOrder)  

	log.Println("Start sever:")
	log.Fatal(http.ListenAndServe(":4000",mux))
}

func unmarshal(data []byte)(*bd.Order,bool){
	var ord bd.Order
	er := json.Unmarshal(data,&ord)
	if er != nil {
		log.Println("incorrect data was sent to the channel")
		return nil, false
	}
	if (ord.OrderUid == "" || ord.Items[0].ChrtId == 0){
		log.Println("structure came without unique parameters")
		return nil, false
	}
	return &ord, true
}

