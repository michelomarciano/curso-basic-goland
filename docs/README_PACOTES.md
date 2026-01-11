# PACOTES

**INTERNO**: Um módulo Go é um conjunto de pacotes que compõem o projeto, tanto os pacotes criados internamente quanto os pacotes usados externamente (GitHub, padrão do Go). A ideia é que ele compile em um único lugar e centralize as dependências do projeto, facilitando o gerenciamento de versões e dependências.

**EXTERNO**: Para importar pacotes externos, basta executar o seguinte comando `go get <pacote>`. Ele será adicionado ao arquivo `go.mod` e poderá ser usado em qualquer pacote do projeto, como no `main`.

## Casos de Uso
- **Organização de código**: Separar funcionalidades em pacotes lógicos (ex: `models`, `handlers`, `utils`)
- **Reutilização**: Criar bibliotecas internas que podem ser usadas em múltiplos projetos
- **Dependências externas**: Usar bibliotecas de terceiros (ex: frameworks web, ORMs, validadores)
- **Encapsulamento**: Agrupar funcionalidades relacionadas e controlar o que é exportado
- **Testes**: Organizar testes em pacotes separados

**Exemplo prático:**
```go
// Estrutura de projeto
projeto/
  ├── models/      // Modelos de dados
  ├── handlers/    // Handlers HTTP
  ├── utils/       // Funções utilitárias
  └── main.go      // Ponto de entrada

// Usar pacote externo
import "github.com/gin-gonic/gin"

// Usar pacote interno
import "projeto/models"
```

