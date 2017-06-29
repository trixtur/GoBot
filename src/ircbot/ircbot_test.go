package ircbot

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"log"
)

var (
	myresponse = `{"port":6667}`
)

func TestConfigure(t *testing.T) {
	test_server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, myresponse)
	}))

	defer test_server.Close()

	request,err := http.NewRequest("GET", test_server.URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	config := configure(request)

	port := fmt.Sprint(config["port"])
	expect := "6667"

	// Test if port matches
	if port != expect {
		fmt.Println("port:", port)
		t.Error("expected " + expect + " received " + port);
	}
}
