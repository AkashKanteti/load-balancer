package be

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":69",
		Handler:        new(ok),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type ok struct {

}
func (h *ok) ServeHTTP(rs http.ResponseWriter,req *http.Request){
	fmt.Printf("%v from load balancer\n",req.URL)
	be.ReceiveRequests(rs,req)
}
