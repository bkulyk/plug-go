package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namsral/flag"
	"github.com/stianeikeland/go-rpio"
)

var (
	port uint

	pinMap = map[string]rpio.Pin{
		"19": rpio.Pin(19),
		"26": rpio.Pin(26),
	}
)

func readAndValidateConfig() {
	flag.UintVar(&port, "port", 8080, "Port to serve api on")
	flag.Parse()
}

func togglePin(pin string) {
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

	readAndValidateConfig()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.GET("/toggle", toggleHandler)
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func toggleHandler(c echo.Context) error {
	pin := c.QueryParam("pin")
	togglePin(pin)
	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/?pin=%s", pin))
}
