package operadores

import "fmt"

func OperadoresAritmeticos() {
	fmt.Println("Operadores aritmeticos")
}

func OperadoresRelacionais() {
	fmt.Println("Operadores relacionais")
}

func OperadoresLogicos() {
	fmt.Println("\n=== OPERADOR && (AND) - Todas as combinações ===")

	// COMBINAÇÃO 1: true && true = true
	fmt.Println("\n--- COMBINAÇÃO 1: true && true ---")
	num1 := 7
	num2 := 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: true && true = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if num1 > num2 && num1 < 10 {
		fmt.Println("✅ num1 é maior que num2 E menor que 10")
	}

	// COMBINAÇÃO 2: true && false = false
	fmt.Println("\n--- COMBINAÇÃO 2: true && false ---")
	num1 = 10
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: true && false = false")
	fmt.Println("RESULTADO: NÃO executa o bloco if")
	if num1 > num2 && num1 < 10 {
		fmt.Println("Esta mensagem não aparece")
	} else {
		fmt.Println("❌ num1 é maior que num2 MAS NÃO é menor que 10")
	}

	// COMBINAÇÃO 3: false && true = false (curto-circuito)
	fmt.Println("\n--- COMBINAÇÃO 3: false && true (CURTO-CIRCUITO) ---")
	num1 = 3
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é false, então NÃO avalia BLOCO 2 (CURTO-CIRCUITO)")
	fmt.Println("         O compilador para aqui e não verifica num1 < 10")
	fmt.Println("PASSO 3: false && (não avaliado) = false")
	fmt.Println("RESULTADO: NÃO executa o bloco if")
	if num1 > num2 && num1 < 10 {
		fmt.Println("Esta mensagem não aparece")
	} else {
		fmt.Println("❌ num1 NÃO é maior que num2 (BLOCO 2 não foi avaliado)")
	}

	// COMBINAÇÃO 4: false && false = false (curto-circuito)
	fmt.Println("\n--- COMBINAÇÃO 4: false && false (CURTO-CIRCUITO) ---")
	num1 = 3
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é false, então NÃO avalia BLOCO 2 (CURTO-CIRCUITO)")
	fmt.Println("         O compilador para aqui e não verifica num1 < 10")
	fmt.Println("PASSO 3: false && (não avaliado) = false")
	fmt.Println("RESULTADO: NÃO executa o bloco if")
	if num1 > num2 && num1 < 10 {
		fmt.Println("Esta mensagem não aparece")
	} else {
		fmt.Println("❌ num1 NÃO é maior que num2 (BLOCO 2 não foi avaliado)")
	}

	fmt.Println("\n=== OPERADOR || (OR) - Todas as combinações ===")

	// COMBINAÇÃO 1: true || true = true (curto-circuito)
	fmt.Println("\n--- COMBINAÇÃO 1: true || true (CURTO-CIRCUITO) ---")
	num1 = 7
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então NÃO avalia BLOCO 2 (CURTO-CIRCUITO)")
	fmt.Println("         O compilador para aqui e não verifica num1 < 10")
	fmt.Println("PASSO 3: true || (não avaliado) = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if num1 > num2 || num1 < 10 {
		fmt.Println("✅ num1 é maior que num2 OU menor que 10 (BLOCO 2 não foi avaliado)")
	}

	// COMBINAÇÃO 2: true || false = true (curto-circuito)
	fmt.Println("\n--- COMBINAÇÃO 2: true || false (CURTO-CIRCUITO) ---")
	num1 = 10
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então NÃO avalia BLOCO 2 (CURTO-CIRCUITO)")
	fmt.Println("         O compilador para aqui e não verifica num1 < 10")
	fmt.Println("PASSO 3: true || (não avaliado) = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if num1 > num2 || num1 < 10 {
		fmt.Println("✅ num1 é maior que num2 (BLOCO 2 não foi avaliado)")
	}

	// COMBINAÇÃO 3: false || true = true
	fmt.Println("\n--- COMBINAÇÃO 3: false || true ---")
	num1 = 3
	num2 = 5
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é false, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: false || true = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if num1 > num2 || num1 < 10 {
		fmt.Println("✅ num1 NÃO é maior que num2 MAS é menor que 10")
	}

	// COMBINAÇÃO 4: false || false = false
	fmt.Println("\n--- COMBINAÇÃO 4: false || false ---")
	num1 = 15
	num2 = 20
	fmt.Printf("num1 = %d, num2 = %d\n", num1, num2)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é false, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: false || false = false")
	fmt.Println("RESULTADO: NÃO executa o bloco if")
	if num1 > num2 || num1 < 10 {
		fmt.Println("Esta mensagem não aparece")
	} else {
		fmt.Println("❌ num1 NÃO é maior que num2 E NÃO é menor que 10")
	}
}

