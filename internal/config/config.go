package config

import (
	"flag"
	"fmt"
	"sync"

	goconfig "github.com/Yalantis/go-config"
)

var (
	config        *Config
	isInitialised bool
	mutex         = new(sync.Mutex)
)

const configFilename = "config.json"

type Config struct {
	HTTP     HTTP     `json:"http"`
	JWT      JWT      `json:"jwt"`
	Oauth    Oauth    `json:"oauth"`
	Log      Logger   `json:"logger"`
	Services Services `json:"services"`
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
}

type HTTP struct {
	Address string `json:"address" default:":8081"`
}

type JWT struct {
	Secret           string `json:"secret" default:"my-totally-secret-key"`
	AccessExpiresIn  int    `json:"access_expires_in" default:"1"`
	RefreshExpiresIn int    `json:"refresh_expires_in" default:"168"`
}

type Oauth struct {
	Google Google `json:"google"`
}

type Google struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url" default:"http://localhost:8081/oauth2callback"`
	Scopes       []string `json:"scopes" default:"https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email openid"`
}

type Logger struct {
	Level string `json:"level" default:"info"`
}

type Services struct {
	API APIService `json:"api"`
}

type APIService struct {
	Address string `json:"address" default:"localhost:8180"`
}

type Postgres struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"5432"`
	Database string `json:"database" default:"oss"`
	User     string `json:"user" default:"oss"`
	Password string `json:"password"`
	Log      bool   `json:"log" default:"true"`
}

type Redis struct {
	Address  string `json:"address" default:"localhost:6379"`
	Password string `json:"password"`
	DB       int    `json:"db" default:"0"`
}

func New() (*Config, error) {
	var cfg Config

	filename := configFilename

	flag.Parse()

	if flag.Arg(0) != "" {
		filename = flag.Arg(0)
	}

	if err := goconfig.Init(&cfg, filename); err != nil {
		return nil, fmt.Errorf("init config: %w", err)
	}

	config = &cfg
	isInitialised = true

	return &cfg, nil
}

func Get() *Config {
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
