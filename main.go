package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Pokemon struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Base_experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Is_default      bool   `json:"is_default"`
	Order           int    `json:"order"`
	Weight          int    `json:"weight"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Pokemon
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("id:", responseObject.Id)
	fmt.Println("name:", responseObject.Name)
	fmt.Println("base_experience:", responseObject.Base_experience)
	fmt.Println("height:", responseObject.Height)
	fmt.Println("Is Default:", responseObject.Is_default)
	fmt.Println("order", responseObject.Order)
	fmt.Println("Weight", responseObject.Weight)
}
