package main

func SomaT[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaGeneric() {
	m := map[string]int{"Wesley": 1000, "Jhonathan": 2000, "Woody": 3000}
	println(SomaT(m))

	m2 := map[string]float64{"Wesley": 100.10, "Jhonathan": 200.20, "Woody": 300.30}
	println(SomaT(m2))
}
