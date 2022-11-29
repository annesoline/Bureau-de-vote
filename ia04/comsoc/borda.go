package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	// renvoie un décompte à partir d'un profil
	err = CheckProfileAlternative(p, p[0])
	if err == nil {
		count := make(Count)
		for i := range p[0] {
			count[p[0][i]] = 0
		}
		nbAlts := len(p[0])
		for i := range p {
			for j := range p[i] {
				count[p[i][j]] += nbAlts - j - 1
			}
		}
		return count, err
	} else {
		return count, err
	}
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	if err == nil {
		bestAlts = MaxCount(count)
	}
	return bestAlts, err
}
