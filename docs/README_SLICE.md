# SLICE

Slices são estruturas de dados dinâmicas em Go, muito mais utilizadas que arrays por serem flexíveis e redimensionáveis. Um slice é uma referência a um array subjacente e permite trabalhar com coleções de tamanho variável.

**OBSERVAÇÃO**: Ao contrário de arrays, slices não têm tamanho fixo e podem crescer ou diminuir dinamicamente.

## Diferença entre Array e Slice

| Característica | Array | Slice |
|----------------|-------|-------|
| Tamanho | Fixo (definido na declaração) | Dinâmico (pode crescer/diminuir) |
| Declaração | `[5]int` | `[]int` |
| Uso | Menos comum | Muito comum |
| Passagem | Por valor (cópia) | Por valor (cópia do slice, mas compartilha o array subjacente) |

## Criando Slices

### 1. Declaração e Inicialização
Cria um slice e inicializa com valores diretamente.

**Exemplo:**
```go
slice := []int{1, 2, 3, 4, 5}
fmt.Println(slice) // [1 2 3 4 5]
```

### 2. Slice Vazio e Slice Nil
É importante entender a diferença entre um slice nil e um slice vazio, pois eles têm comportamentos diferentes.

**Slice Nil:**
Um slice nil é um slice que não foi inicializado. Ele não aponta para nenhum array subjacente.

```go
var slice []int        // slice nil
// ou
var slice []int = nil // explicitamente nil
```

**Slice Vazio:**
Um slice vazio é um slice inicializado, mas sem elementos. Ele aponta para um array subjacente (mesmo que vazio).

```go
slice := []int{}       // slice vazio (não nil)
slice := make([]int, 0) // slice vazio usando make
```

**Comparação Prática:**
```go
var nilSlice []int           // nil
emptySlice := []int{}        // vazio
makeSlice := make([]int, 0)  // vazio

fmt.Println(nilSlice == nil)    // true
fmt.Println(emptySlice == nil)  // false
fmt.Println(makeSlice == nil)   // false

fmt.Println(len(nilSlice))      // 0
fmt.Println(len(emptySlice))    // 0
fmt.Println(len(makeSlice))    // 0

// Todos funcionam com append:
nilSlice = append(nilSlice, 1)    // OK - cria novo array
emptySlice = append(emptySlice, 1) // OK
makeSlice = append(makeSlice, 1)   // OK
```

**Características:**

| Característica | Slice Nil | Slice Vazio |
|----------------|-----------|-------------|
| `== nil` | `true` | `false` |
| `len()` | `0` | `0` |
| `cap()` | `0` | `0` |
| Array subjacente | Não existe | Existe (vazio) |
| `append()` funciona? | Sim | Sim |
| Uso comum | Indicar não inicializado | Indicar inicializado mas vazio |

**Quando Usar Cada Um:**

**Slice Nil:**
- Quando você quer indicar explicitamente que o slice não foi inicializado
- Útil para verificar se um slice foi inicializado: `if slice == nil { ... }`
- Padrão quando você declara uma variável slice sem inicializar

**Slice Vazio:**
- Quando você quer garantir que o slice não é `nil`
- Útil quando você precisa diferenciar entre "não inicializado" e "vazio mas inicializado"
- Quando você quer passar um slice vazio para uma função que pode verificar `nil`

**Exemplo Prático:**
```go
func processarSlice(slice []int) {
    if slice == nil {
        fmt.Println("Slice não foi inicializado")
        return
    }
    
    if len(slice) == 0 {
        fmt.Println("Slice está vazio mas foi inicializado")
    }
}

var nilSlice []int
emptySlice := []int{}

processarSlice(nilSlice)    // "Slice não foi inicializado"
processarSlice(emptySlice)  // "Slice está vazio mas foi inicializado"
```

**OBSERVAÇÃO**: Em muitos casos, o comportamento é o mesmo (ambos têm tamanho 0 e funcionam com `append`). A diferença principal é semântica: nil indica "não inicializado" e vazio indica "inicializado mas sem elementos".

### 3. Criar Slice com `make`
A função `make` permite criar slices com tamanho e capacidade inicial especificados. Isso é útil para otimização de performance. A função `make` recebe 3 parâmetros (tipo, tamanho, capacidade). A função `make` cria um array interno e retorna um slice de acordo com o tamanho das posições. Quando o slice atinge a capacidade máxima, o Go cria mais posições e duplica o tamanho da capacidade automaticamente.

