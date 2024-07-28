package main
                     4683452
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Handler:        new(lbReqHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type lbReqHandler struct {

}
func (h *lbReqHandler) ServeHTTP(rs http.ResponseWriter,req *http.Request){
	fmt.Printf("%v from load balancer\n",req.URL)
	http
}

