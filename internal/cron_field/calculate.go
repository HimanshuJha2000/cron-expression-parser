package cron_field

import (
	"fmt"
	"strconv"
	"strings"
)

func HandleAsteriskSlash(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, "/")
	dividingFactor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	for i := minRange; i <= maxRange; i++ {
		if i%dividingFactor == 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, err
}

func HandleSlash(maxRange int, incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, "/")
	startingPoint, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	gap, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	for i := startingPoint; i <= maxRange; i += gap {
		result = append(result, strconv.Itoa(i))
	}
	return result, err
}

func HandleAsterisk(minRange int, maxRange int) ([]string, error) {
	var result []string
	for i := minRange; i <= maxRange; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result, nil
}

func HandleComma(incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, ",")
	for i := 0; i < len(parts); i++ {
		result = append(result, parts[i])
	}
	return result, nil
}

func HandleDash(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, "-")
	startingPoint, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	// Handle the invalid starting point case
	if startingPoint < minRange {
		return nil, fmt.Errorf("err %s", "Not a valid input")
	}

	endingPoint, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	// Handle the invalid ending point case
	if endingPoint > maxRange {
		return nil, fmt.Errorf("err %s", "Not a valid input")
	}
	for i := startingPoint; i <= endingPoint; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result, nil
}
