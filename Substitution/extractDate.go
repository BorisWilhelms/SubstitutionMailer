package SubstitutionPlan

import (
	"regexp"
	"time"
)

func extractDate(s string) (time.Time, error) {
	r, _ := regexp.Compile("<div class=\"mon_title\">(\\d{2}.\\d?\\d.\\d{4}) \\w+</div>")
	m := r.FindAllStringSubmatch(s, -1)
	date, err := time.Parse("02.1.2006", m[0][1])

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func extractLastChange(s string) (time.Time, error) {
	r, _ := regexp.Compile("Stand: (\\d{2}.\\d{2}.\\d{4} \\d{2}:\\d{2})")

	m := r.FindAllStringSubmatch(s, -1)
	date, err := time.Parse("02.1.2006 15:04", m[0][1])

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
