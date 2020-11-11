#1/usr/bin/env sh

# https://www.thepolyglotdeveloper.com/2017/04/cross-compiling-golang-applications-raspberry-pi/
env GOOS=linux GOARCH=arm GOARM=5 go build server.go
