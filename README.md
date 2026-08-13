# Continuum Demo

A deliberately small **brownfield** Go service for demonstrating [Continuum](https://github.com/dapi/continuum): persistent project context, issue-driven agent execution, isolated work, review/fix convergence, and CI-backed delivery.

The baseline intentionally starts **without `memory-bank/`**. It already has code, tests, product notes, architecture constraints, repository instructions, and CI — enough history for an agent to discover and preserve.

## Run

```sh
go test ./...
go run ./cmd/api
```

The API listens on `:8080`.

```sh
curl -s localhost:8080/healthz
curl -s -X POST localhost:8080/orders \
  -H 'content-type: application/json' \
  -d '{"customer":"acme","amount":12500}'
curl -s localhost:8080/orders
```

## What makes this brownfield

The project already has behavioral contracts and design constraints spread across multiple sources:

- `AGENTS.md` — repository-level working rules;
- `docs/product.md` — users, current API contract, and a known operational problem;
- `docs/architecture.md` — package boundaries, constraints, and quality bar;
- `internal/orders/*` — executable domain behavior and tests;
- `.github/workflows/ci.yml` — delivery verification.

Continuum adoption must discover these sources before treating generated context as authoritative.

## The two-prompt demo

### Prompt 1 — attach Continuum

Give a capable coding agent this prompt from the repository root:

> Attach Continuum to this existing brownfield repository. Follow the official Memory Bank brownfield adaptation protocol. First discover and record evidence from existing repository sources without consulting generic Memory Bank content as project truth. Then install/adapt Memory Bank, preserve provenance and unknowns, run its validation, and stop after the repository is ready for issue-driven work. Do not implement product changes yet.

Expected outcome: the repository gains an evidence-backed, validated `memory-bank/` that reflects the existing project rather than replacing it with generic template assumptions.

### Prompt 2 — execute the issue

Then run the real issue through the normal Continuum flow:

> Execute GitHub issue #1 through Continuum. Use the adapted Memory Bank as durable context, preserve existing contracts, implement and test the change, then converge review and CI until the change is ready as a verified PR. Stop on a genuine human-gate decision instead of guessing.

The operator-facing interface is two prompts. Internally, Continuum performs the discovery, context adaptation, isolated execution, verification, review/fix, and CI work required to make those prompts reliable.

## Reproducing with Continuum tools

Typical local flow after Prompt 1:

```sh
start-issue 1 --agent codex
# agent implements the issue in its worktree
code-converge
```

See the parent project: [dapi/continuum](https://github.com/dapi/continuum).
