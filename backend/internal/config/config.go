package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServerConfig
	HTTPClientsConfig
}

type HTTPServerConfig struct {
	Host string `env:"HTTP_SERVER_HOST"`
	Port int    `env:"HTTP_SERVER_PORT"`
}

type HTTPClientsConfig struct {
	MWSClient          MWSClient
	YouTubeParseClient YouTubeParseClientConfig
	LLMClient          LLMClientConfig
}

type YouTubeParseClientConfig struct {
	YouTubeAPIKey string `env:"YOUTUBE_API_KEY"`
	VideosLimit   int64  `env:"YOUTUBE_VIDEOS_LIMIT"`
	CommentsLimit int64  `env:"YOUTUBE_COMMENTS_LIMIT"`
	Chart         string `env:"YOUTUBE_CHART"`
	RegionCode    string `env:"YOUTUBE_REGION_CODE"`
}

type LLMClientConfig struct {
	LLMUrl    string `env:"LLM_API_URL"`
	LLMApiKey string `env:"LLM_API_KEY"`
	LLMModel  string `env:"LLM_MODEL"`
}

type MWSClient struct {
	MWSUrl    string `env:"MWS_URL"`
	MWSToken  string `env:"MWS_TOKEN"`
	MWSViewID string `env:"MWS_VIEW_ID"`
}

func MustLoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	var cfg Config

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("Error reading env: " + err.Error())
	}

	return cfg
}
