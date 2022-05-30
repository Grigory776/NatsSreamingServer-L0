package nats

import (
	//"NatsStreamingServer/MainServer/bd"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{ 
		http.NotFound(w,r)
		return
	}
	if r.Method == "POST" {
        er := r.ParseForm()
		if er != nil {
			log.Println(er)
		}
		uid = r.FormValue("model")
		http.Redirect(w,r,"/show_order",http.StatusFound)
    }else{
        http.ServeFile(w,r, "../ui/html/home.html")
    }
}

func ShowOrder (w http.ResponseWriter, r *http.Request){
	data, fl := cache[uid]
	if fl{
		t, _ := template.ParseFiles("../ui/html/showOrder.html")
		t.Execute(w, &data)
	}else{
		t, _ := template.ParseFiles("../ui/html/notFound.html")
		t.Execute(w, &data)
	}
}