**Sintaxe:**
```go
make([]tipo, tamanho, capacidade)
```

**Exemplos:**
```go
// Slice vazio (tamanho 0, capacidade 0)
slice1 := make([]int, 0)

// Slice vazio com capacidade inicial (tamanho 0, capacidade 10)
slice2 := make([]int, 0, 10)

// Slice pré-preenchido com zeros (tamanho 5, capacidade 5)
slice3 := make([]int, 5)  // [0 0 0 0 0]

// Slice com tamanho e capacidade diferentes (tamanho 3, capacidade 10)
slice4 := make([]int, 3, 10)  // [0 0 0]
```

**Por que usar `make`?**

1. **Especificar Capacidade Inicial (Performance)**
   A principal vantagem é definir a capacidade inicial, evitando realocações quando você adiciona elementos com `append`.

   **Sem `make` (pode ser mais lento):**
   ```go
   slice := []int{}  // capacidade 0
   // Cada append pode causar realocação de memória
   slice = append(slice, 1)  // realoca
   slice = append(slice, 2)  // realoca
   slice = append(slice, 3)  // realoca
   // ... muitas realocações
   ```

   **Com `make` (mais eficiente):**
   ```go
   slice := make([]int, 0, 100)  // capacidade 100
   // Os primeiros 100 appends não causam realocação
   slice = append(slice, 1)  // sem realocação
   slice = append(slice, 2)  // sem realocação
   // ... até 100 elementos sem realocação
   ```

2. **Criar Slice Pré-preenchido com Zeros**
   Útil quando você precisa de um slice com tamanho fixo inicializado com valores zero.

   ```go
   // Criar um slice com 10 elementos, todos zeros
   slice := make([]int, 10)
   fmt.Println(slice)  // [0 0 0 0 0 0 0 0 0 0]
   ```

3. **Diferença entre `nil` e Slice Vazio**
   Há uma diferença sutil entre `nil` e um slice vazio:

   ```go
   var slice1 []int        // slice nil
   slice2 := []int{}       // slice vazio (não nil)
   slice3 := make([]int, 0) // slice vazio (não nil)

   fmt.Println(slice1 == nil)  // true
   fmt.Println(slice2 == nil)  // false
   fmt.Println(slice3 == nil)  // false
   ```

**Quando usar cada método:**

| Método | Quando usar |
|--------|-------------|
| `[]int{}` | Quando você não sabe quantos elementos vai ter e não se importa com performance |
| `make([]int, 0)` | Quando você quer um slice vazio explícito (não nil) |
| `make([]int, 0, capacidade)` | Quando você sabe aproximadamente quantos elementos vai adicionar (melhor performance) |
| `make([]int, tamanho)` | Quando você precisa de um slice pré-preenchido com zeros |

**Resumo**: Use `make` quando você sabe aproximadamente quantos elementos vai adicionar (especifique capacidade) ou precisa de um slice pré-preenchido. Isso evita múltiplas realocações de memória e melhora a performance.

### 4. Converter Array em Slice
Você pode converter um array em slice usando a sintaxe de slicing. Isso cria um slice que referencia o array subjacente.

**Exemplo - Converter array completo em slice:**
```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[:]  // Converte todo o array em slice
fmt.Println(slice) // [1 2 3 4 5]
```

**Exemplo - Converter parte do array em slice (por índices):**
```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[1:3]  // Converte elementos do índice 1 até 2 (3 não incluído)
fmt.Println(slice)   // [2 3]
```

**Explicação:**
- `array[:]` - Cria um slice com todos os elementos do array
- `array[1:3]` - Cria um slice com elementos dos índices 1 e 2 (o índice 3 não é incluído)
- O slice criado referencia o array original, então modificações no slice podem afetar o array (e vice-versa)

**IMPORTANTE**: Quando você converte um array em slice, o slice compartilha a mesma memória do array. Alterações no slice podem modificar o array original.

## Operações com Slices

### Append - Adicionar Elementos
A função `append` adiciona elementos ao final do slice e retorna um novo slice. É necessário reatribuir o resultado ao slice original.

**Exemplo - Adicionar um elemento:**
```go
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 6)
fmt.Println(slice) // [1 2 3 4 5 6]
```

**Exemplo - Adicionar múltiplos elementos:**
```go
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 6, 7, 8, 9, 10)
fmt.Println(slice) // [1 2 3 4 5 6 7 8 9 10]
```

**IMPORTANTE**: O `append` retorna um novo slice. Sempre reatribua o resultado: `slice = append(slice, valor)`

