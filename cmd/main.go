package main

import (
	"task2/config"
	"task2/routes"
)

func main() {
	// Init config with default values
	cfg := config.Config{
		Addr:            "localhost:8080",
		InitialDataFile: "cmd/blog_data.json",
		UrlOrigin:       "http://localhost:8080",
	}
	// Update config with environment variables
	cfg.UpdateFromEnv()
	// Run main application
	routes.RunApp(&cfg)
}
