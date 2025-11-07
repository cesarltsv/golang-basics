package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)
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
	http.HandleFunc("/", AddressController)
	http.ListenAndServe(":8080", nil)
}

func AddressController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path %s \n", r.URL.Path)
	zipcode := r.URL.Query().Get("zipcode")
	if len(zipcode) < 7 || zipcode == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("address invalid1"))
		return
	}

	address, err := searchAddress(zipcode)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(address)
	fmt.Printf("%v", address)
}

func searchAddress(zipcode string) (*Address, error) {
	var address Address
	url := "https://viacep.com.br/ws/" + zipcode + "/json/"
	response, err := http.Get(url)
	if err != nil {
		return &Address{}, errors.New("Address not found")
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return &Address{}, errors.New("internal server error")
	}
	err = json.Unmarshal(data, &address)
	if err != nil {
		return &Address{}, errors.New("internal server error")
	}
	return &address, nil
}