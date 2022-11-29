package comsoc

import (
	"errors"
	"fmt"
)

const MAXIMUM_POINTS int = 10000
const NONE Alternative = 0

func Rank(alt Alternative, prefs []Alternative) int {
	// renvoie l'indice ou se trouve alt dans prefs
	var indice int = 0
	for i := 0; i < len(prefs); i++ {
		if prefs[i] == alt {
			indice = i
		}
	}
	return indice
}

func IsPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	// renvoie vrai ssi alt1 est préférée à alt2
	return Rank(alt1, prefs) < Rank(alt2, prefs)
}

func MaxCount(count Count) (bestAlts []Alternative) { // renvoie les meilleures alternatives pour un décompte donné
	max := -1
	for _, element := range count {
		if max < element {
			max = element
		}
	}
	for key, element := range count {
		if element == max {
			bestAlts = append(bestAlts, key)
		}
	}
	return bestAlts
}

func MaxCountx(count Count, param int) (bestAlts []Alternative) {
	// renvoie les meilleures alternatives pour un décompte donné
	copycount := make(Count)
	for key, value := range count {
		copycount[key] = value
	}

	for param-len(bestAlts) > 0 {
		max := -1
		for _, element := range count {
			if max < element {
				max = element
			}
		}
		for key, element := range count {
			if element == max {
				bestAlts = append(bestAlts, key)
				delete(copycount, key)
			}
		}
	}
	return bestAlts
}

func CheckProfile(prefs Profile) error {
	// vérifie le profil donné, par ex. qu'ils sont tous complets
	// et que chaque alternative n'apparaît qu'une seule fois par préférences

	for i := range prefs {

		if i > 0 && len(prefs[i-1]) != len(prefs[i]) {
			return errors.New("Erreur : profil incomplet ou trop long")
		}

		m := make(map[Alternative]bool)

		for j := range prefs[i] {
			_, exists := m[prefs[i][j]] // check si l'alternative existe
			if exists {
				fmt.Printf("Agent %d | ", i+1)
				return errors.New("Erreur : doublon d'alternative")
			} else {
				m[prefs[i][j]] = true
			}
		}
	}
	return nil
}

func CheckProfileAlternative(prefs Profile, alt []Alternative) error {
	// vérifie le profil donné, par ex. qu'ils sont tous complets
	// et que chaque alternative de alts apparaît exactement une fois par préférences

	for i := range prefs {

		if len(prefs[i]) != len(alt) {
			fmt.Printf("Agent %d | ", i+1)
			return errors.New("Erreur : alternative(s) manquante(s) ou en trop")
		}

		m := make(map[Alternative]bool)
		for i := range alt {
			m[alt[i]] = false
		}

		for j := range prefs[i] {
			bool_value, exists := m[prefs[i][j]] // check si l'alternative existe
			if exists && !bool_value {
				m[prefs[i][j]] = true

			} else if exists && bool_value {
				fmt.Printf("Agent %d | ", i+1)
				return errors.New("Erreur : doublon d'alternative")

			} else if !exists {
				fmt.Printf("Agent %d | ", i+1)
				return errors.New("Erreur : alternative incorrecte")
			}
		}
	}
	return nil
}

// func Remove(slice []Alternative, a Alternative) []Alternative {
// 	for i := range slice {
// 		if slice[i] == a {
// 			return append(slice[:i], slice[i+1:]...)
// 		}
// 	}
// 	return nil
// }

// Fonctions utilitaires pour STV_SCF

func CountPointsSTV(p Profile) Count {
	count := make(Count)
	for i := range p[0] { // initialiser les points de chaque alternative à 0
		count[p[0][i]] = 0
	}

	for j := 0; j < len(p); j++ { // pour chaque votant
		count[p[j][0]]++ // on incrémente le score de l'alternative préférée
	}

	return count
}

func MinCount(count Count) (perdant Alternative, gagnantExaequo bool) { // renvoie les pires alternatives pour un décompte donné
	min := MAXIMUM_POINTS
	for _, element := range count {
		if element < min {
			min = element
		}
	}
	worstAlts := []Alternative{}
	for key, element := range count {
		if element == min {
			worstAlts = append(worstAlts, key)
		}
	}

	if len(worstAlts) > 1 && len(worstAlts) != len(count) { // ex aequo sur les perdants => Tie Break
		return TieBreakSmallest(worstAlts), false

	} else if len(worstAlts) > 1 && len(worstAlts) == len(count) { // ex aequo sur les gagnants
		return NONE, true

	} else { // un seul perdant
		return worstAlts[0], false
	}
}

func RetirerPerdant(p Profile, perdant Alternative) Profile {

	for i := range p {
		for j := 0; j < len(p[i]); j++ {
			if p[i][j] == perdant {
				p[i] = append(p[i][:j], p[i][j+1:]...) // on retire sur chaque ligne le perdant
				j = len(p[i])
			}
		}

		// p[j] = Remove(p[j], worstAlts[0])
	}
	return p

}
