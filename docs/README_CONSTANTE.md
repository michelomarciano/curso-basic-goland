# CONSTANTE

Para declarar uma constante é necessário usar a palavra reservada `const`. 

**Exemplo:** `const PI float64 = 3.1415`

## Casos de Uso
- **Valores matemáticos**: Constantes como PI, e, etc.
- **Configurações da aplicação**: URLs base, portas padrão, timeouts
- **Códigos de status**: Códigos HTTP, códigos de erro
- **Valores mágicos**: Evitar "números mágicos" no código, substituindo por constantes nomeadas
- **Valores de configuração**: Tamanhos máximos, limites, thresholds
- **Strings de formatação**: Templates, formatos de data/hora

**Exemplo prático:**
```go
const (
    PI          = 3.14159
    MaxUsers    = 1000
    DefaultPort = 8080
    APIBaseURL  = "https://api.example.com"
    Timeout     = 30 * time.Second
)
```

