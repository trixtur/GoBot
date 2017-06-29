/*
 * The main driver of the application. You must supply a config url as a command-line argument. The config API must return JSON.
 */
package ircbot

import (
		"os"
		"fmt"
		"io/ioutil"
		"log"
		"net/http"
		"encoding/json"
)

func configure(url string) (config map[string]interface{}) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &config); err != nil {
			panic(err)
		}
	}

	return
}

func main() {
	url := os.Args[1]

	config := configure(url)
	fmt.Println("map:", config)
	fmt.Println("port:", config["port"])
	fmt.Println("channels:", config["channels"])
}

