package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	localhostTime := time.Now().Round(0)
	currentNtpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Smth goes wrong. Error: \"%v\"", err)
	}

	//no need to do Round() due to stripMono has been already done into Local()
	fmt.Printf("current time: %s\nexact time: %s\n", localhostTime.String(), currentNtpTime.Local().String())
}
