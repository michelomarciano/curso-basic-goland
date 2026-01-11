# SWITCH

O `switch` em Go é uma estrutura de controle que permite executar diferentes blocos de código baseado no valor de uma expressão. Diferente de outras linguagens, Go não precisa de `break` - ele sai automaticamente quando encontra uma condição correspondente.

**OBSERVAÇÃO**: Em Go, o `switch` não precisa de `break` como em outras linguagens. Ele sai automaticamente quando encontra uma condição correspondente.

## Switch Básico
```go
var numero int = 2

switch numero {
case 1:
    fmt.Println("O número é 1")
case 2:
    fmt.Println("O número é 2")
case 3:
    fmt.Println("O número é 3")
default:
    fmt.Println("O número é diferente de 1, 2 e 3")
}
```

## Switch com Retorno
O `switch` pode ser usado diretamente em um `return`:

```go
func DiaDaSemana(numero int) string {
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
```

## Switch sem Expressão (Switch Verdadeiro)
Quando você omite a expressão após `switch`, ele avalia cada `case` como uma condição booleana:

```go
var numero int = 5

switch {
case numero < 5:
    fmt.Println("Menor que 5")
case numero == 5:
    fmt.Println("Igual a 5")
case numero > 5:
    fmt.Println("Maior que 5")
}
```

## Switch com Inicialização de Variável
Você pode inicializar uma variável no `switch`:

```go
switch numero := 5; numero {
case 1, 2, 3:
    fmt.Println("Entre 1 e 3")
case 4, 5, 6:
    fmt.Println("Entre 4 e 6")
default:
    fmt.Println("Outro valor")
}
```

## Múltiplos Valores no Case
Você pode ter múltiplos valores em um único `case`:

```go
switch numero {
case 1, 3, 5, 7, 9:
    fmt.Println("Ímpar")
case 2, 4, 6, 8, 10:
    fmt.Println("Par")
}
```

## Fallthrough
Por padrão, Go não executa o próximo `case` após encontrar uma correspondência. Se você quiser esse comportamento, use `fallthrough`:

```go
switch numero {
case 1:
    fmt.Println("Um")
    fallthrough  // Continua para o próximo case
case 2:
    fmt.Println("Dois")
case 3:
    fmt.Println("Três")
}
// Se numero == 1, imprime "Um" e "Dois"
```

**OBSERVAÇÃO**: O `fallthrough` deve ser a última instrução no `case` e não pode ser usado no último `case` de um `switch`.

## Casos de Uso
- **Menus e opções**: Processar escolhas do usuário em menus interativos
- **Máquinas de estado**: Implementar lógica de estados (ex: pedido: pendente, processando, concluído)
- **Parsers**: Processar tokens ou comandos com diferentes valores
- **Conversão de tipos**: Converter entre tipos diferentes baseado em valores
- **Validação de entrada**: Validar diferentes tipos de entrada
- **Roteamento**: Direcionar requisições baseado em rotas ou métodos HTTP

**Exemplo prático:**
```go
// Processar comando do usuário
switch comando {
case "create":
    criarItem()
case "read":
    lerItem()
case "update":
    atualizarItem()
case "delete":
    deletarItem()
default:
    fmt.Println("Comando inválido")
}

// Determinar tipo de arquivo
switch extensao {
case ".jpg", ".png", ".gif":
    processarImagem()
case ".mp4", ".avi":
    processarVideo()
case ".pdf":
    processarPDF()
}
```

