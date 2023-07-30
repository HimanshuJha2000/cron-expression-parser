package main

import (
	"bufio"
	"fmt"
	"github.com/MyOrg/cron-expression-parser/internal/cron_expression"
	"log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your cron string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character from the input string
	inputCron := input[:len(input)-1]

	var cronExp cron_expression.CronExpression
	err = cronExp.GetCronExpression(inputCron)
	if err != nil {
		log.Println("Error in processing the cron expression")
	}
	return
}
