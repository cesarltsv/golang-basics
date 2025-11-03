package organize

import (
	"strings"

	"loto.com/script/game"
	"loto.com/script/utils"
)



func HandleGameObject(currentCol int, numbers []string, date string) game.Game {
	if currentCol > utils.StartColumn {
		return game.Create(strings.Join(numbers, ","), date, currentCol)
	}
	return game.Create("", "", currentCol)
}

func GetGamesData(row []string) ([]string, string) {
	game := []string{}
	releaseDate := ""

	for index, colCell := range row {
		if index >  utils.StartColumn  && index < utils.FinishColumn {
			game = append(game, colCell)
		}
		if index ==  utils.StartColumn  {
			releaseDate = colCell
		}
	}

	return game, releaseDate

}