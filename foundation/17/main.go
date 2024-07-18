package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Wesley": 1000, "Jhonathan": 2000, "Woody": 3000}
	println(SomaInteiro(m))

	m2 := map[string]float64{"Wesley": 100.10, "Jhonathan": 200.20, "Woody": 300.30}
	println(SomaFloat(m2))

	SomaGeneric()
}
