# 🕸  GraphQL

## Dependency

```bash
go get github.com/graphql-go/graphql
```

## cURL example
```bash
curl -X POST http://localhost:8080/graphql -d '{"query":"{ book(ID: 1) { Title Author { Name } } }"}'
```
