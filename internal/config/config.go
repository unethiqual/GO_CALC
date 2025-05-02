package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	TimeAddition       time.Duration
	TimeSubtraction    time.Duration
	TimeMultiplication time.Duration
	TimeDivision       time.Duration
	ComputingPower     int
}

func LoadConfig() Config {
	cfg := Config{
		TimeAddition:       2000 * time.Millisecond,
		TimeSubtraction:    2000 * time.Millisecond,
		TimeMultiplication: 3000 * time.Millisecond,
		TimeDivision:       3000 * time.Millisecond,
		ComputingPower:     3,
	}

	file, err := os.Open(".env")
	if err != nil {
		log.Println("File .env not found. Using default values.")
		return cfg
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "TIME_ADDITION_MS":
			if v, err := strconv.Atoi(value); err == nil && v > 0 {
				cfg.TimeAddition = time.Duration(v) * time.Millisecond
			}
		case "TIME_SUBTRACTION_MS":
			if v, err := strconv.Atoi(value); err == nil && v > 0 {
				cfg.TimeSubtraction = time.Duration(v) * time.Millisecond
			}
		case "TIME_MULTIPLICATIONS_MS":
			if v, err := strconv.Atoi(value); err == nil && v > 0 {
				cfg.TimeMultiplication = time.Duration(v) * time.Millisecond
			}
		case "TIME_DIVISIONS_MS":
			if v, err := strconv.Atoi(value); err == nil && v > 0 {
				cfg.TimeDivision = time.Duration(v) * time.Millisecond
			}
		case "COMPUTING_POWER":
			if v, err := strconv.Atoi(value); err == nil && v > 0 {
				cfg.ComputingPower = v
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading .env file:", err)
	}

	return cfg
}
