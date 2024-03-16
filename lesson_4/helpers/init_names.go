package helpers

import "errors"


type PossibleNames struct {
	part1 []string
	part2 []string
	part3 []string
}

func NewPossibleNames(part1, part2, part3 []string) (*PossibleNames, error) {
	if len(part1) == 0 || len(part2) == 0 || len(part3) == 0 {
		return nil, errors.New("slice with names can't be empty")
	}

	return &PossibleNames{
		part1: part1,
		part2: part2,
		part3: part3,
	},
	nil
}

func (r *PossibleNames) GroupUpNames() [][]string {
	mainSlice := make([][]string, 0, 3)
	mainSlice = append(mainSlice, r.part1, r.part2, r.part3)
	return mainSlice
}
