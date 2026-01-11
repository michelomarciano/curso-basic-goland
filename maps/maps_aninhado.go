package maps

import "fmt"

func MapAninhado() {
	dados := map[string]map[string]string{
		"nome": {
			"nome": "Mike",
			"idade": "30",
			"email": "mike@example.com",
		},
		"endereco": {
			"rua": "Rua das Flores",
			"numero": "123",
			"cidade": "São Paulo",
			"estado": "SP",
			"cep": "1234567890",
		},
	}
	fmt.Println(dados)
	fmt.Println(dados["nome"]["nome"])
	fmt.Println(dados["endereco"]["rua"])
}
