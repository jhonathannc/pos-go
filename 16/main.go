package main

func main() {
	var minhaVar interface{} = "Jhonathan Candeu"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	println("O valor de res é %v e o valor de ok é %v", res, ok)

	res2 := minhaVar.(int)
	println("O valor de res2 é %v", res2)
}