func OperadoresLogicosTresCombinacoes() {
	fmt.Println("\n=== OPERADORES LÓGICOS COM 3 COMBINAÇÕES ===")

	// COMBINAÇÃO 1: (true && true) && true = true
	fmt.Println("\n--- COMBINAÇÃO 1: (BLOCO1 && BLOCO2) && BLOCO3 = true ---")
	num1 := 7
	num2 := 5
	num3 := 3
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d\n", num1, num2, num3)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: BLOCO 1 && BLOCO 2 = true && true = true")
	fmt.Println("PASSO 4: (true) && BLOCO 3 -> num1 > num3")
	fmt.Printf("         %d > %d = %v\n", num1, num3, num1 > num3)
	fmt.Println("PASSO 5: true && true = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if (num1 > num2 && num1 < 10) && num1 > num3 {
		fmt.Println("✅ Todas as condições são verdadeiras: num1 > num2 E num1 < 10 E num1 > num3")
	}

	// COMBINAÇÃO 2: (true && false) && true = false
	fmt.Println("\n--- COMBINAÇÃO 2: (BLOCO1 && BLOCO2) && BLOCO3 = false ---")
	num1 = 10
	num2 = 5
	num3 = 3
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d\n", num1, num2, num3)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então avalia BLOCO 2 -> num1 < 10")
	fmt.Printf("         %d < 10 = %v\n", num1, num1 < 10)
	fmt.Println("PASSO 3: BLOCO 1 && BLOCO 2 = true && false = false")
	fmt.Println("PASSO 4: (false) && BLOCO 3 -> NÃO avalia BLOCO 3 (CURTO-CIRCUITO)")
	fmt.Println("         O compilador para aqui e não verifica num1 > num3")
	fmt.Println("PASSO 5: false && (não avaliado) = false")
	fmt.Println("RESULTADO: NÃO executa o bloco if")
	if (num1 > num2 && num1 < 10) && num1 > num3 {
		fmt.Println("Esta mensagem não aparece")
	} else {
		fmt.Println("❌ BLOCO 1 && BLOCO 2 resultou em false, então BLOCO 3 não foi avaliado")
	}

	// COMBINAÇÃO 3: (true || false) && true = true
	fmt.Println("\n--- COMBINAÇÃO 3: (BLOCO1 || BLOCO2) && BLOCO3 = true ---")
	num1 = 10
	num2 = 5
	num3 = 3
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d\n", num1, num2, num3)
	fmt.Println("PASSO 1: Avalia BLOCO 1 -> num1 > num2")
	fmt.Printf("         %d > %d = %v\n", num1, num2, num1 > num2)
	fmt.Println("PASSO 2: BLOCO 1 é true, então NÃO avalia BLOCO 2 (CURTO-CIRCUITO do ||)")
	fmt.Println("         O compilador para aqui e não verifica num1 < 10")
	fmt.Println("PASSO 3: BLOCO 1 || (não avaliado) = true || (não avaliado) = true")
	fmt.Println("PASSO 4: (true) && BLOCO 3 -> num1 > num3")
	fmt.Printf("         %d > %d = %v\n", num1, num3, num1 > num3)
	fmt.Println("PASSO 5: true && true = true")
	fmt.Println("RESULTADO: Executa o bloco if")
	if (num1 > num2 || num1 < 10) && num1 > num3 {
		fmt.Println("✅ (num1 > num2 OU num1 < 10) E num1 > num3 - BLOCO 2 não foi avaliado")
	}

}

