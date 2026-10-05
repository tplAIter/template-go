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

## Native offline build scope

The approved native offline CLI/MCP `run build` and default `gen`/`gen batch`
build metadata supports **`workflow=true` only**, with the exact pinned module
manifest, including `go.temporal.io/sdk v1.29.1`. Default generation commits its
files only after the approved offline compiler succeeds. This bounded acceptance
covers compilation; it does not establish Temporal client/worker execution or
acceptance of every template variant or the full template lifecycle.

**`workflow=false` is the template default and is unsupported by this v2 build
metadata.** Native `run build` and default generation refuse with
`TRUST_GO_MODULE_CLOSURE_UNAVAILABLE` before compiler execution; default generation
refuses before applying files. There is no automatic dependency-free v1 fallback,
ambient cache/toolchain fallback or online dependency resolution. Explicit
`--no-build` requests file-only generation.

The supported offline runner is **Darwin/arm64 with the authenticated, pinned
Go 1.27.1 toolchain**. Other operating systems and architectures are unsupported
and return a typed refusal; they are not covered by this acceptance.

These restrictions concern tplAIter's approved native offline build metadata.
The rendered Go code and manual `go test ./...` / `go build ./...` workflows remain
unchanged. Published-source installed CLI/MCP replay against the exact public
metadata commit is a separate completion requirement; the historical signed local
metadata-overlay proof is not a published-metadata certificate.

## Verification

The repository workflow runs on pushes, pull requests, and manual dispatch. It first uses the pinned core template-check action to validate the manifest and render every fixture combination. Each rendered Go project is then checked with:

```sh
go test ./...
go build ./...
```

Those commands apply to each rendered fixture. The shared checker validates template files and does not run template hooks or manifest commands; its boundaries are documented in the [core validation contract](https://github.com/tplAIter/tplaiter/blob/main/docs/template-validation.md).

## Status

This is a local `0.0.0-local` preview. Native offline build/default generation acceptance is limited to the scope above; it is not whole-template or all-variant acceptance. Template checks do not claim live CLI generation coverage.
