# plug-go

Turn my lights on and off. Provides an HTTP endpoint and simple web page to
turn the lights on and off.

I found a cheap RF remote ([similar to this]) to turn on/off lamps, broke open the remote and connected
the buttons to a [relay], then connected the relay to a Raspberry Pi.

Port of <https://github.com/bkulyk/plug> server that was written in Ruby. I remember
Ruby taking a long time to install of the Raspberry Pi and I don't want to do that. Also I'm sure the version of ruby and gems were out of date by now. I like that I can compile a go application and target the Raspberry Pi's architecture.

## compile for pi

```sh
env GOOS=linux GOARCH=arm GOARM=5 go build server.go
```

## run on schedule

To schedule the lights I just used crontab to call the server.

```cron
30 16 * * * /usr/bin/curl http://localhost:8080/toggle?pin=26
30 21 * * * /usr/bin/curl http://localhost:8080/toggle?pin=19
```

[similar to this]: https://www.amazon.ca/Wireless-Electrical-Household-Appliances-ETL-Listed/dp/B07BH5S222/ref=sr_1_110?dchild=1&keywords=socket+remote+control+switch&qid=1605124921&sr=8-110
[relay]: https://www.amazon.ca/gp/product/B0057OC6D8/ref=ppx_yo_dt_b_search_asin_title?ie=UTF8&psc=1
