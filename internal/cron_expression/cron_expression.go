package cron_expression

import (
	"fmt"
	"github.com/MyOrg/cron-expression-parser/internal/cron_field"
	"github.com/MyOrg/cron-expression-parser/internal/utils"
	"strings"
)

type CronExpression struct {
	Minutes    string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

func (ce *CronExpression) GetCronExpression(inputCron string) error {
	return ce.formatOutputCronExpression(inputCron)
}

func (ce *CronExpression) calculateExpression(inputCron string) error {
	inputCronSplit := strings.Split(inputCron, " ")

	var cf cron_field.CronField

	minutes, err := cf.GetCronFieldData(inputCronSplit[0], utils.MINUTES)
	if err != nil {
		return fmt.Errorf("%s %s", "Error while fetching minutes from cron parser string", err)
	}
	ce.Minutes = strings.Join(minutes, " ")

	hour, err := cf.GetCronFieldData(inputCronSplit[1], utils.HOUR)
	if err != nil {
		return fmt.Errorf("%s %s", "Error while fetching hour from cron parser string", err)
	}
	ce.Hour = strings.Join(hour, " ")

	dayOfMonth, err := cf.GetCronFieldData(inputCronSplit[2], utils.DAY_OF_MONTH)
	if err != nil {
		return fmt.Errorf("%s %s", "Error while fetching day of month from cron parser string", err)
	}
	ce.DayOfMonth = strings.Join(dayOfMonth, " ")

	month, err := cf.GetCronFieldData(inputCronSplit[3], utils.MONTH)
	if err != nil {
		return fmt.Errorf("%s %s", "Error while fetching month from cron parser string", err)
	}
	ce.Month = strings.Join(month, " ")

	dayOfWeek, err := cf.GetCronFieldData(inputCronSplit[4], utils.DAY_OF_WEEK)
	if err != nil {
		return fmt.Errorf("%s %s", "Error while fetching day of week from cron parser string", err)
	}
	ce.DayOfWeek = strings.Join(dayOfWeek, " ")

	ce.Command = inputCronSplit[5]

	return nil
}

//func strValues(values []string) []string {
//	strVals := make([]string, len(values))
//	for i, v := range values {
//		strVals[i] = strconv.Itoa(v)
//	}
//	return strVals
//}

func (ce *CronExpression) formatOutputCronExpression(inputCron string) error {
	err := ce.calculateExpression(inputCron)
	if err != nil {
		return err
	}

	fmt.Printf("%-14s %s\n", utils.MINUTES, ce.Minutes)
	fmt.Printf("%-14s %s\n", utils.HOUR, ce.Hour)
	fmt.Printf("%-14s %s\n", utils.DAY_OF_MONTH, ce.DayOfMonth)
	fmt.Printf("%-14s %s\n", utils.MONTH, ce.Month)
	fmt.Printf("%-14s %s\n", utils.DAY_OF_WEEK, ce.DayOfWeek)
	fmt.Printf("%-14s %s\n", utils.COMMAND, ce.Command)

	return nil
}
