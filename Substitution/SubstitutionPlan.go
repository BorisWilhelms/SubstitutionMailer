package SubstitutionPlan

import (
	"errors"
	"strings"
	"time"
)

type SubstitutionPlan struct {
	Date         time.Time
	LastChange   time.Time
	Substitution []Substitution
}

func Parse(source string, grade string) (SubstitutionPlan, error) {
	res := SubstitutionPlan{}

	if strings.Index(source, grade) == -1 {
		return res, nil
	}

	date, err := extractDate(source)
	if err != nil {
		return SubstitutionPlan{}, errors.New("Could not extract date from source")
	}

	change, err := extractLastChange(source)
	if err != nil {
		return SubstitutionPlan{}, errors.New("Could not extract last change from source")
	}

	res.Date = date
	res.LastChange = change
	subst := extractSubstitutions(source)

	for _, v := range subst {
		if strings.Index(v.Grade, grade) != -1 {
			res.Substitution = append(res.Substitution, v)
		}
	}

	return res, nil
}
