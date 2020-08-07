package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {

	formData := url.Values{
		"name": {"Fabio Azevedo"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}
