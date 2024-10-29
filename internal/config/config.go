package config

import "github.com/lpernett/godotenv"

func Load(path string) error {
	err := godotenv.Load(path)
	return err
}
