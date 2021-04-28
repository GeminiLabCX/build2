package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"
)

var (
	//go:embed VERSION
	version string
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
	fmt.Printf("🍖: %s\n", version)
}
