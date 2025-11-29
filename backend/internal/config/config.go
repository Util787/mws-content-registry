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
	MWSUrl             string `env:"MWS_URL"`
	YouTubeParseClient YouTubeParseClient
}

type YouTubeParseClient struct {
	YouTubeAPIKey string `env:"YOUTUBE_API_KEY"`
	VideosLimit   int64    `env:"YOUTUBE_VIDEOS_LIMIT"`
	CommentsLimit int64    `env:"YOUTUBE_COMMENTS_LIMIT"`
	Chart         string `env:"YOUTUBE_CHART"`
	RegionCode    string `env:"YOUTUBE_REGION_CODE"`
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
