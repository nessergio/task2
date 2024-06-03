package config

import "os"

type Config struct {
	Addr            string
	InitialDataFile string
	UrlOrigin       string
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func (c Config) UpdateFromEnv() {
	c.Addr = getEnv("PORT", c.Addr)
	c.InitialDataFile = getEnv("INITIAL_DATA_FILE", c.InitialDataFile)
	c.UrlOrigin = getEnv("URL_ORIGIN", c.UrlOrigin)
}
