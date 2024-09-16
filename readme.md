# API Go

---

Esta é uma API simples desenvolvida em Go, que retorna uma mensagem JSON de boas-vindas.

## Endpoint

- `GET /`

Resposta:

```json
{
  "message": "hello go"
}
```

---

## Site

A API está disponível na Vercel: [API Go](https://api-go-one.vercel.app/)

---

## Como rodar localmente

1. Clone o repositório:

   ```bash
   git clone https://github.com/danielvor/api-go.git
   ```

2. Acesse o diretório do projeto:

   ```bash
   cd api-go
   ```

3. Instale as dependências:

   ```bash
   go mod tidy
   ```

4. Execute o servidor:

   ```bash
   go run main.go
   ```

5. Acesse a API em:

   ```
   http://localhost:3000/
   ```



### Dependências

- Go
- [rs/cors](https://github.com/rs/cors)

