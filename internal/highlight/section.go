package highlight

type Section struct {
	StartLine int
	StartIdx  int
	EndLine   int
	EndIdx    int

	Region *Region
}

func (s Section) Analyze() []Highlight {
	return nil
}
