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
	//dividingFactor should not be more than maxRange
	if dividingFactor > maxRange {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}
	for i := minRange; i <= maxRange; i++ {
		if i%dividingFactor == 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, err
}

func HandleSlash(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, "/")
	startingPoint, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	//startingPoint should not be less than minRange
	if startingPoint < minRange {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}
	gap, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	//gap should not be more than maxRange
	if gap > maxRange {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
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

func HandleComma(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	parts := strings.Split(incomingText, ",")
	for i := 0; i < len(parts); i++ {

		// validation for whether the number is within the allowed range or not
		number, err := strconv.Atoi(parts[i])
		if err != nil {
			return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
		}
		if number < minRange || number > maxRange {
			return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
		}
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
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}

	endingPoint, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	// Handle the invalid ending point case
	if endingPoint > maxRange {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}
	for i := startingPoint; i <= endingPoint; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result, nil
}

func HandleSingleInteger(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	number, err := strconv.Atoi(incomingText)
	if err != nil {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}
	if number < minRange || number > maxRange {
		return nil, fmt.Errorf("err : %s %s", "invalid input passed", incomingText)
	}
	result = append(result, incomingText)
	return result, nil
}
