package comsoc

import (
	"errors"
	// "fmt"
)

func TieBreakFactory(alt []Alternative) func(alt []Alternative) (gagnant Alternative, err error) {
	// fonction qui crée une fonction TieBreak renvoyant 1 seule Alternative gagnante (le plus grand nombre),
	// à partir d'un tableau d'ordre stricte d'Alternatives dont les scores sont égaux

	return func(alt []Alternative) (gagnant Alternative, err error) {

		if len(alt) == 0 {
			return 0, errors.New("Erreur : tableau d'alternatives vide")

		} else if len(alt) == 1 {
			return alt[0], nil

		} else {
			gagnant = 0
			for i := range alt {
				if alt[i] > gagnant {
					gagnant = alt[i]
				}
			}
			return gagnant, nil
		}
	}
}

func TieBreakSmallest(alt []Alternative) (gagnant Alternative) {

	if len(alt) == 1 {
		return alt[0]

	} else {
		gagnant = alt[0]
		for i := range alt {
			if alt[i] < gagnant {
				gagnant = alt[i]
			}
		}
		return gagnant
	}
}

func SWFFactory(SWF func(p Profile) (countSWF Count, errSWF error), TieBreak func(bestAlts []Alternative) (gagnant Alternative, err error)) func(p Profile) (resSWF []Alternative, err error) {
	// fontion qui, à partir d'une fonction SWF, et d'une fonction TieBreak,
	// renvoie une fonction SWF sans ex aequo => tableau d'alternatives, d'ordre strict, du meilleur au pire

	return func(p Profile) (resSWF []Alternative, err error) {

		countSWF, errSWF := SWF(p)

		if errSWF == nil {
			i := 0
			for i < len(countSWF)+1 { // à chaque tour, on récupère l'alternative gagnant, et on le retire du countSWF
				// au cas où il y ait plusieurs égalités de scores différents
				// fmt.Println(countSWF)
				bestAlts := MaxCount(countSWF) // on récupère la ou les meilleurs alternatives

				if len(bestAlts) > 1 { // il y a ex aequo
					gagnant, errTB := TieBreak(bestAlts) // on récupère donc le gagnant avec un tie break

					if errTB == nil {
						resSWF = append(resSWF, gagnant)
						delete(countSWF, gagnant)

					} else {
						return nil, errTB
					}

				} else if len(bestAlts) == 1 { //
					resSWF = append(resSWF, bestAlts[0])
					delete(countSWF, bestAlts[0])
				}
				i++
			}
			bestAlts := MaxCount(countSWF)
			resSWF = append(resSWF, bestAlts[0])
			return resSWF, nil

		} else {
			return nil, errSWF
		}
	}
}

func SCFFactory(SCF func(p Profile) (resSCF []Alternative, errSCF error), TieBreak func(resSCF []Alternative) (gagnant Alternative, err error)) func(p Profile) (gagnant Alternative, err error) {
	// fonction qui, à partir d'une fonction SCF, et d'une fonction TieBreak,
	// renvoie une fonction SCF sans ex aequo => valeur unique

	return func(p Profile) (gagnant Alternative, err error) {

		resSCF, errSCF := SCF(p)

		if errSCF == nil {

			if len(resSCF) > 1 { // il y a ex aequo !
				gagnant, errTB := TieBreak(resSCF) // on détermine le gagnant avec le TieBreak

				if errTB == nil {
					return gagnant, nil
				} else {
					return 0, errTB
				}

			} else { // pas d'ex aequo
				return resSCF[0], nil // resSCF[0] est le seul gagnant (et seule valeur dans le tableau)
			}

		} else {
			return 0, errSCF
		}
	}
}

func SWFFactoryApproval(SWF func(p Profile, thresholds []int) (countSWF Count, errSWF error), TieBreak func(bestAlts []Alternative) (gagnant Alternative, err error)) func(p Profile, thresholds []int) (resSWF []Alternative, err error) {
	// fontion qui, à partir d'une fonction SWF, et d'une fonction TieBreak,
	// renvoie une fonction SWF sans ex aequo => tableau d'alternatives, d'ordre strict, du meilleur au pire

	return func(p Profile, thresholds []int) (resSWF []Alternative, err error) {

		countSWF, errSWF := SWF(p, thresholds)

		if errSWF == nil {
			i := 0
			for i < len(countSWF)+1 { // à chaque tour, on récupère l'alternative gagnant, et on le retire du countSWF
				// au cas où il y ait plusieurs égalités de scores différents
				// fmt.Println(countSWF)
				bestAlts := MaxCount(countSWF) // on récupère la ou les meilleurs alternatives

				if len(bestAlts) > 1 { // il y a ex aequo
					gagnant, errTB := TieBreak(bestAlts) // on récupère donc le gagnant avec un tie break

					if errTB == nil {
						resSWF = append(resSWF, gagnant)
						delete(countSWF, gagnant)

					} else {
						return nil, errTB
					}

				} else if len(bestAlts) == 1 { //
					resSWF = append(resSWF, bestAlts[0])
					delete(countSWF, bestAlts[0])
				}
				i++
			}
			bestAlts := MaxCount(countSWF)
			resSWF = append(resSWF, bestAlts[0])
			return resSWF, nil

		} else {
			return nil, errSWF
		}
	}
}
