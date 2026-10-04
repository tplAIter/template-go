<p align="center">
  <img src="https://raw.githubusercontent.com/tplAIter/.github/main/assets/banner.png?v=20260928" alt="tplAIter — Build with blocks. Spend fewer tokens." width="100%">
</p>

<h1 align="center">template-go</h1>

<p align="center">A Go service template assembled with tplAIter's existing Go renderer.</p>

<p align="center"><strong>Status: public development preview · local manifest</strong></p>

<p align="center"><a href="https://github.com/tplAIter/tplaiter">core CLI</a> · <a href="https://github.com/tplAIter/template-base">base template</a> · <a href="https://github.com/tplAIter/template-rust">Rust template</a> · <a href="https://github.com/tplAIter/tplaiter/blob/main/docs/template-validation.md">validation workflow</a></p>

`template-go` provides a small Go service skeleton rendered by Go's `text/template` engine. It contains typed standard-library configuration, an explicit composition root, and a neutral entity generator that creates controller, service, and repository packages under a parent entity type.

In the intended MCP-assisted workflow, an agent selects the parameters and
blocks it needs instead of retyping this boilerplate. The rendered result stays
plain Go and can be checked, reviewed, and adapted locally. This preview does
not claim that the live MCP or project lifecycle is complete.

## Capabilities

- An `entity` generator with safe name-derived paths and manual dependency-injection wiring.
- An optional Temporal adapter pinned to `go.temporal.io/sdk v1.29.1`. Its explicitly empty registry is idle: no dialing, polling, or automatic execution. Adding registrations enables a client and worker with finite dial and stop timeouts.
- A standard-library-only default render with bounded HTTP timeouts and graceful shutdown.
- A native template contract with no declared external dependencies.

## Verification

The repository workflow runs on pushes, pull requests, and manual dispatch. It first uses the pinned core template-check action to validate the manifest and render every fixture combination. Each rendered Go project is then checked with:

```sh
go test ./...
go build ./...
```

Those commands apply to each rendered fixture. The shared checker validates template files and does not run template hooks or manifest commands; its boundaries are documented in the [core validation contract](https://github.com/tplAIter/tplaiter/blob/main/docs/template-validation.md).

## Status

This is a local `0.0.0-local` preview. The core CLI's live generation path remains pending. Actual `gen.Generate` preflight regression coverage requires a core-owned checker interface; template checks do not claim live CLI generation coverage.
