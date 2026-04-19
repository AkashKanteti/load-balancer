# DLB — Agent Guide

## Project Overview

DLB is an internal L4/L7 load balancer prototype. It distributes traffic across backend service clusters with configurable balancing strategies, health checking, and observability integration.

**Status:** Prototype (pre-approval)
**Owner:** @AkashKanteti
**Jira:** `TRST` board

## Repository Structure

```
lb/                → Load balancer entry point
be/                → Test backend server
```

Target structure (in progress):

```
cmd/dlb/           → Main binary entry point
internal/
  config/          → YAML config loading, validation, hot-reload
  server/          → L4 (TCP) and L7 (HTTP/S) listeners
  router/          → L7 request routing and matching
  balancer/        → Load balancing algorithm implementations
  pool/            → Backend pool management and state
  health/          → Health check probes (TCP, HTTP, gRPC)
  middleware/      → L7 middleware chain (rate limit, circuit breaker, retry)
  proxy/           → L4/L7 reverse proxying
  metrics/         → Prometheus metrics and collectors
configs/           → Example and default configuration files
deployments/       → Dockerfile and Kubernetes manifests
scripts/           → Local dev helpers (echo servers, test traffic)
```

## Building and Running

```bash
# Build
make build

# Run locally
make run

# Run with custom config
./bin/dlb --config configs/dlb.yaml

# Run tests
make test

# Run linter
make lint
```

## Testing

- Unit tests live alongside source files (`*_test.go`)
- Run the full suite before raising a PR:
  ```bash
  make test
  ```
- For local manual testing, spin up echo backends:
  ```bash
  go run scripts/test-backends/echo_server.go --port 8081
  go run scripts/test-backends/echo_server.go --port 8082
  ```

## Configuration

Config is YAML-based. See `configs/dlb.yaml` for the full schema. Key sections:

- `listeners` — L4/L7 listener binds
- `pools` — Backend pools with algorithm selection and backend addresses
- `health_checks` — Probe type, interval, thresholds per pool
- `middleware` — Rate limiting, circuit breaker, retry (L7 only)
- `metrics` — Prometheus export settings
- `logging` — Log level and format

## Code Conventions

- Standard Go project layout (`cmd/`, `internal/`)
- All balancer algorithms implement the `balancer.Balancer` interface
- All middleware implements `middleware.Middleware` (wraps `http.Handler`)
- Health probes implement `health.Probe`
- Use `zerolog` for structured logging — no `fmt.Print` or `log` stdlib in production code
- Errors wrapped with context: `fmt.Errorf("failed to do X: %w", err)`
- Metrics use Prometheus client_golang — register in `internal/metrics/`

## Key Design Decisions

1. **Interface-driven algorithms** — All balancers implement `Balancer`, swappable per-pool via config
2. **Separate L4/L7 paths** — L4 operates on raw TCP connections; L7 uses `net/http` with middleware
3. **Health state in pool** — Unhealthy backends excluded from selection automatically
4. **Config-driven** — No hardcoded defaults for production-facing settings
5. **Graceful lifecycle** — Context-based shutdown with connection draining
