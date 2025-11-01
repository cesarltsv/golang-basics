package main

import "fmt"

func useMaps() {
	users := map[string]string{"elsa": "geladinha", "zuko": "esquentadinho", "gelado": "fica frio aí"}
	fmt.Printf("map: %v \n", users)
	fmt.Printf("else nickName: %v \n", users["elsa"])
	fmt.Printf("zuko nickName: %v \n", users["zuko"])
	fmt.Printf("gelado nickName: %v \n", users["gelado"])
	users["izuku"] = "bomberMan"
	fmt.Printf("map updated: %v \n", users)
	delete(users, "zuko")
	fmt.Printf("map updated 2: %v \n", users)

	defaultMap := make(map[string]int)
	fmt.Printf("Default value: %v \n", defaultMap)

	fmt.Printf("\n **** loop a map **** \n \n")
	for name, nickName := range users {
		fmt.Printf("%s has a nickName equal: %s \n", name, nickName)
	}
}
