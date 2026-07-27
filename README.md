# GoFramework
A Go framework inspired by Laravel, built on GoFiber.

## Directory Structure

```
├── app/
│   ├── console/commands/  # Flash commands
│   ├── http/
│   │   ├── controllers/   # HTTP Controllers
│   │   ├── middleware/     # HTTP Middleware
│   │   └── requests/       # Form Requests
│   ├── models/             # Eloquent Models (GORM)
│   ├── providers/          # Service Providers
│   └── service/            # Business Logic
├── bootstrap/              # Application Bootstrapping
├── cmd/                    # CLI Entry Points
├── config/                 # Configuration
├── database/
│   ├── migrations/         # Database Migrations
│   └── seeds/              # Database Seeders
├── framework/              # Core Framework
├── public/                 # Public Assets
├── resources/views/        # Views (HTML Templates)
├── routes/                 # Route Definitions
└── storage/                # Storage
```

## How to Use

### Via Go Run (no installation required)

```bash
go run ./cmd/flash serve
go run ./cmd/flash migrate
go run ./cmd/flash make controller User
go run ./cmd/flash route:list
```

### Via Go install (set as a global command)

```bash
go install ./cmd/flash
flash serve
flash make model Product
flash migrate
flash route:list
```

### Via Makefile

```bash
make flash serve
make flash migrate
make flash route:list
```

### HTTP Server

```bash
# Via main.go
go run main.go

# Via flash
go run ./cmd/flash serve
```

## Available Commands

| Command | Description |
|---|---|
| `serve` | Start development server |
| `migrate` | Run database migrations |
| `db:seed` | Seed database |
| `key:generate` | Generate app key |
| `route:list` | List registered routes |
| `make controller <name>` | Generate controller |
| `make model <name>` | Generate model |
| `make migration <name>` | Generate migration |
| `make seeder <name>` | Generate seeder |
