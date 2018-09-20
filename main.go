package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	producerMessageCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "producer_message_count",
		Help: "Number of message produce.",
	})
	producerDisconnectCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "producer_disconnect_count",
		Help: "Number of disconnection happen from producer.",
	})
	producerReconnectCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "producer_reconnect_count",
		Help: "Number of reconnection happen from producer.",
	})
	producerConnectCloseCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "producer_connect_close_count",
		Help: "Number of closed connection from producer.",
	})
	consumerMessageCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "consumer_message_count",
		Help: "Number of message consume.",
	})
	consumerDisconnectCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "consumer_disconnect_count",
		Help: "Number of disconnection happen from consumer.",
	})
	consumerReconnectCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "consumer_Reconnect_count",
		Help: "Number of Reconnection happen from consumer.",
	})
	consumerConnectCloseCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "consumer_connect_close_count",
		Help: "Number of closed connection from consumer.",
	})
)

func registerPrometheusMetrics() {
	prometheus.MustRegister(producerMessageCount)
	prometheus.MustRegister(producerDisconnectCount)
	prometheus.MustRegister(producerReconnectCount)
	prometheus.MustRegister(producerConnectCloseCount)
	prometheus.MustRegister(consumerMessageCount)
	prometheus.MustRegister(consumerDisconnectCount)
	prometheus.MustRegister(consumerReconnectCount)
	prometheus.MustRegister(consumerConnectCloseCount)

}

func main() {
	registerPrometheusMetrics()
	http.HandleFunc("/promesgcount", promesgcount)
	http.HandleFunc("/prodisconncount", prodisconncount)
	http.HandleFunc("/proreconncount", proreconncount)
	http.HandleFunc("/proconnclosscount", proconnclosscount)
	http.HandleFunc("/conmesgcount", conmesgcount)
	http.HandleFunc("/condisconncount", condisconncount)
	http.HandleFunc("/conreconncount", conreconncount)
	http.HandleFunc("/conconnclosscount", conconnclosscount)
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
}

func promesgcount(w http.ResponseWriter, r *http.Request) {
	producerMessageCount.Inc()
}
func prodisconncount(w http.ResponseWriter, r *http.Request) {
	producerDisconnectCount.Inc()
}
func proreconncount(w http.ResponseWriter, r *http.Request) {
	producerReconnectCount.Inc()
}
func proconnclosscount(w http.ResponseWriter, r *http.Request) {
	producerConnectCloseCount.Inc()
}
func conmesgcount(w http.ResponseWriter, r *http.Request) {
	consumerMessageCount.Inc()
}
func condisconncount(w http.ResponseWriter, r *http.Request) {
	consumerDisconnectCount.Inc()
}
func conreconncount(w http.ResponseWriter, r *http.Request) {
	consumerReconnectCount.Inc()
}
func conconnclosscount(w http.ResponseWriter, r *http.Request) {
	consumerConnectCloseCount.Inc()
}
