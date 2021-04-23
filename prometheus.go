package common

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PrometheusBoot(port string)  {
	http.Handle("/metrics",promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf("%s:%s",IP,port),nil)
	if err != nil{
		log.Fatal(err)
	}
}