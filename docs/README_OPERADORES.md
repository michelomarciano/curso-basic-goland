# OPERADORES

## ARITMÉTICOS

**IMPORTANTE**: Não podemos somar um inteiro de 32 bits com outro de 64 bits, por exemplo. Isso acontece porque Go é uma linguagem fortemente tipada e requer conversão explícita entre tipos diferentes. 

| Operador | Nome | Descrição | Exemplo |
|----------|------|-----------|---------|
| `+` | Adição | Soma dois valores | `x := 5 + 3` // x = 8 |
| `-` | Subtração | Subtrai dois valores | `x := 5 - 3` // x = 2 |
| `*` | Multiplicação | Multiplica dois valores | `x := 5 * 3` // x = 15 |
| `/` | Divisão | Divide dois valores | `x := 10 / 2` // x = 5 |
| `%` | Módulo | Retorna o resto da divisão | `x := 10 % 3` // x = 1 |
| `++` | Incremento | Incrementa o valor em 1 | `x++` // x = x + 1 |
| `--` | Decremento | Decrementa o valor em 1 | `x--` // x = x - 1 |

**OBSERVAÇÃO**: Em Go, `++` e `--` são apenas operadores de pós-incremento/decremento. Não podem ser usados em expressões, apenas como instruções separadas. 

**Exemplo:** `x++` é válido, mas `y := x++` não é válido.

### Casos de Uso - Aritméticos
- **Cálculos matemáticos**: Operações básicas em calculadoras, sistemas financeiros
- **Contadores e loops**: Incrementar/decrementar variáveis em iterações
- **Processamento de dados**: Somar totais, calcular médias, estatísticas
- **Jogos**: Pontuação, vidas, recursos
- **Sistemas de medição**: Converter unidades, calcular distâncias, áreas

**Exemplo prático:**
```go
// Calcular total de vendas
total := 0
for _, venda := range vendas {
    total += venda.Valor  // Adição
}
media := total / len(vendas)  // Divisão
```

## ATRIBUIÇÃO COMPOSTA

Operadores de atribuição composta combinam uma operação aritmética com uma atribuição. Eles são uma forma abreviada e mais concisa de modificar o valor de uma variável.

**Sintaxe Geral:**
```go
variavel operador= valor
// Equivale a:
variavel = variavel operador valor
```

| Operador | Nome | Equivalência | Exemplo |
|----------|------|--------------|---------|
| `+=` | Adição e atribuição | `x += y` é igual a `x = x + y` | `total += 10` // total = total + 10 |
| `-=` | Subtração e atribuição | `x -= y` é igual a `x = x - y` | `saldo -= 5` // saldo = saldo - 5 |
| `*=` | Multiplicação e atribuição | `x *= y` é igual a `x = x * y` | `preco *= 1.1` // preco = preco * 1.1 |
| `/=` | Divisão e atribuição | `x /= y` é igual a `x = x / y` | `media /= 2` // media = media / 2 |
| `%=` | Módulo e atribuição | `x %= y` é igual a `x = x % y` | `resto %= 3` // resto = resto % 3 |

**Exemplo Básico - Operador `+=`:**
```go
total := 0
total += 5   // total = 0 + 5 = 5
total += 10  // total = 5 + 10 = 15
total += 3   // total = 15 + 3 = 18
fmt.Println(total) // 18
```

**O que acontece com `+=`:**
O operador `+=` pega o valor atual da variável, soma com o valor especificado e guarda o resultado de volta na mesma variável.

**Passo a passo:**
```go
total := 0
numero := 5

total += numero
// 1. Lê o valor atual de total: 0
// 2. Lê o valor de numero: 5
// 3. Soma: 0 + 5 = 5
// 4. Atribui o resultado de volta em total: total = 5
```

**Exemplo Prático - Acumulador em Loop:**
```go
func SomarNumeros(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero  // Adiciona cada número ao total
    }
    return total
}

// Uso:
soma := SomarNumeros(1, 2, 3, 4, 5)
fmt.Println(soma) // 15

// Execução passo a passo:
// total = 0
// total += 1  → total = 1
// total += 2  → total = 3
// total += 3  → total = 6
// total += 4  → total = 10
// total += 5  → total = 15
```

**Exemplos com Outros Operadores:**
```go
// Subtração
saldo := 100
saldo -= 20  // saldo = 100 - 20 = 80
saldo -= 10  // saldo = 80 - 10 = 70

// Multiplicação
preco := 10.0
preco *= 1.5  // preco = 10.0 * 1.5 = 15.0 (aumento de 50%)

// Divisão
media := 100.0
media /= 4  // media = 100.0 / 4 = 25.0

// Módulo
resto := 10
resto %= 3  // resto = 10 % 3 = 1
```

