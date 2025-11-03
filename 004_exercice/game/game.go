package game

type Game struct {
	Id int
	ReleaseDate string
	Numbers string
}

func Create(numbers string, date string, index int) Game {
	return Game{
		Id: index + 1,
		Numbers: numbers,
		ReleaseDate: date,
	}
}