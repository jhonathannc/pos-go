package main

type MyNumber int

type Number interface {
	~int | float64
}

func SomaC[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaConstraint() {
	m := map[string]int{"Wesley": 1000, "Jhonathan": 2000, "Woody": 3000}
	println(SomaC(m))

	m2 := map[string]float64{"Wesley": 100.10, "Jhonathan": 200.20, "Woody": 300.30}
	println(SomaC(m2))

	m3 := map[string]MyNumber{"Wesley": 1000, "Jhonathan": 2000, "Woody": 3000}
	println(SomaC(m3))
}
