package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

const FILE_NAME = "results.xlsx"
const START_NUMBER_COL = 1
const FINISH_NUMBER_COL = 8
var gameList = make(map[int] Game)

type Game struct {
	ReleaseDate string
	Numbers string
}


func main() {
	file, err := excelize.OpenFile(FILE_NAME)
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
		game := []string{}
		releaseDate := ""

		for index, colCell := range row {
			if index > START_NUMBER_COL && index < FINISH_NUMBER_COL {
				game = append(game, colCell)
			}
			if index == START_NUMBER_COL {
				releaseDate = colCell
			}
		}

		if i > START_NUMBER_COL {
			gameList[i] = Game{
				Numbers: strings.Join(game, ","),
				ReleaseDate: releaseDate,
			} 
		}
	}

	for _, v := range gameList {
		fmt.Printf("date: %v, value: %v \n", v.ReleaseDate, v.Numbers )
	}

}