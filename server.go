package main

// https://github.com/BPI-SINOVOIP/RPi.GPIO/blob/master/create_gpio_user_permissions.py

import (
	"fmt"
	"log"
	"github.com/namsral/flag"
	"github.com/stianeikeland/go-rpio"
	"net/http"
  "os"
  "time"
)

var (
	port uint

	pinMap = map[string]rpio.Pin{
		"19": rpio.Pin(19),
		"26": rpio.Pin(26),
	}
)

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	pins, ok := r.URL.Query()["pin"]

	if !ok || len(pins[0]) < 1 {
		log.Println("Url Param 'pin' is missing")
		return
	}

	pin := string(pins[0])

	toggle(pin)
}

func readAndValidateConfig() {
	flag.UintVar(&port, "port", 8080, "Port to serve api on")
	flag.Parse()
}

func toggle(pin string) {
	pinMap[pin].High()
	time.Sleep(time.Second)
	pinMap[pin].Low()
}

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	for key, pin := range pinMap {
		fmt.Printf("Set pin for output %s\n", key)
		pin.Output()
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/toggle", toggleHandler)

	readAndValidateConfig()

	portNumber := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server at port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}
