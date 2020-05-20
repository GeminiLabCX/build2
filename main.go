package main

import (
	"fmt"
	"os"
	"time"
)

var (
	buildstamp string = "not set"
	githash    string = "not set"
	goversion  string = "not set"
)

func main() {
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Printf("🗿: Hello Build!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format(time.RFC3339))
	fmt.Printf("🎆: %s\n", githash)
	fmt.Printf("💈: %s\n", buildstamp)
	fmt.Printf("💎: %s\n", goversion)
}
