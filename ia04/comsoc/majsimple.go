package comsoc

func MajoritySWF(p Profile) (count Count, err error) {
	// renvoie un décompte à partir d'un profil
	err = CheckProfileAlternative(p, p[0])
	if err == nil {
		count := make(Count)
		for i := range p[0] {
			count[p[0][i]] = 0
		}
		for i := range p {
			count[p[i][0]]++
		}
		return count, err
	} else {
		return count, err
	}
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	// renvoie les meilleurs alternatives à partir d'un profil
	count := make(Count)
	count, err = MajoritySWF(p)
	if err == nil {
		bestAlts = MaxCount(count)
	}
	return bestAlts, err
}
