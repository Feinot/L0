package pub

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"nats/comands"

	"nats/models"
)

func Pub() {

	sc, err := comands.Init(models.Prod)
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	defer sc.Close()

	file, err := os.Open("./modesl.json")
	if err != nil {
		fmt.Println("publish error:", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("file write error:", err)
		return
	}
	ach := func(s string, err2 error) {}

	if _, err = sc.PublishAsync("test", []byte(data), ach); err != nil {
		fmt.Println("publish error:", err)
	}
	fmt.Println("Publish complite")

	time.Sleep(5 * time.Second)

}
