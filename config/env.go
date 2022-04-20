package config

import (
	"os"
)

type Env struct {
	Host     string
	Port     string
	Protocol string
	BaseUrl  string
}

func GetEnv() Env {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	protocol := "http"
	if port == "443" {
		protocol = "https"
	}

	baseUrl := protocol + "://" + host + ":" + port
	if port == "443" || port == "80" {
		baseUrl = protocol + "://" + host
	}

	return Env{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		BaseUrl:  baseUrl,
	}
}
