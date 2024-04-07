package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  Address
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  Company
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var users []User

	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatalln(err)
	}

	//for _, user := range users {
	//	fmt.Println(user.Name)
	//}
	fmt.Printf("%s", body)
}
