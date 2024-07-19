package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Nome  string `json:"nome"`
	Saldo int    `json:"saldo"`
}

func main() {
	conta := Conta{Nome: "Jhonathan", Saldo: 1000}
	res, err := json.Marshal(conta) // serializa a struct para um json
	if err != nil {
		println(err)
	}

	// println(res) // por padrao o marshal retorna a serializacao em bytes
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	/**
	\ o marshal codifica em json e armazena em um variavel
	\ o encoder codifica e retorna direto em algum ponto, no exemplo acima ele
	\ joga em um terminal os.Stdout
	*/

	jsonPuro := []byte(`{"n": "woody", "s": 200}`)
	var contax Conta
	err = json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		println(err)
	}
	println(contax.Saldo)
}
