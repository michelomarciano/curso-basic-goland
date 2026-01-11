package heranca

import "fmt"

type Animal struct {
	Nome  string
	Idade int
	Cor   string
	Peso  float64
}

type Cachorro struct {
	Animal
	Raça string
}

type Gato struct {
	Animal
	Raça string
}

func Heranca() {
	fmt.Println("Herança")
}
