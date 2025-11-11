package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1.2233)
	err = useInsert(db, product)
	if err != nil {
		panic(err)
	}
	product.ID = "b6611c5e-a343-4477-bbce-ddc420a2ea3c"
	product.Name = "Naruto Uzumaki"
	err = useUpdate(db, product)
	if err != nil {
		panic(err)
	}
	selecteProduct, err := useSelect(db, "b6611c5e-a343-4477-bbce-ddc420a2ea3c")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %s \n", selecteProduct.Name)
	fmt.Printf("Price: %v \n", selecteProduct.Price)
	fmt.Printf("id: %s \n", selecteProduct.ID)

	productList, err := useSelectAll(db)
	if err != nil {
		panic(err)
	}
	for _, p := range productList {
		fmt.Printf("\n =================================== \n")
		fmt.Printf("Product: %s \n", p.Name)
		fmt.Printf("Price: %v \n", p.Price)
		fmt.Printf("id: %s \n", p.ID)
	}

}






