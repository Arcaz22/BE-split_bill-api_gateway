package config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Services struct {
		Auth struct {
			URL string `yaml:"url"`
		} `yaml:"auth"`
	} `yaml:"services"`
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func LoadConfig() error {
	cfg = Config{
		Server: struct {
			Port string `yaml:"port"`
		}{
			Port: "8080",
		},
		Services: struct {
			Auth struct {
				URL string `yaml:"url"`
			} `yaml:"auth"`
		}{
			Auth: struct {
				URL string `yaml:"url"`
			}{
				URL: "http://localhost:3000/v1",
			},
		},
	}
	return nil
}
