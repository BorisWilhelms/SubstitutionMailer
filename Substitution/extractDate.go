package SubstitutionPlan

import (
	"regexp"
	"time"
)

func extractDate(s string) (time.Time, error) {
	return extractTime(s,
		"<div class=\"mon_title\">(\\d{2}.\\d?\\d.\\d{4}) \\w+</div>",
		"02.1.2006")
}

func extractLastChange(s string) (time.Time, error) {
	return extractTime(s, "Stand: (\\d{2}.\\d{2}.\\d{4} \\d{2}:\\d{2})",
		"02.1.2006 15:04")
}

func extractTime(s string, rexp string, format string) (time.Time, error) {
	r, _ := regexp.Compile(rexp)

	m := r.FindAllStringSubmatch(s, -1)
	date, err := time.Parse(format, m[0][1])

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
