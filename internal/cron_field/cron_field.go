package cron_field

import (
	"github.com/MyOrg/cron-expression-parser/internal/utils"
	"strconv"
	"strings"
)

type CronField struct{}

var cronFieldRangeMap = make(map[string][]string)

func init() {
	cronFieldRangeMap[utils.MINUTES] = []string{utils.MinuteMin, utils.MinuteMax}
	cronFieldRangeMap[utils.HOUR] = []string{utils.HourMin, utils.HourMax}
	cronFieldRangeMap[utils.DAY_OF_MONTH] = []string{utils.DayOfMonthMin, utils.DayOfMonthMax}
	cronFieldRangeMap[utils.MONTH] = []string{utils.MonthMin, utils.MonthMax}
	cronFieldRangeMap[utils.DAY_OF_WEEK] = []string{utils.DayOfWeekMin, utils.DayOfWeekMax}
}

func (cf *CronField) GetCronFieldData(incomingText string, cronFieldType string) ([]string, error) {

	rangeArray := cronFieldRangeMap[cronFieldType]
	minRange, _ := strconv.Atoi(rangeArray[0])
	maxRange, _ := strconv.Atoi(rangeArray[1])

	result, err := expandCronField(minRange, maxRange, incomingText)
	return result, err
}

func expandCronField(minRange int, maxRange int, incomingText string) ([]string, error) {
	var result []string
	if strings.Contains(incomingText, "*/") {
		return HandleAsteriskSlash(minRange, maxRange, incomingText)
	} else if strings.Contains(incomingText, "/") {
		return HandleSlash(maxRange, incomingText)
	} else if strings.Contains(incomingText, "*") {
		return HandleAsterisk(minRange, maxRange)
	} else if strings.Contains(incomingText, ",") {
		return HandleComma(incomingText)
	} else if strings.Contains(incomingText, "-") {
		return HandleDash(minRange, maxRange, incomingText)
	} else {
		result = append(result, incomingText)
	}
	return result, nil
}
