# ESTRUTURAS DE CONTROLE

As estruturas de controle permitem que o programa tome decisões e execute diferentes blocos de código baseado em condições. Em Go, as principais estruturas de controle são `if/else` e `switch`.

## IF/ELSE

O `if/else` é a estrutura de controle mais básica e permite executar código condicionalmente baseado em uma expressão booleana.

### If Básico

```go
var idade int = 18

if idade >= 18 {
    fmt.Println("Você é maior de idade")
} else {
    fmt.Println("Você é menor de idade")
}
```

### If sem Else

Quando você só precisa executar código se uma condição for verdadeira:

```go
var saldo float64 = 100.0

if saldo > 0 {
    fmt.Println("Você tem saldo disponível")
}
```

### If com Múltiplas Condições

Você pode usar operadores lógicos para combinar condições:

```go
var idade int = 20
var temCarteira bool = true

if idade >= 18 && temCarteira {
    fmt.Println("Pode dirigir")
} else {
    fmt.Println("Não pode dirigir")
}
```

### If/Else If/Else

Para múltiplas condições, use `else if`:

```go
var nota float64 = 8.5

if nota >= 9.0 {
    fmt.Println("Excelente")
} else if nota >= 7.0 {
    fmt.Println("Bom")
} else if nota >= 5.0 {
    fmt.Println("Regular")
} else {
    fmt.Println("Reprovado")
}
```

### If com Inicialização de Variável

Uma característica poderosa do Go é poder inicializar uma variável diretamente no `if`. A variável só existe no escopo do `if`:

```go
var numero int = 12

if idade := numero; idade >= 18 {
    fmt.Println("Você é maior de idade")
} else {
    fmt.Println("Você é menor de idade")
}
// idade não existe aqui fora do if
```

**OBSERVAÇÃO**: A variável declarada no `if` só existe dentro do escopo do bloco `if/else`. Isso é útil para evitar poluir o escopo externo com variáveis temporárias.

### If com Múltiplas Inicializações

Você pode inicializar múltiplas variáveis separadas por vírgula:

```go
if valor, erro := calcular(); erro == nil {
    fmt.Println("Resultado:", valor)
} else {
    fmt.Println("Erro:", erro)
}
```

### Casos de Uso - If/Else

- **Validações**: Verificar se dados estão corretos antes de processar
- **Controle de acesso**: Verificar permissões, idade, autenticação
- **Tratamento de erros**: Verificar se operações foram bem-sucedidas
- **Lógica de negócio**: Implementar regras condicionais
- **Filtros**: Processar apenas dados que atendem certas condições
- **Valores padrão**: Definir valores quando condições não são atendidas

**Exemplo prático:**
```go
// Validar entrada do usuário
if email == "" {
    fmt.Println("Email é obrigatório")
} else if !strings.Contains(email, "@") {
    fmt.Println("Email inválido")
} else {
    processarEmail(email)
}

// Verificar permissões
if !usuario.Autenticado {
    redirecionarParaLogin()
} else if !usuario.TemPermissao("admin") {
    mostrarErro("Acesso negado")
} else {
    mostrarPainelAdmin()
}

// Tratamento de erro
if err := salvarDados(dados); err != nil {
    log.Printf("Erro ao salvar: %v", err)
    return err
} else {
    fmt.Println("Dados salvos com sucesso")
}
```

## SWITCH

O `switch` em Go é uma estrutura de controle que permite executar diferentes blocos de código baseado no valor de uma expressão. Diferente de outras linguagens, Go não precisa de `break` - ele sai automaticamente quando encontra uma condição correspondente.

**OBSERVAÇÃO**: Em Go, o `switch` não precisa de `break` como em outras linguagens. Ele sai automaticamente quando encontra uma condição correspondente.

### Switch Básico

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

### Switch com Retorno

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

### Switch sem Expressão (Switch Verdadeiro)

Quando você omite a expressão após `switch`, ele avalia cada `case` como uma condição booleana. Isso é útil para condições complexas:

```go
var numero int = 12

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
```

**OBSERVAÇÃO**: O switch sem expressão é equivalente a uma série de `if/else if`, mas pode ser mais legível em alguns casos.

### Switch com Inicialização de Variável

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

### Múltiplos Valores no Case

Você pode ter múltiplos valores em um único `case`:

```go
var numero int = 5

switch numero {
case 1, 3, 5, 7, 9:
    fmt.Println("Ímpar")
case 2, 4, 6, 8, 10:
    fmt.Println("Par")
default:
    fmt.Println("Número fora do range")
}
```

### Fallthrough

Por padrão, Go não executa o próximo `case` após encontrar uma correspondência. Se você quiser esse comportamento, use `fallthrough`:

```go
var numero int = 1

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

### Casos de Uso - Switch

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
default:
    fmt.Println("Tipo de arquivo não suportado")
}

// Máquina de estados
switch estadoPedido {
case "pendente":
    processarPagamento()
case "pago":
    enviarProduto()
case "enviado":
    aguardarConfirmacao()
case "concluido":
    finalizarPedido()
default:
    tratarErro()
}
```

## QUANDO USAR IF/ELSE VS SWITCH

### Use If/Else quando:
- **Condições complexas**: Quando você precisa avaliar expressões booleanas complexas com múltiplos operadores lógicos
- **Ranges de valores**: Quando você precisa verificar se um valor está dentro de um range (ex: `if idade >= 18 && idade < 65`)
- **Condições não discretas**: Quando os valores não são discretos ou não fazem sentido agrupar em cases
- **Poucas condições**: Quando há apenas 2-3 condições simples

**Exemplo:**
```go
// Melhor com if/else - condições complexas
if idade >= 18 && idade < 65 && temCarteira && !temMultas {
    permitirDirigir()
}

// Melhor com if/else - ranges
if nota >= 9.0 {
    // Excelente
} else if nota >= 7.0 {
    // Bom
}
```

### Use Switch quando:
- **Múltiplos valores discretos**: Quando você precisa verificar um valor contra múltiplos valores específicos
- **Código mais limpo**: Quando o switch torna o código mais legível que múltiplos if/else
- **Múltiplos cases**: Quando há 4 ou mais condições para o mesmo valor
- **Valores constantes**: Quando você está comparando contra valores constantes conhecidos

**Exemplo:**
```go
// Melhor com switch - múltiplos valores discretos
switch dia {
case 1:
    return "Domingo"
case 2:
    return "Segunda"
// ... mais cases
}

// Melhor com switch - múltiplos valores no mesmo case
switch mes {
case 1, 3, 5, 7, 8, 10, 12:
    return 31  // Meses com 31 dias
case 4, 6, 9, 11:
    return 30  // Meses com 30 dias
case 2:
    return 28  // Fevereiro
}
```

## BOAS PRÁTICAS

1. **Use nomes descritivos**: Variáveis e condições devem ter nomes claros
2. **Evite aninhamento excessivo**: Se você tem muitos `if` aninhados, considere refatorar
3. **Use switch para múltiplos valores**: Quando apropriado, switch é mais legível
4. **Sempre trate o caso default**: Em switches, considere sempre ter um `default`
5. **Mantenha condições simples**: Evite condições muito complexas - extraia para funções se necessário
6. **Use inicialização no if/switch**: Aproveite a inicialização de variáveis para manter o escopo limpo

**Exemplo de código limpo:**
```go
// Bom - condição clara e simples
if usuario.TemPermissao("admin") {
    executarAcaoAdmin()
}

// Evite - condição muito complexa
if usuario != nil && usuario.Perfil != nil && usuario.Perfil.Tipo == "admin" && usuario.Perfil.Ativo && !usuario.Perfil.Suspenso {
    executarAcaoAdmin()
}

// Melhor - extrair para função
if isAdminAtivo(usuario) {
    executarAcaoAdmin()
}
```