### Remover Item por Índice
Para remover um item de um slice, você precisa usar a técnica de "cortar" (slicing) combinada com `append`. Não existe uma função nativa para remover itens.

**Como funciona:**
1. Pegue todos os elementos antes do índice a ser removido: `slice[:índice]`
2. Pegue todos os elementos depois do índice: `slice[índice+1:]`
3. Use `append` com o operador `...` para desempacotar o segundo slice

**Exemplo:**
```go
slice := []int{1, 2, 3, 4, 5}
// Remove o item no índice 2 (valor 3)
slice = append(slice[:2], slice[3:]...)
fmt.Println(slice) // [1 2 4 5]
```

**Explicação detalhada:**
- `slice[:2]` - Pega elementos do índice 0 até 1: `[1, 2]`
- `slice[3:]` - Pega elementos do índice 3 até o final: `[4, 5]`
- `slice[3:]...` - O operador `...` desempacota o slice para passar como argumentos individuais ao `append`
- `append(slice[:2], slice[3:]...)` - Concatena `[1, 2]` com `[4, 5]`, resultando em `[1, 2, 4, 5]`

**Para remover outros índices:**
- Remover índice 0: `slice = append(slice[1:])`
- Remover índice 1: `slice = append(slice[:1], slice[2:]...)`
- Remover último índice: `slice = slice[:len(slice)-1]`

## Slicing (Cortar Slices)
Slicing permite criar um novo slice a partir de uma porção de um slice existente usando a sintaxe `slice[início:fim]`.

**Sintaxe:**
- `slice[início:fim]` - Elementos do índice `início` até `fim-1` (fim não incluído)
- `slice[:fim]` - Do início até `fim-1`
- `slice[início:]` - Do índice `início` até o final
- `slice[:]` - Todo o slice (cria uma view do mesmo array subjacente, não uma cópia completa)

**Exemplo:**
```go
slice := []int{1, 2, 3, 4, 5}
parte1 := slice[1:3]  // [2, 3] - índices 1 e 2
parte2 := slice[:3]   // [1, 2, 3] - índices 0, 1 e 2
parte3 := slice[2:]   // [3, 4, 5] - índices 2, 3 e 4
```

## Verificar o Tipo de um Slice
Para verificar o tipo de um slice em tempo de execução, você pode usar o pacote `reflect`.

**Exemplo:**
```go
import "reflect"

slice := []int{1, 2, 3, 4, 5}
fmt.Println(reflect.TypeOf(slice)) // []int
```

**OBSERVAÇÃO**: O `reflect.TypeOf()` retorna o tipo dinâmico da variável. Para slices, mostrará `[]tipo`, onde `tipo` é o tipo dos elementos (ex: `[]int`, `[]string`, etc.).

## Operador `...` (Unpack/Spread)
O operador `...` desempacota um slice, permitindo passar seus elementos como argumentos individuais para uma função.

**Exemplo:**
```go
slice1 := []int{1, 2, 3}
slice2 := []int{4, 5, 6}
// Sem ... não funciona:
// append(slice1, slice2) // ERRO!

// Com ... funciona:
slice3 := append(slice1, slice2...) // [1, 2, 3, 4, 5, 6]
```

**Uso comum:**
- Adicionar todos os elementos de um slice a outro: `append(slice1, slice2...)`
- Passar elementos de um slice como argumentos para uma função variádica: `funcao(slice...)`

## Casos de Uso
- **Coleções dinâmicas**: Armazenar listas que podem crescer ou diminuir (usuários, produtos, pedidos)
- **Processamento de dados**: Filtrar, mapear, reduzir coleções de dados
- **APIs REST**: Retornar listas de recursos (GET /users retorna []User)
- **Buffers dinâmicos**: Acumular dados de tamanho desconhecido
- **Pilhas e filas**: Implementar estruturas de dados lineares
- **Paginação**: Dividir grandes conjuntos de dados em páginas
- **Cache**: Armazenar itens em memória para acesso rápido

**Exemplo prático:**
```go
// Lista de produtos em um carrinho
var carrinho []Produto
carrinho = append(carrinho, produto1)
carrinho = append(carrinho, produto2)

// Filtrar usuários ativos
var usuariosAtivos []Usuario
for _, u := range usuarios {
    if u.Ativo {
        usuariosAtivos = append(usuariosAtivos, u)
    }
}

// Processar logs em lote
var logs []string
logs = append(logs, "Log 1", "Log 2", "Log 3")
processarLogs(logs...)
```

