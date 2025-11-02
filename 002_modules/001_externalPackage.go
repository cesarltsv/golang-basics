package main

import (
	"fmt"

	"github.com/google/uuid"
)

func useExternalPackage() {
	id := uuid.New().String()
	fmt.Printf("UUID: %s \n", id)
}