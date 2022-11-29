package comsoc

func STV_SWF(p Profile) (countRes Count, err error) {
	errProfil := CheckProfileAlternative(p, p[0])
	nbVotants := len(p)
	nbAlternatives := len(p[0])

	if errProfil == nil {
		countRes = make(Count)

		for i := 0; i < nbAlternatives-1; i++ { // tant qu'on n'a pas retiré tous les perdants

			countPoints := CountPointsSTV(p)

			for key, value := range countPoints { // vérifie s'il y a un score majoritaire
				if value > nbVotants/2 {
					countRes[key] = i + 1

					for k := range countPoints { // et mettre les autres à i
						if k != key {
							countRes[k] = i
						}
					}

					return countRes, nil
				}
			}

			perdant, gagnantExaequo := MinCount(countPoints) // on récupère la pire alternative
			if gagnantExaequo {
				for key := range countPoints {
					countRes[key] = i
				}

			} else {
				p = RetirerPerdant(p, perdant)
				countRes[perdant] = i

				if i == nbAlternatives-2 {
					countRes[p[0][0]] = i + 1
				}
			}
		}

		return countRes, nil

	} else {
		return countRes, errProfil
	}
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	res, err := STV_SWF(p)
	if err == nil {
		bestAlts = MaxCount(res)
	}
	return bestAlts, err
}
