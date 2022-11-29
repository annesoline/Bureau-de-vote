package comsoc

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = CheckProfileAlternative(p, p[0])

	if err != nil {
		return nil, err
	}

	count = make(Count)
	for _, val := range p[0] {
		count[val] = 0
	}

	for index, valeur := range thresholds {
		for i := 0; i < valeur; i++ {
			count[p[index][i]]++
		}
	}

	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	//count := make(Count)
	count, err := ApprovalSWF(p, thresholds)

	if err != nil {
		return nil, err
	}

	return MaxCount(count), err
}
