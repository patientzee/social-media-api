# Social Media API
REST API using Gin and PostgreSQL

## Project Structure
```
social-media-api/
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── main.go
├── database/
│   └── schema.sql
├── handlers/
│   ├── users.go
│   ├── posts.go
│   └── comments.go
├── models/
│   ├── user.go
│   ├── post.go
│   └── comment.go
└── config/
    └── database.go
```

## Setup
```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin your-repository-url
git push -u origin main
```

## Local Development
1. Create database: `createdb socialmedia`
2. Run migrations: `psql socialmedia < database/schema.sql`
3. Start server: `go run main.go`
