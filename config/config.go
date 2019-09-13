package config

import (
	"log"
	"os"
	"strconv"
)

const defaultLimit = 20

// Settings holds the application config
type Settings struct {
	Port       string
	Title      string
	Banner     string
	LabelPlus  string
	LabelMinus string
	Limit      int
	Day        int
}

var settings *Settings

// Read the config from env vars
func Read() *Settings {
	banner := getEnvVar("BANNER", "/static/images/grass.jpg")
	if banner != "/static/images/grass.jpg" {
		banner = "data:image/png;base64," + banner
	}

	l := getEnvVar("LIMIT", string(defaultLimit))
	limit, err := strconv.Atoi(l)
	if err != nil {
		log.Println("Invalid limit, using", defaultLimit)
		limit = defaultLimit
	}
	d := getEnvVar("DAY", "0")
	day, err := strconv.Atoi(d)
	if err != nil || day > 6 {
		log.Print("Invalid day, using", "0")
		day = 0
	}

	settings = &Settings{
		Port:       getEnvVar("PORT", "8000"),
		Banner:     banner,
		Title:      getEnvVar("TITLE", "Sign-up Sheet"),
		LabelPlus:  getEnvVar("LABELPLUS", "Attending"),
		LabelMinus: getEnvVar("LABELMINUS", "Not attending"),
		Limit:      limit,
		Day:        day,
	}
	return settings
}

// Get the settings
func Get() *Settings {
	return settings
}

func getEnvVar(v, dflt string) string {
	value := os.Getenv(v)
	if len(value) == 0 {
		return dflt
	}
	return value
}
