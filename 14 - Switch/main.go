package main

import "fmt"

//verificação explicita
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough 
		//essa clausula avalia a condição e 
		// joga ele pra dentro da proxima condição
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Numero inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("14 - Switches")
	//dia := diaDaSemana(3)
	//dia2 := diaDaSemana2(3)
	dia := diaDaSemana2(1)
	fmt.Println(dia)
	//fmt.Println(dia2)
}
