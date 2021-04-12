package entity

type Position struct {
	rule calculationRuleProvider
}

type calculationRuleProvider interface {
	Calculates(employee Employee) float64
}

func NewPosition(calculationRule calculationRuleProvider) *Position {
	return &Position{
		rule: calculationRule,
	}
}
