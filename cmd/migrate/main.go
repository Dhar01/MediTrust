package main

import (
	"errors"
	"fmt"
	"io/fs"
	"medicine-app/config"
	"medicine-app/internal/database"
	"os"
	"os/exec"
	"strings"
)

type migration string

const (
	migrationUp   migration = "up"
	migrationDown migration = "down"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Welcome to the migration!")
	fmt.Println("Note: This migration only works for postgresql")

	cfg, err := getConfig()
	if err != nil {
		return err
	}

	dbURL := database.GetDSN(cfg.Database)

	if dbURL == "" {
		return fmt.Errorf("no URL found")
	}

	input, err := getInput()
	if err != nil {
		return err
	}

	directory := "sql/schema"
	ok, err := exists(directory)
	if err != nil || !ok {
		fmt.Printf("%s directory not exists!\n", directory)
		return err
	}

	fmt.Printf("%s directory exists.\n", directory)

	if err := os.Chdir(directory); err != nil {
		return err
	}

	fmt.Printf("Running goose %s...\n", string(input))

	cmd := exec.Command("goose", "postgres", dbURL, string(input))
	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("goose failed: %w\nOutput:\n%s", err, out) 
	}

	fmt.Println(string(out))
	return nil
}

// get the user input
func getInput() (migration, error) {
	var input string
	fmt.Print("\nEnter command (up/down): ")
	fmt.Scan(&input)

	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case string(migrationUp):
		return migrationUp, nil
	case string(migrationDown):
		return migrationDown, nil
	default:
		return "", fmt.Errorf("invalid input: must be 'up' or 'down'")
	}
}

// get the config for DSN
func getConfig() (*config.Config, error) {
	if err := config.LoadConfig(); err != nil {
		return nil, err
	}

	cfg := config.GetConfig()

	if cfg == nil {
		return nil, fmt.Errorf("can't access config")
	}

	return cfg, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}
