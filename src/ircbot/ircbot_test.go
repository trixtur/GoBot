package ircbot

import (
	"testing"
	"fmt"
	"os"
)

func TestConfigure(t *testing.T) {
	url := os.Args[2]

	config := configure(url)

	port := fmt.Sprint(config["port"])

	// Test if port matches
	if port != "6667" {
		fmt.Println("port:", port)
		t.Error("expected 6667 received " + fmt.Sprintf("",port));
	}
}
