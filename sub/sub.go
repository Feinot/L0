package sub

import (
	"encoding/json"
	"fmt"
	"log"

	"nats/models"
	"nats/storage/db"

	"nats/comands"

	"github.com/nats-io/stan.go"
)

func Subscribe() {
	or := models.Order{}
	//db.Connect()

	sc, err := comands.Init(models.Sub)
	if err != nil {
		fmt.Println(err)
	}

	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("failed to ACK msg:%v", err)

		}
		if err := json.Unmarshal(msg.Data, &or); err != nil {
			fmt.Println("can't unmarshal: ", err.Error())
		}

		if err = or.Validate(); err != nil {
			fmt.Printf("Nats.validate: %s\n", err)

			return
		}
		fmt.Printf("msg: %s\n", string(msg.Data))

		//cashe.Set(or)
		db.Insert(or)

	}

	_, err = sc.QueueSubscribe("test",
		"test", mcb,
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName("test1"))

	if err != nil {
		log.Println(err)
	}

}
