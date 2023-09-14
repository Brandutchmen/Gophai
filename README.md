# Gophai
> A Go backend GraphQL api boilerplate with Authentication, Postgres, Hot Reload, and more.

- 🐳 [Docker](https://www.docker.com/) Development Environment
- ⚡️ Live Reloading with [air](https://github.com/cosmtrek/air)
- 🚀 Routing with [Chi](https://go-chi.io/)
- 💾 [Postgres](https://www.postgresql.org/) with [Gorm](https://gorm.io/)
- 🧩 Fully Typed GraphQL API with [gqlgen](https://gqlgen.com/)
- 🔑 Authentication (Coming Soon)

# Roadmap
- [ ] Remove Gorm in favor of sqlc for db
- [ ] Add authentication
- [ ] Add authorization middleware
- [ ] Create role/permission scaffold

# Quick Start

Clone this project

```bash
git clone https://github.com/Brandutchmen/Gophai.git
```

or generate a new project with this [template](https://github.com/new?template_name=Gophai&template_owner=Brandutchmen)

---
Copy the example environment file
```bash
cp .env.example .env
```

---
Start the development environment
```bash
docker-compose up
```

---

Open the GraphQL Playground at [http://localhost:3000/api/graphql/playground](http://localhost:3000/api/graphql/playground)

---

# Running Tests

```bash
make run-tests
```
