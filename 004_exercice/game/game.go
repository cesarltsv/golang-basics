package game

type Game struct {
	ReleaseDate string
	Numbers string
}

func Create(numbers string, date string) Game {
	return Game{
		Numbers: numbers,
		ReleaseDate: date,
	}
}