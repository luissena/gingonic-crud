package config

import "fmt"

var (
	logger *Logger
)

func Init() error {
	var err error

	if err != nil {
		return fmt.Errorf("Error initializing | %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
