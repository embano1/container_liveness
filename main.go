package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	twait := flag.Int("t", 10, "Specifies the amount of time in sec before exiting")
	errcode := flag.Int("e", 1, "Specifies the wanted exit code")
	flag.Parse()

	fmt.Printf("Starting...sleeping for %v seconds\n", *twait)
	time.Sleep(time.Duration(*twait) * time.Second)
	fmt.Printf("Oops, simulating weird behaviour and exiting(%v)\n", *errcode)
	os.Exit(*errcode)

}
