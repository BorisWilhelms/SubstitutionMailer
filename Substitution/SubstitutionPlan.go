package SubstitutionPlan

import (
	"errors"
	"strings"
)

func Parse(source string, grade string) (SubstitutionPlan, error) {
	res := SubstitutionPlan{}

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
