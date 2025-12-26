# Go Web Application Template

A production-ready Go web application template with modern best practices, including:

- Clean architecture with `cmd/`, `internal/`, and `ui/` structure
- HTTP server with graceful shutdown
- Structured logging with `slog`
- Configuration management (flags + environment variables)
- Request validation utilities
- Security headers middleware
- Docker support with multi-arch builds
- GitHub Actions CI/CD with semantic versioning
- Embedded static files and templates
- Makefile for common tasks

## Project Structure

```
.
├── cmd/
│   └── app/              # Application entry point
│       └── main.go
├── internal/
│   ├── config/           # Configuration parsing
│   ├── log/              # Logging utilities
│   ├── server/           # HTTP server setup and handlers
│   └── validator/        # Validation utilities
├── ui/
│   ├── static/           # Static assets (CSS, JS)
│   └── templates/        # HTML templates
├── scripts/              # Build and utility scripts
├── .github/
│   └── workflows/       # CI/CD workflows
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod
```

## Getting Started

### 1. Copy the Template

```bash
cp -r go-template /path/to/your-new-project
cd /path/to/your-new-project
```

### 2. Update Module Name

Replace all occurrences of `github.com/yourusername/yourproject` with your actual module path:

```bash
# Update go.mod
sed -i '' 's|github.com/yourusername/yourproject|github.com/yourusername/yourproject|g' go.mod

# Update all Go files
find . -name "*.go" -type f -exec sed -i '' 's|github.com/yourusername/yourproject|github.com/yourusername/yourproject|g' {} +
```

Or use your IDE's find-and-replace feature.

### 3. Customize Configuration

Edit `internal/config/config.go` to add your application-specific configuration:

```go
type Config struct {
    Port        string
    DatabaseURL string
    APIKey      string
    // Add your fields here
}
```

### 4. Add Your Handlers

Edit `internal/server/server.go` and `internal/server/handlers.go` to add your routes and handlers:

```go
// In registerAPIRoutes
apiGroup.HandleFunc("GET /users", getUsersHandler(logger, cfg))
apiGroup.HandleFunc("POST /users", createUserHandler(logger, cfg))
```

### 5. Initialize Go Modules

```bash
go mod tidy
```

### 6. Build and Run

```bash
# Build
make build

# Run
./dist/yourproject --port 8000 --verbose

# Or run directly
go run ./cmd/app --port 8000 --verbose
```

## Features

### Configuration

Configuration supports both command-line flags and environment variables:

```bash
# Via flag
./app --port 8080

# Via environment variable
APP_PORT=8080 ./app
```

### Logging

Logging is controlled by the `--verbose` or `-v` flag:

```bash
# Verbose logging (debug level)
./app --verbose

# Silent mode (errors only)
./app
```

### Docker

Build and run with Docker:

```bash
# Build
docker build -t yourproject:latest .

# Run
docker run -p 8000:8000 yourproject:latest

# Or use docker-compose
docker-compose up
```

### Multi-Architecture Builds

Build for multiple platforms:

```bash
make build-prod
```

This creates binaries in `dist/` for:
- Linux (amd64, arm64)
- macOS (amd64, arm64)

### CI/CD

The GitHub Actions workflow:
- Uses semantic versioning based on commit messages
- Builds and pushes Docker images on version releases
- Skips builds for `chore:` commits
- Supports multi-arch Docker builds

**Required Secrets:**
- `DOCKERHUB_USERNAME` - Your Docker Hub username
- `DOCKERHUB_TOKEN` - Your Docker Hub access token

**Commit Message Format:**
- `feat:` - Minor version bump
- `fix:` - Patch version bump
- `feat!:` or `BREAKING CHANGE:` - Major version bump
- `chore:` - No release (Docker build skipped)

## Makefile Targets

- `make build` - Build binary locally
- `make build-prod` - Build for multiple platforms
- `make clean` - Remove build artifacts
- `make format` - Format Go code
- `make test` - Run tests
- `make docker-build` - Build Docker image
- `make docker-clean` - Clean up Docker resources
- `make docker-clean-all` - Clean Docker resources and build cache

## UI Templates

The template includes a basic HTML template structure:

- `ui/templates/layouts/base.tmpl.html` - Base layout
- `ui/templates/pages/` - Page templates
- `ui/static/` - Static assets (CSS, JS)

To use templates, uncomment the embed directive in `ui/efs.go`:

```go
//go:embed "templates/*" "static/*"
var Files embed.FS
```

Then use the template cache in your handlers (see example in `internal/server/server.go`).

## Validation

The `internal/validator` package provides validation utilities:

```go
import "github.com/yourusername/yourproject/internal/validator"

type UserRequest struct {
    Email string `json:"email"`
    validator.Validator
}

func (r *UserRequest) Validate(v *validator.Validator) {
    v.CheckField(validator.NotBlank(r.Email), "email", "email is required")
    v.CheckField(validator.Matches(r.Email, emailRegex), "email", "invalid email format")
}
```

## Security

The template includes:
- Security headers middleware (CSP, X-Frame-Options, etc.)
- Request throttling
- Request size limits
- Graceful shutdown
- Health check endpoint

## Dependencies

The template uses:
- `github.com/en9inerd/go-pkgs` - Router and middleware utilities

You can replace this with your preferred router/middleware library if needed.

## Customization Checklist

- [ ] Update module name in `go.mod` and all Go files
- [ ] Add your configuration fields in `internal/config/config.go`
- [ ] Implement your handlers in `internal/server/handlers.go`
- [ ] Register your routes in `internal/server/server.go`
- [ ] Update Docker image name in `.github/workflows/release-and-docker.yml`
- [ ] Customize UI templates in `ui/templates/`
- [ ] Add your static assets to `ui/static/`
- [ ] Update `.env.example` with your environment variables
- [ ] Update `docker-compose.yml` with your service configuration
- [ ] Add your dependencies: `go get <package>`
- [ ] Write tests for your handlers
- [ ] Update this README with your project-specific information

## License

This template is provided as-is. Customize it for your needs.
