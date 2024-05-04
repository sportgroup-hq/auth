package config

import (
	"flag"
	"fmt"
	"sync"

	goconfig "github.com/Yalantis/go-config"
)

var (
	config        Config
	isInitialised bool
	mutex         = new(sync.Mutex)
)

const configFilename = "config.json"

type Config struct {
	Address   string   `json:"address" default:":8080"`
	JwtSecret string   `json:"jwt_secret" default:"my-totally-secret-key"`
	Oauth     Oauth    `json:"oauth"`
	Log       Logger   `json:"logger"`
	Services  Services `json:"services"`
}

type Oauth struct {
	Google Google `json:"google"`
}

type Google struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url" default:"http://localhost:8080/api/v1/test"`
	Scopes       []string `json:"scopes" default:"https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email openid"`
}

type Logger struct {
	Level string `json:"level" default:"info"`
}

type Services struct {
	API APIService `json:"api"`
}

type APIService struct {
	Address string `json:"address" default:"localhost:8081"`
}

func New() (Config, error) {
	var cfg Config

	filename := configFilename

	flag.Parse()

	if flag.Arg(0) != "" {
		filename = flag.Arg(0)
	}

	if err := goconfig.Init(&cfg, filename); err != nil {
		return Config{}, fmt.Errorf("init config: %w", err)
	}

	config = cfg
	isInitialised = true

	return cfg, nil
}

func Get() Config {
	mutex.Lock()
	if !isInitialised {
		cfg, err := New()
		if err != nil {
			panic(err)
		}
		config = cfg
	}
	mutex.Unlock()

	return config
}
