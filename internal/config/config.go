package config

//Server config
type Server struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`

		JokeURL string `yaml:*"joke-url" env"JOKE_URL"`
	} `yaml:"server"`
}
