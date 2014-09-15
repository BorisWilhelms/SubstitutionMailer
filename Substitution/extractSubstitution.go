package SubstitutionPlan

import (
	"errors"
	"log"
	"regexp"
)

var substRegex, _ = regexp.Compile(">([[:alpha:]\\d,äÄöÖüÜß?\\-\\.\\/\\+ |&nbsp;]+)<")

func extractSubstitutions(source string) []Substitution {
	var res []Substitution

	r := extractRows(source)

	for _, v := range r {
		if len(v) != 2 {
			log.Print("Found line with length not equal to two")
			log.Print(v)
			continue
		}

		subst, err := extractSubstitution(v[0])
		if err != nil {
			log.Print("Error parsing row")
			log.Print(v[0])
			continue
		}
		res = append(res, subst)
	}

	return res
}

func extractSubstitution(r string) (Substitution, error) {
	var result = Substitution{}
	m := substRegex.FindAllStringSubmatch(r, -1)

	if len(m) != 9 {
		return result, errors.New("The row has not nine items...")
	}

	result.Grade = m[0][1]
	result.Hour = m[1][1]
	result.Class = m[2][1]
	result.SubstTeacher = m[3][1]
	result.Teacher = m[4][1]
	result.Room = m[5][1]
	result.UnkownField = m[6][1]
	result.Description = m[7][1]
	result.Type = m[8][1]

	return result, nil
}

func extractRows(source string) [][]string {
	r, _ := regexp.Compile("(<tr.*\n)")
	m := r.FindAllStringSubmatch(source, -1)

	return m[:]
}