**Exemplo com Strings (Concatenação):**
```go
texto := "Hello"
texto += " "      // texto = "Hello "
texto += "World"  // texto = "Hello World"
fmt.Println(texto) // "Hello World"
```

**Vantagens:**
- **Código mais conciso**: `total += numero` é mais curto que `total = total + numero`
- **Mais legível**: Deixa claro que estamos acumulando/modificando um valor
- **Menos propenso a erros**: Evita repetir o nome da variável
- **Idiomático em Go**: É a forma preferida para modificar valores

**OBSERVAÇÕES:**
- Os operadores de atribuição composta funcionam com todos os tipos numéricos (int, float, etc.)
- Para strings, apenas `+=` está disponível (concatenação)
- A variável deve estar declarada antes de usar operadores de atribuição composta
- Não pode ser usado em expressões, apenas como instrução separada

**Exemplo Avançado - Múltiplas Operações:**
```go
// Calcular estatísticas
func CalcularEstatisticas(numeros []int) (soma, produto int) {
    soma = 0
    produto = 1
    
    for _, num := range numeros {
        soma += num      // Acumula a soma
        produto *= num    // Acumula o produto
    }
    
    return
}

// Uso:
soma, produto := CalcularEstatisticas([]int{2, 3, 4})
fmt.Println(soma)    // 9 (2+3+4)
fmt.Println(produto) // 24 (2*3*4)
```

### Casos de Uso - Atribuição Composta
- **Acumuladores**: Somar valores em loops (contadores, totais, médias)
- **Contadores**: Incrementar/decrementar valores
- **Cálculos iterativos**: Processar dados em múltiplas iterações
- **Concatenação de strings**: Juntar múltiplas strings
- **Ajustes de valores**: Aplicar descontos, aumentos, reduções percentuais
- **Processamento de dados**: Transformar valores em coleções

**Exemplo prático:**
```go
// Acumular pontos em um jogo
pontos := 0
pontos += 10  // Matou inimigo
pontos += 5   // Coletou item
pontos += 20  // Completou missão
fmt.Println("Pontos totais:", pontos) // 35

// Aplicar desconto progressivo
preco := 100.0
preco *= 0.9   // 10% de desconto: 90.0
preco *= 0.95  // Mais 5% de desconto: 85.5

// Construir mensagem
mensagem := "Erro: "
mensagem += "Falha na conexão"
mensagem += " - Tente novamente"
fmt.Println(mensagem) // "Erro: Falha na conexão - Tente novamente"
```

## RELACIONAIS

| Operador | Nome | Descrição | Exemplo |
|----------|------|-----------|---------|
| `==` | Igual a | Verifica se dois valores são iguais | `x == y` // `true` se x for igual a y |
| `!=` | Diferente de | Verifica se dois valores são diferentes | `x != y` // `true` se x for diferente de y |
| `<` | Menor que | Verifica se o valor da esquerda é menor que o da direita | `x < y` // `true` se x for menor que y |
| `<=` | Menor ou igual a | Verifica se o valor da esquerda é menor ou igual ao da direita | `x <= y` // `true` se x for menor ou igual a y |
| `>` | Maior que | Verifica se o valor da esquerda é maior que o da direita | `x > y` // `true` se x for maior que y |
| `>=` | Maior ou igual a | Verifica se o valor da esquerda é maior ou igual ao da direita | `x >= y` // `true` se x for maior ou igual a y |

**OBSERVAÇÃO**: Os operadores relacionais retornam um valor booleano (`true` ou `false`).

### Casos de Uso - Relacionais
- **Validações**: Verificar se valores estão dentro de limites aceitáveis
- **Comparações**: Comparar dados, ordenar elementos
- **Condicionais**: Tomar decisões baseadas em comparações
- **Filtros**: Filtrar dados baseado em condições
- **Sistemas de permissão**: Verificar se usuário tem idade suficiente, permissões adequadas

**Exemplo prático:**
```go
// Validar idade para acesso
if idade >= 18 {
    fmt.Println("Acesso permitido")
}

// Verificar se temperatura está em faixa segura
if temp >= 20 && temp <= 30 {
    fmt.Println("Temperatura OK")
}
```

## LÓGICOS

