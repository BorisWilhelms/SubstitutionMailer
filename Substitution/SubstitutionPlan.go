package SubstitutionPlan

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type SubstitutionPlan struct {
	Date time.Time
}

func Parse(source string, class string) (SubstitutionPlan, error) {
	s := SubstitutionPlan{}

	if strings.Index(source, class) == -1 {
		return s, nil
	}

	date, err := extractDate(source)
	if err != nil {
		return SubstitutionPlan{}, errors.New("Could not extract date from source")
	}

	s.Date = date

	return s, nil
}

func extractDate(source string) (time.Time, error) {
	r, _ := regexp.Compile("<div class=\"mon_title\">(\\d{2}.\\d?\\d.\\d{4}) \\w+</div>")
	m := r.FindAllStringSubmatch(source, -1)
	date, err := time.Parse("02.1.2006", m[0][1])

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
