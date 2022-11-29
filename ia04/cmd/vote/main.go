package main

func main() {
	// alt := []comsoc.Alternative{4, 2, 6, 1, 5, 3}
	// var alt1 comsoc.Alternative = 5
	// var alt2 comsoc.Alternative = 2

	// c := make(comsoc.Count)

	// c[1] = 13
	// c[2] = 12
	// c[3] = 34
	// c[4] = 61
	// c[5] = 61
	// c[6] = 61
	// c[7] = 61
	// c[7] = 11

	// p := comsoc.Profile{{1, 2, 5, 12},
	// 	{2, 12, 1, 5},
	// 	{2, 5, 12, 1},
	// 	{12, 1, 5, 2},
	// 	{5, 12, 1, 2}}

	// p2 := comsoc.Profile{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4},
	// 	{2, 3, 4, 1}, {2, 3, 4, 1}, {2, 3, 4, 1}, {2, 3, 4, 1},
	// 	{4, 3, 1, 2}, {4, 3, 1, 2}, {4, 3, 1, 2}}

	// // Test Copeland_SWF
	// resCop, errCop := comsoc.CopelandSWF(p)
	// if errCop == nil {
	// 	fmt.Println(resCop)
	// } else {
	// 	fmt.Println(errCop)
	// }

	// // Test Copeland_SCF
	// resCop_SCF, errCop_SCF := comsoc.CopelandSCF(p)
	// if errCop_SCF == nil {
	// 	fmt.Println(resCop_SCF)
	// } else {
	// 	fmt.Println(errCop_SCF)
	// }

	// Test SVT_SWF
	// resSTV_SWF, errSTV_SWF := comsoc.STV_SWF(p)

	// if errSTV_SWF == nil {
	// 	fmt.Println(resSTV_SWF)
	// } else {
	// 	fmt.Println(errSTV_SWF)
	// }

	// Test SVT_SCF
	// resSTV_SCF, errSTV_SCF := comsoc.STV_SCF(p)
	// if errSTV_SCF == nil {
	// 	fmt.Println(resSTV_SCF)
	// } else {
	// 	fmt.Println(errSTV_SCF)
	// }

	// giveRanking := comsoc.SWFFactory(comsoc.STV_SWF, comsoc.TieBreakFactory([]comsoc.Alternative{}))
	// ranking, errRanking := giveRanking(p)
	// if errRanking == nil {
	// 	fmt.Println(ranking)
	// } else {
	// 	fmt.Println(errRanking)
	// }

	// Test ApprovalSWF
	// thresholds1 := []int{2, 3, 4, 2, 1}

	// resApprovalSWF, errSWF := comsoc.ApprovalSWF(p, thresholds1)
	// if errSWF == nil {
	// 	fmt.Println(resApprovalSWF)
	// } else {
	// 	fmt.Println(errSWF)
	// }

	// resApprovalSCF, err := comsoc.ApprovalSCF(p, thresholds1)
	// if err == nil {
	// 	fmt.Println(resApprovalSCF)
	// } else {
	// 	fmt.Println(err)
	// }

	// // Test BordaSWF
	// countBordaSWF, errBordaSWF := comsoc.BordaSWF(p)

	// if errBordaSWF == nil {
	// 	fmt.Println(countBordaSWF)
	// } else {
	// 	fmt.Println(errBordaSWF)
	// }

	// // Test SWFFactory
	// bestAlts := []comsoc.Alternative{}
	// SWFsansExAequo := comsoc.SWFFactory(comsoc.BordaSWF, comsoc.TieBreakFactory(bestAlts))
	// classement, err := SWFsansExAequo(p)
	// if err == nil {
	// 	fmt.Println(classement)
	// } else {
	// 	fmt.Println(err)
	// }

	// // Test BordaSCF
	// resBordaSCF, errBorda := comsoc.BordaSCF(p)

	// if errBorda == nil {
	// 	fmt.Println(resBordaSCF)
	// } else {
	// 	fmt.Println(errBorda)
	// }

	// // Test du TieBreakFactory sur BordaSCF
	// tiebreak := comsoc.TieBreakFactory(resBordaSCF)
	// result, errTieBreak := tiebreak(resBordaSCF)

	// if errTieBreak == nil {
	// 	fmt.Println(result)
	// } else {
	// 	fmt.Println(errTieBreak)
	// }

	// Test SCFFactory
	// resSCF := []comsoc.Alternative{}
	// SCFsansExAequo := comsoc.SCFFactory(comsoc.BordaSCF, comsoc.TieBreakFactory(resSCF))
	// gagnant, err := SCFsansExAequo(p)
	// if err == nil {
	// 	fmt.Println(gagnant)
	// } else {
	// 	fmt.Println(err)
	// }

	// Test MajoritySWF
	// count := make(comsoc.Count)
	// count, err := comsoc.MajoritySWF(p)

	// if err == nil {
	// 	fmt.Println(count)
	// } else {
	// 	fmt.Println(err)
	// }

	// Test MajoritySCF
	// bestAlts, err2 := comsoc.MajoritySCF(p, 1)

	// if err2 == nil {
	// 	fmt.Println(bestAlts)
	// } else {
	// 	fmt.Println(err2)
	// }

	// Test Rank
	// fmt.Println(comsoc.Rank(1, p[1]))
	// fmt.Println(comsoc.Rank(2, p[1]))

	// Test IsPref
	// fmt.Println(comsoc.IsPref(1, 2, p[1]))

	// Test MaxCount
	// tab := comsoc.MaxCount(c, 3)
	// fmt.Println(tab)

	// Test CheckProfile
	// err := comsoc.CheckProfile(p)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Profile ok")
	// }

	// Test CheckProfileAlternative(p, alt)
	// err := comsoc.CheckProfileAlternative(p, alt)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Profile ok")
	// }

	// ----------------------- PARTIE TEST CONDORCET ----------------------- //

	// var alt comsoc.Alternative = 3
	// var preference1 = []comsoc.Alternative{5, 2, 1, 3, 6, 4}

	// fmt.Println("alt =", alt, "dans", preference1, "?")
	// fmt.Println("reponse : ", comsoc.CheckAlternativeDansTableau(alt, preference1))

	// fmt.Println("alt =", alt, "index = ", comsoc.Rank(alt, preference1))
	// fmt.Println("5 > 2 oui", comsoc.IsPref(5, 2, preference1))

	// //var m map[comsoc.Alternative] int
	// m := make(map[comsoc.Alternative]int)
	// m[1] = 14
	// m[2] = 12
	// m[3] = 17
	// m[4] = 16
	// m[5] = 18
	// m[6] = 4

	// fmt.Println("meilleures alternatives ", comsoc.MaxCount(m))

	// p := comsoc.Profile{{1, 2, 3}, {2, 3, 1}, {2, 1, 3}}

	// fmt.Println("erreur ? ", comsoc.CheckProfile(p), " fin de l'erreur ")

	// p2 := comsoc.Profile{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4},
	// 	{2, 3, 4, 1}, {2, 3, 4, 1}, {2, 3, 4, 1}, {2, 3, 4, 1},
	// 	{4, 3, 1, 2}, {4, 3, 1, 2}, {4, 3, 1, 2}}

	// p1 := comsoc.Profile{{1, 2, 3},
	// 	{2, 3, 1},
	// 	{3, 1, 2}}

	// res, err := comsoc.CondorcetWinner(p1)
	// if err == nil {
	// 	fmt.Println(res)
	// } else {
	// 	fmt.Println(err)
	// }
}