| Operador | Nome | Descrição | Exemplo |
|----------|------|-----------|---------|
| `&&` | E (AND) | Retorna `true` se ambas as condições forem verdadeiras | `x > 5 && x < 10` // `true` se x estiver entre 5 e 10 |
| `\|\|` | OU (OR) | Retorna `true` se pelo menos uma das condições for verdadeira | `x < 5 \|\| x > 10` // `true` se x for menor que 5 ou maior que 10 |
| `!` | NÃO (NOT) | Inverte o valor booleano | `!true` // `false`, `!false` // `true` |

**OBSERVAÇÃO**: Os operadores lógicos trabalham com valores booleanos e retornam um valor booleano. 

- O operador `&&` usa **avaliação de curto-circuito**: se a primeira condição for `false`, a segunda não é avaliada.
- O operador `||` também usa **curto-circuito**: se a primeira condição for `true`, a segunda não é avaliada.

### Casos de Uso - Lógicos
- **Validações complexas**: Combinar múltiplas condições (ex: email válido E senha forte)
- **Controle de fluxo**: Decisões baseadas em múltiplas condições
- **Filtros avançados**: Buscar dados que atendam várias condições simultaneamente
- **Lógica de negócio**: Implementar regras de negócio complexas
- **Tratamento de erros**: Verificar múltiplas condições de erro

**Exemplo prático:**
```go
// Validar login
if email != "" && senha != "" && len(senha) >= 8 {
    // Processar login
}

// Verificar permissões
if isAdmin || (isUser && temPermissao) {
    // Permitir acesso
}
```

## UNÁRIOS

| Operador | Nome | Descrição | Exemplo |
|----------|------|-----------|---------|
| `+` | Positivo | Mantém o valor positivo (geralmente implícito) | `x := +5` // x = 5 |
| `-` | Negativo | Inverte o sinal do valor | `x := -5` // x = -5 |
| `!` | Negação Lógica | Inverte o valor booleano | `!true` // `false` |
| `^` | Negação Bitwise | Inverte os bits (complemento de 1) | `^5` // inverte todos os bits de 5 |
| `*` | Dereferência | Obtém o valor apontado por um ponteiro | `*ptr` // valor apontado por ptr |
| `&` | Endereço | Obtém o endereço de memória de uma variável | `&x` // endereço de memória de x |
| `++` | Incremento | Incrementa o valor em 1 (pós-incremento) | `x++` // x = x + 1 |
| `--` | Decremento | Decrementa o valor em 1 (pós-decremento) | `x--` // x = x - 1 |
| `<-` | Receive | Recebe um valor de um canal | `valor := <-ch` // recebe valor do canal ch |

**OBSERVAÇÃO**: 
- Os operadores `++` e `--` são apenas pós-incremento/decremento em Go. Não podem ser usados em expressões, apenas como instruções separadas.
- O operador `<-` é usado para comunicação entre goroutines através de canais.

### Casos de Uso - Unários
- **Ponteiros**: Trabalhar com endereços de memória, modificar valores por referência
- **Canais**: Comunicação entre goroutines (programação concorrente)
- **Operações bitwise**: Manipulação de bits para otimização, flags, máscaras
- **Negação lógica**: Inverter condições booleanas
- **Operações matemáticas**: Inverter sinais, operações unárias

**Exemplo prático:**
```go
// Usar ponteiro para modificar valor
valor := 10
ptr := &valor  // Obter endereço
*ptr = 20      // Modificar através do ponteiro

// Negação lógica
if !estaLogado {
    redirecionarParaLogin()
}

// Comunicação via canal
resultado := <-canal  // Receber valor do canal
```

## TERNÁRIO

**OBSERVAÇÃO**: Não existe operador ternário em Go como `idade := numero >= 18 ? "Maior de idade" : "Menor de idade"`. O que podemos fazer é usar `if else` para resolver.

**Exemplo:**
```go
var idade string
if numero >= 18 {
    idade = "Maior de idade"
} else {
    idade = "Menor de idade"
}
```

### Casos de Uso - Ternário (Simulado)
- **Atribuições condicionais**: Quando você precisa atribuir um valor baseado em uma condição
- **Valores padrão**: Definir valores padrão quando uma condição não é atendida
- **Formatação condicional**: Escolher formatos diferentes baseado em condições
- **Mensagens dinâmicas**: Gerar mensagens diferentes baseado em estados

**Exemplo prático:**
```go
// Definir status baseado em idade
var status string
if idade >= 18 {
    status = "Adulto"
} else {
    status = "Menor"
}

// Definir desconto baseado em tipo de cliente
var desconto float64
if isVIP {
    desconto = 0.20
} else {
    desconto = 0.10
}
```

