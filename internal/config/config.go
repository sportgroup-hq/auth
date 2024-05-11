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

type Config struct {
	HTTP     HTTP     `json:"http"`
	JWT      JWT      `json:"jwt"`
	Oauth    Oauth    `json:"oauth"`
	Log      Logger   `json:"logger"`
	GRPC     GRPC     `json:"grpc"`
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
}

type HTTP struct {
	Address string `json:"address" envconfig:"HTTP_ADDRESS" default:":8080"`
}

type JWT struct {
	Secret           string `json:"secret"             envconfig:"JWT_SECRET"             default:"my-totally-secret-key"`
	AccessExpiresIn  int    `json:"access_expires_in"  envconfig:"JWT_ACCESS_EXPIRES_IN"  default:"1"`
	RefreshExpiresIn int    `json:"refresh_expires_in" envconfig:"JWT_REFRESH_EXPIRES_IN" default:"168"`
}

type Oauth struct {
	Google Google `json:"google"`
}

type Google struct {
	ClientID     string   `json:"client_id"     envconfig:"OAUTH_GOOGLE_CLIENT_ID"`
	ClientSecret string   `json:"client_secret" envconfig:"OAUTH_GOOGLE_CLIENT_SECRET"`
	RedirectURL  string   `json:"redirect_url"  envconfig:"OAUTH_GOOGLE_REDIRECT_URL" default:"http://localhost:8081/oauth2callback"`
	Scopes       []string `json:"scopes"        envconfig:"OAUTH_GOOGLE_SCOPES"       default:"https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email openid"`
}

type Logger struct {
	Level string `json:"level" envconfig:"LOGGER_LEVEL" default:"info"`
}

type GRPC struct {
	API APIService `json:"api"`
}

type APIService struct {
	Address string `json:"address" envconfig:"GRPC_API_ADDRESS" default:"localhost:8180"`
}

type Postgres struct {
	Host     string `json:"host"     envconfig:"POSTGRES_HOST"     default:"localhost"`
	Port     int    `json:"port"     envconfig:"POSTGRES_PORT"     default:"5432"`
	Database string `json:"database" envconfig:"POSTGRES_DATABASE" default:"sportgroup_auth"`
	User     string `json:"user"     envconfig:"POSTGRES_USER"     default:"sportgroup_api_user"`
	Password string `json:"password" envconfig:"POSTGRES_PASSWORD"`
	Log      bool   `json:"log"      envconfig:"POSTGRES_LOG"      default:"true"`
}

type Redis struct {
	Host     string `json:"host"     envconfig:"REDIS_HOST"    default:"localhost:6379"`
	DB       int    `json:"db"       envconfig:"REDIS_DB"      default:"0"`
	Password string `json:"password" envconfig:"REDIS_PASSWORD"`
}

func New() (*Config, error) {
	var cfg Config

	flag.Parse()

	if err := goconfig.Init(&cfg, flag.Arg(0)); err != nil {
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
