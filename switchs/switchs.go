package switchs

import "fmt"

func Switch() {
	var numero int = 12

	switch numero {
	case 1:
		fmt.Println("O numero é 1")
	case 2:
		fmt.Println("O numero é 2")
	case 3:
		fmt.Println("O numero é 3")
	default:
		fmt.Println("O numero é diferente de 1, 2 e 3")
	}
}

func SwitchComRetorno() string {
	var numero int = 12

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func SwitcComAtribuicaoDeVariavel(numero int) string {
	var diaDaSemana string 

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia inválido"
	}
	return diaDaSemana
}