package api

type Clien interface {
	//GetJoke returns one joke
	GetJoke() (*JokeResponse, error)
}
