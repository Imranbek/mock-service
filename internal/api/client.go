package api

type Client interface {
	//GetJoke returns one joke
	GetJoke() (*JokeResponse, error)
}
