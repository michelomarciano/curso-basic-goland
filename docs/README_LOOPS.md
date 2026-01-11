# LOOPS

Loops (laços de repetição) em Go permitem executar um bloco de código repetidamente. Go tem apenas uma palavra-chave para loops: `for`. No entanto, ela pode ser usada de diferentes formas para criar diferentes tipos de loops.

**OBSERVAÇÕES**: 
- Go não tem `while`, `do-while` ou `foreach` como outras linguagens. Tudo é feito com variações do `for`.
- O `range` funciona apenas com tipos iteráveis: **slices**, **arrays**, **maps**, **strings** e **canais**. Não é possível usar `range` diretamente em **structs** (para iterar sobre campos de struct, é necessário usar reflection do pacote `reflect`).

## 1. Loop For Tradicional
O loop `for` tradicional tem três componentes: inicialização, condição e pós-instrução.

**Sintaxe:**
```go
for inicialização; condição; pós-instrução {
    // código
}
```

**Exemplo:**
```go
for i := 0; i < 10; i++ {
    fmt.Println("Loop For: ", i)
}
// Imprime números de 0 a 9
```

**Explicação:**
- `i := 0` - Inicializa a variável `i` com valor 0 (executa uma vez no início)
- `i < 10` - Condição que é verificada antes de cada iteração
- `i++` - Incrementa `i` após cada iteração

## 2. Loop While (For sem Inicialização)
Em Go, não existe `while`. Você usa `for` com apenas a condição para simular um loop `while`.

**Sintaxe:**
```go
for condição {
    // código
}
```

**Exemplo:**
```go
i := 0
for i < 10 {
    fmt.Println("Loop While: ", i)
    i++
}
// Imprime números de 0 a 9
```

**OBSERVAÇÃO**: É importante garantir que a condição eventualmente se torne `false`, caso contrário você terá um loop infinito.

## 3. Loop Do-While (For Infinito com Break)
Go não tem `do-while` nativo. Você simula usando um `for` infinito com `break` condicional.

**Sintaxe:**
```go
for {
    // código
    if condição {
        break
    }
}
```

**Exemplo:**
```go
i := 0
for {
    fmt.Println("Loop Do While: ", i)
    i++
    if i >= 10 {
        break  // Sai do loop quando i >= 10
    }
}
// Imprime números de 0 a 9
```

**Vantagem**: Garante que o código execute pelo menos uma vez antes de verificar a condição.

## 4. Loop For Range
O `for range` é usado para iterar sobre coleções (slices, arrays, maps, strings e canais). Ele retorna dois valores: o índice/chave e o valor.

**Sintaxe:**
```go
for índice, valor := range coleção {
    // código
}
```

### Com Slices e Arrays
```go
slice := []string{"Golang", "Python", "Java", "JavaScript", "C#"}
for indice, valor := range slice {
    fmt.Println(indice, valor)
}
// Saída:
// 0 Golang
// 1 Python
// 2 Java
// 3 JavaScript
// 4 C#
```

### Ignorando o Índice
Se você não precisa do índice, use `_` para ignorá-lo:
```go
slice := []string{"Golang", "Python", "Java"}
for _, valor := range slice {
    fmt.Println(valor)  // Só imprime os valores
}
```

### Ignorando o Valor
Se você só precisa do índice:
```go
for indice := range slice {
    fmt.Println(indice)  // Só imprime os índices
}
```

### Com Maps
```go
mapa := map[string]int{
    "Golang": 2009,
    "Python": 1991,
    "Java":   1995,
}

for chave, valor := range mapa {
    fmt.Println(chave, valor)
}
```

### Com Strings
```go
texto := "Go"
for indice, caractere := range texto {
    fmt.Println(indice, string(caractere))
}
// Saída:
// 0 G
// 1 o
```

**OBSERVAÇÃO**: Com strings, `range` itera sobre runes (caracteres Unicode), não bytes.

## 5. Loop Infinito
Um loop infinito executa indefinidamente até ser interrompido com `break` ou `return`.

**Sintaxe:**
```go
for {
    // código
    // Precisa de break ou return para sair
}
```

**Exemplo:**
```go
for {
    fmt.Println("Loop infinito")
    time.Sleep(time.Second)
    // Sem break, este loop nunca termina
}
```

## 6. Break e Continue

### Break
O `break` interrompe o loop imediatamente e continua a execução após o loop.

**Exemplo:**
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Sai do loop quando i == 5
    }
    fmt.Println(i)
}
// Imprime: 0, 1, 2, 3, 4
```

### Continue
O `continue` pula o restante da iteração atual e vai para a próxima iteração.

**Exemplo:**
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Pula números pares
    }
    fmt.Println(i)
}
// Imprime: 1, 3, 5, 7, 9 (apenas ímpares)
```

## 7. Labels e Break/Continue Aninhados
Você pode usar labels para controlar loops aninhados.

**Exemplo:**
```go
externo:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 2 {
            break externo  // Sai do loop externo
        }
        fmt.Println(i, j)
    }
}
```

## Casos de Uso
- **Iteração sobre coleções**: Processar elementos de slices, arrays, maps
- **Processamento de dados**: Filtrar, transformar, agregar dados
- **Contadores**: Executar código um número específico de vezes
- **Leitura de dados**: Ler arquivos linha por linha, processar streams
- **APIs e servidores**: Loops infinitos para servidores que ficam escutando requisições
- **Jogos**: Game loops que executam continuamente
- **Polling**: Verificar condições periodicamente até que sejam atendidas
- **Processamento em lote**: Processar múltiplos itens de uma vez

**Exemplo prático:**
```go
// Processar lista de usuários
usuarios := []Usuario{...}
for _, usuario := range usuarios {
    if !usuario.Ativo {
        continue  // Pula usuários inativos
    }
    processarUsuario(usuario)
}

// Servidor HTTP (loop infinito)
for {
    requisicao := aguardarRequisicao()
    processarRequisicao(requisicao)
}

// Buscar até encontrar
for {
    resultado, err := buscar()
    if err == nil {
        break  // Sai quando encontra
    }
    time.Sleep(time.Second)
}

// Processar map de configurações
config := map[string]string{
    "host": "localhost",
    "port": "8080",
}
for chave, valor := range config {
    fmt.Printf("%s = %s\n", chave, valor)
}
```

