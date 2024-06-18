package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// func main() {
// 	// database, err := mongodb.SetUpDatabase("mongodb://root:example@localhost:27018/")
// 	// if err != nil {
// 	// 	panic("Falha ao se conectar ao banco \n")
// 	// }

// 	// broker, errBroker := broker.NewBrokerRabbit("amqp://iot:iot@192.168.1.156:5673/")
// 	// if errBroker != nil {
// 	// 	panic("Falha ao conectar no broker")
// 	// }
// 	// startApi(broker, database)
// 	http.Handle("/metrics", promhttp.Handler())
// 	http.ListenAndServe(":2112", nil)
// }

// func startApi(broker broker.IBroker, database *mongo.Database) {
// 	server := echo.New()
// 	api.SetRouts(broker, api.SetWebsocket(server), server, database)
// 	if err := server.Start(":1323"); err != nil {
// 		server.Logger.Fatal()
// 	}
// 	fmt.Println("server running")
// }

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
