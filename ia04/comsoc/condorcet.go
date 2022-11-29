package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	err = CheckProfileAlternative(p, p[0])

	if err != nil {
		return nil, err
	}

	var sc_candidat int = 0
	var sc_opposant int = 0

	for _, candidat := range p[0] {
		sc_opposant = 0
		sc_candidat = 0
		for _, opposant := range p[0] {
			for _, val := range p {
				if candidat != opposant && IsPref(candidat, opposant, val) {
					sc_candidat = sc_candidat + 1
				}
				if candidat != opposant && IsPref(opposant, candidat, val) {
					sc_opposant = sc_opposant + 1
				}
			}
		}
		if sc_candidat > sc_opposant {
			result := []Alternative{candidat}
			return result, err
		}

	}
	return nil, err
}
