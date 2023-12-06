# GraphQL Examples

GraphQL のサーバーサイド実装の検証

- 一対多の関係にある部署(Department)と従業員(Employee)のケースを想定
- それぞれを全件取得するクエリに対応
- 特に DataLoader を使用して、N+1問題を回避する方法を確認

```graphql
"""
部署
"""
type Department {
  id: ID!
  name: String!
  employees: [Employee!]!
}

"""
従業員
"""
type Employee {
  id: ID!
  name: String!
  department: Department!
}

type Query {
  departments: [Department!]!
  employees: [Employee!]!
}
```

## Go

- [Golang](https://go.dev/) (v1.21.4)
- [Gin](https://github.com/gin-gonic/gin) (v1.9.1)
- [GORM](https://github.com/go-gorm/gorm) (v1.25.5)
- [gqlgen](https://github.com/99designs/gqlgen) (v0.17.40)
- [DataLoader](https://github.com/graph-gophers/dataloader) (v7.1.0)
