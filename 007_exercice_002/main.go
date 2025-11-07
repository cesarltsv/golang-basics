package main

import (
	"encoding/json"
	"fmt"
)
var Zipcode = ""
type Account struct {
	Number int `json:"numero"`
	Amount int `json:"saldo"`
}

type Address struct {
      Zipcode string `json:"cep"`
      Street string  `json:"logradouro"`
      Complement string `json:"complemento"`
      Unit string  `json:"unidade"`
      Neighbourhood string `json:"bairro"`
      Location string  `json:"localidade"`
      StateId string  `json:"uf"`
      State string  `json:"estado"`
      Region string  `json:"regiao"`
      Ibge string  `json:"ibge"`
      Gia string  `json:"gia"`
      Ddd string  `json:"ddd"`
      Siafi string `json:"siafi"` 
}

func main() {
	var data Address
	showWelcomeMessage()
	getZipcode()
	err := json.Unmarshal(searchAddress(Zipcode), &data)
	if err != nil {
		errorToGetResult()
		return
	}
	showResult(data)
}

func getZipcode() {
	fmt.Print("zipcode: ")
	fmt.Scanln(&Zipcode)
}
