package main

type Account struct {
	Number int `json:"numero"`
	Amount int `json:"saldo"`
}

var account = Account{
	Number: 123,
	Amount: 213123123,
}


func main() {
	splitSection("BASIC JSON - MARSHEL")
	useMarshel()

	splitSection("BASIC JSON - ENCONDER")
	useEncoder()

	splitSection("BASIC JSON - json to struct")
	useJsonToStruct()

	
}