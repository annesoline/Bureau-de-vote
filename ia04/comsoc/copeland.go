package comsoc

func CopelandSWF(p Profile) (Count, error) {

	err := CheckProfileAlternative(p, p[0])

	if err == nil {
		count := make(Count)
		nbVotants := len(p)
		alts := p[0]

		for i := 0; i < len(alts)-1; i++ {
			for j := i + 1; j < len(alts); j++ {
				points := 0
				for k := range p {
					if IsPref(alts[i], alts[j], p[k]) {
						points++
					}
				}

				if points > nbVotants/2 { // alts[i] a la majorité
					count[alts[i]]++
					count[alts[j]]--
				} else if points == nbVotants/2 { // égalité => +0
					continue
				} else { // alts[j] a la majorité
					count[alts[j]]++
					count[alts[i]]--
				}
			}
		}
		return count, nil
	} else {
		return nil, err
	}
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	res, err := CopelandSWF(p)
	if err == nil {
		bestAlts = MaxCount(res)
	}
	return bestAlts, err
}
