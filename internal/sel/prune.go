package sel

func Prune(sels []*Selection) []*Selection {
	newSels := make([]*Selection, 0, len(sels))
	merged := make([]bool, len(sels))

	for i, s := range sels {
		if merged[i] {
			continue
		}

		newS := s
		for j := i + 1; j < len(sels); j++ {
			if merged[j] {
				continue
			}
			s2 := sels[j]
			if newS.Overlaps(s2) {
				newS = newS.Merge(s2)
				merged[j] = true
			}
		}

		newSels = append(newSels, newS)
	}

	return newSels
}
