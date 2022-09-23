package comands

import (
	"log"
	"nats/models"

	"github.com/nats-io/stan.go"
)

func Init(clientID string) (stan.Conn, error) {
	sc, err := stan.Connect(models.Cluster, clientID,
		stan.NatsURL(stan.DefaultNatsURL),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Printf("connect error: %v", reason)
		}))
	return sc, err

}
