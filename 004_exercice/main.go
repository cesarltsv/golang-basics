package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	"loto.com/script/game"
	"loto.com/script/organize"
	"loto.com/script/utils"
)


var gameList = make(map[int] game.Game)

func main() {
	file, err := excelize.OpenFile(utils.FileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	
	for i, row := range rows {
		game, releaseDate := organize.GetGamesData(row)
		gameList[i] = organize.HandleGameObject(i, game, releaseDate)
	}

	for _, v := range gameList {
		fmt.Printf("date: %v, value: %v \n", v.ReleaseDate, v.Numbers )
	}

}