package v1

import (
	"fmt"
	"log"
	"os"
)

func init() {
	fmt.Println(".....")
	home := os.Getenv("HOME11")
	os.Getuid()
	if home == "" {
		log.Fatal("home is not set ")
	}
	fmt.Println(home)
}
