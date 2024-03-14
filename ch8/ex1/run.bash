#!/bin/bash

TZ=US/Eastern go run clock/main.go -port 8010 &
TZ=Europe/London go run clock/main.go -port 8020 &
TZ=Asia/Tokyo go run clock/main.go -port 8030 &

go run clockwall/main.go NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
