package main

import (
	"fmt"
	mpg123 "github.com/Golang-Commons/mediaCommons/mpg123bind"
	"golang.org/x/mobile/exp/audio"
	"os"
	"time"
)

func main() {
	mpg123.Initialize()
	gopath := os.Getenv("GOPATH")
	mp3, err := mpg123.Open(gopath + "/src/github.com/Golang-Commons/mediaCommons/mpg123bind/examples/example.mp3")
	if err != nil {
		panic(err)
	}

	rate, channels, encoding, format := mp3.Format()
	fmt.Printf("Rate: %i Channels: %i Encoding: %i Format: %s\n", rate, channels, encoding, format)

	p, err := audio.NewPlayer(mp3, audio.Format(format), rate)
	if err != nil {
		panic(err)
	}

	p.Play()
	for p.State() == audio.Playing {
		time.Sleep(time.Second)
	}
	mpg123.Exit()
}
