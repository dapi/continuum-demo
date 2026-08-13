# Repository instructions

This is an intentionally small brownfield service used to demonstrate Continuum adoption.

Before changing runtime behavior, inspect the existing README, `docs/`, tests, and package boundaries. Treat existing code and docs as evidence; do not invent missing product or architecture facts.

Keep domain behavior in `internal/orders` and HTTP translation in `cmd/api`. Preserve backward-compatible JSON fields unless an issue explicitly changes the contract.

Verification for runtime changes:

```sh
go test ./...
go vet ./...
```

Do not add Memory Bank files manually unless the task is explicitly to adopt Continuum / Memory Bank.
