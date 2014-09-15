package SubstitutionPlan

import (
	"time"
)

type SubstitutionPlan struct {
	Date         time.Time
	LastChange   time.Time
	Substitution []Substitution
}

type Substitution struct {
	Grade        string
	Hour         string
	Class        string
	SubstTeacher string
	Teacher      string
	Room         string
	UnkownField  string
	Description  string
	Type         string
}
