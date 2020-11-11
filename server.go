package main

// https://github.com/BPI-SINOVOIP/RPi.GPIO/blob/master/create_gpio_user_permissions.py

import (
	"fmt"
	"log"
	"github.com/namsral/flag"
	"github.com/stianeikeland/go-rpio"
	"net/http"
  // "os"
  "time"
)

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	pins, ok := r.URL.Query()["pin"]

	if !ok || len(pins[0]) < 1 {
		log.Println("Url Param 'pin' is missing")
		return
	}

	pin := string(pins[0])

  log.Println("Url Param 'pin' is: " + pin)

  pinMap[pin].High()
  time.Sleep(time.Second)
  pinMap[pin].Low()
}

var pinMap map[string]rpio.Pin
var port uint

func initPins() {
  if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  defer rpio.Close()

  pinMap = make(map[string]rpio.Pin)
	pinMap["26"] = rpio.Pin(26)
	pinMap["19"] = rpio.Pin(19)
}

func readAndValidateConfig() {
	flag.UintVar(&port, "port", 8080, "Port to serve api on")
	flag.Parse()
}

func main() {
	initPins()
	readAndValidateConfig()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/toggle", toggleHandler)

	portNumber := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server at port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal(err)
	}
}
