# HERANÇA

No Go não temos herança tradicional como em linguagens orientadas a objetos. Em vez disso, Go usa **composição** através de **structs embutidos** (embedded structs).

Quando você omite o nome do campo ao embutir uma struct em outra, você pode acessar os campos diretamente, como se fossem parte da struct externa.

**Exemplo:**
```go
type Animal struct {
    Nome  string
    Cor   string
}

type Cachorro struct {
    Animal  // Struct embutido (sem nome de campo)
    Raça string
}

// Uso:
cachorro := heranca.Cachorro{}
cachorro.Cor = "Marrom"  // Acesso direto ao campo de Animal
cachorro.Raça = "Labrador"
```

## Casos de Uso
- **Modelagem de relacionamentos**: Representar relações "é um tipo de" (ex: Cachorro é um Animal)
- **Reutilização de código**: Compartilhar campos e métodos comuns entre tipos relacionados
- **Extensibilidade**: Adicionar funcionalidades específicas a tipos base sem modificar o tipo original
- **Polimorfismo**: Usar interfaces para trabalhar com diferentes tipos que compartilham comportamento
- **Composição de funcionalidades**: Combinar múltiplas structs para criar tipos mais complexos

**Exemplo prático:**
```go
// Sistema de veículos
type Veiculo struct {
    Marca string
    Modelo string
}

type Carro struct {
    Veiculo  // Herda campos de Veiculo
    Portas int
}

type Moto struct {
    Veiculo  // Herda campos de Veiculo
    Cilindradas int
}
```

