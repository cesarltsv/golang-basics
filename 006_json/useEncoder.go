package main

import (
	"encoding/json"
	"os"
)

func useEncoder() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(account)
}