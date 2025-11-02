package main

import "fmt"

type Client struct {
	Name string
}

func NewClient() *Client {
	return &Client{Name: "teste user"}
}

func (c *Client) walk() {
	c.Name = "Herry potter"
	fmt.Printf("The Client %s walk \n", c.Name)
}

func usePointerStructs() {
	client := Client{ 
		Name: "naruto",
	}

	client.walk()
	fmt.Printf("Client Name: %s \n", client.Name)
}