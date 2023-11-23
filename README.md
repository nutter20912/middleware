# Middleware for micro service

## Packages
#### GraphQL
- [99designs/gqlgen](github.com/99designs/gqlgen)
- [graph-gophers/dataloader](github.com/graph-gophers/dataloader)
#### service
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [go-micro/go-micro](https://github.com/go-micro/go-micro)
- [go-micro/plugins](https://github.com/go-micro/plugins)
- [gorilla/websocket](github.com/gorilla/websocket)
#### debug tool
- [go-delve/delve](https://github.com/go-delve/delve)

---
## Setup
### Config

- generate local env file
    ```=bash
    cp .env.example .env
    ```

- configuration
    ```
    ./config
    ├── app.yaml
    ```

- override yaml config
    ```
    #app.port
    APP_PORT=1234
    ```
---
### Develop
1. Define schema
    ```
    ./graph/schema
    ├── resource.graphqls
    └── schema.graphqls
    ```
2. Configuration `./gqlgen.yml`
3. Run `make`
4. Implement the resolvers `./graph/resolver/`
> more to see doc [here](https://gqlgen.com/)

---
## Scripts
read Makefile
