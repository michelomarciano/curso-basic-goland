package main

import(
	"fmt"
	m "math" // incluido alias
)

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)
	fmt.Print("cost", a,b)
	

}