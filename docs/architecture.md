# Architecture

This repository is intentionally a small brownfield service for the Continuum demonstration.

## Current boundaries

- `cmd/api` owns HTTP transport and status-code mapping.
- `internal/orders` owns order validation and in-memory state.
- HTTP handlers must not duplicate domain validation.
- Domain code must not depend on `net/http`.

## Current constraints

- IDs are generated in-process and are unique only for one process lifetime.
- Storage is in memory by design for the demo baseline.
- `POST /orders` currently creates a new order for every accepted request.
- Backward compatibility of the existing JSON fields (`id`, `customer`, `amount`) matters.

## Quality bar

Changes should include focused tests, preserve package boundaries, pass `go test ./...` and `go vet ./...`, and avoid introducing dependencies unless they materially simplify the implementation.
