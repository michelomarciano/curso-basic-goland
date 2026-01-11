# ARRAY

Arrays em Go são estruturas de dados de tamanho fixo. Para criar um array em Go, você precisa especificar o tamanho. Caso não especifique o tamanho, ele se tornará um slice (array dinâmico).

**OBSERVAÇÃO**: Arrays têm tamanho fixo e não podem ser redimensionados após a criação.

## Formas de Declarar e Inicializar Arrays

### 1. Declaração e Atribuição Separada
Primeiro declara o array e depois atribui valores aos índices individualmente.

**Exemplo:**
```go
var array [5]string
array[0] = "Mike"
array[1] = "Sophia"
array[2] = "Luisa"
array[3] = "Eduardo"
array[4] = "Familia"
fmt.Println(array) // [Mike Sophia Luisa Eduardo Familia]
```

### 2. Declaração e Inicialização na Mesma Linha
Declara o array e inicializa com valores diretamente usando a sintaxe de literal composto.

**Exemplo:**
```go
array2 := [5]string{"Mike", "Sophia", "Luisa", "Eduardo", "Familia"}
fmt.Println(array2) // [Mike Sophia Luisa Eduardo Familia]
```

**Vantagem**: Mais conciso que a forma anterior, permite declarar e inicializar em uma única linha.

### 3. Inicialização com Tamanho Inferido
Usa `[...]` para deixar o compilador inferir o tamanho do array automaticamente com base nos elementos fornecidos.

**Exemplo:**
```go
array3 := [...]string{"Mike", "Sophia", "Luisa", "Eduardo", "Familia"}
fmt.Println(array3) // [Mike Sophia Luisa Eduardo Familia]
```

**Vantagem**: Não precisa especificar o tamanho manualmente; o compilador calcula automaticamente (neste caso, `[5]string`).

### 4. Inicialização com Índices Específicos
Permite inicializar apenas posições específicas do array usando a sintaxe `índice: valor`. Os índices não especificados recebem o valor zero do tipo (para strings, é uma string vazia `""`).

**Exemplo:**
```go
array4 := [5]string{1: "Sophia", 3: "Eduardo"}
fmt.Println(array4) // [ Sophia  Eduardo ]
```

**Explicação:**
- `1: "Sophia"` - O índice 1 recebe o valor "Sophia"
- `3: "Eduardo"` - O índice 3 recebe o valor "Eduardo"
- Os índices 0, 2 e 4 ficam com strings vazias (valor zero de `string`)

**Vantagem**: Permite inicializar apenas posições específicas do array, deixando as demais com valores zero.

## Casos de Uso
- **Armazenar dados com tamanho fixo conhecido**: Quando você sabe exatamente quantos elementos precisa (ex: dias da semana, meses do ano)
- **Buffer de tamanho fixo**: Para operações de I/O onde você precisa de um buffer com tamanho específico
- **Tabelas de lookup**: Para criar tabelas de conversão ou mapeamento com tamanho fixo
- **Dados de configuração**: Armazenar configurações com número fixo de parâmetros
- **Processamento de dados estruturados**: Quando trabalha com formatos de dados que têm tamanho fixo (ex: headers de protocolos)

**Exemplo prático:**
```go
// Armazenar os 7 dias da semana
var diasDaSemana [7]string = [7]string{
    "Domingo", "Segunda", "Terça", "Quarta", 
    "Quinta", "Sexta", "Sábado",
}

// Buffer para leitura de arquivo (1024 bytes)
var buffer [1024]byte
```

