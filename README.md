# plug-go

Turn my lights on and off. Provides an HTTP endpoint and simple web page to
turn the lights on and off.

I found a cheap RF remote ([similar to this]) to turn on/off lamps, broke open the remote and connected
the buttons to a [relay], then connected the relay to a Raspberry Pi.

To trigger a button on the remote, the codes sets a pin to high for a short time then sets it back to low. This shorts the button on the remote which turns on or off the plug. There are separate buttons for turning a plug on or off.

Port of [Plug], server that was written in Ruby. I remember
Ruby taking a long time to install of the Raspberry Pi and I don't want to do that. Also I'm sure the version of ruby and gems were out of date by now. I like that I can compile a go application and target the Raspberry Pi's architecture.

## compile for pi

```sh
env GOOS=linux GOARCH=arm GOARM=5 go build server.go
```

## run on schedule

To schedule the lights I just used crontab to call the server, to toggle specific pins.

```cron
30 16 * * * /usr/bin/curl http://localhost:8080/toggle?pin=26
30 21 * * * /usr/bin/curl http://localhost:8080/toggle?pin=19
```

[Plug]: https://github.com/bkulyk/plug
[similar to this]: https://www.amazon.ca/Wireless-Electrical-Household-Appliances-ETL-Listed/dp/B07BH5S222/ref=sr_1_110?dchild=1&keywords=socket+remote+control+switch&qid=1605124921&sr=8-110
[relay]: https://www.amazon.ca/Homyl-Channel-Module-Computer-Product/dp/B08375GTC3/ref=sr_1_110?crid=28AK3GQSPJ0GY&dchild=1&keywords=8+channel+relay&qid=1605332021&s=electronics&sprefix=8+channe%2Celectronics%2C270&sr=1-110
