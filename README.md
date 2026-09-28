<p align="center">
  <img src="https://raw.githubusercontent.com/tplAIter/.github/main/assets/banner.png" alt="tplAIter — Templates for the way you build." width="100%">
</p>

<h1 align="center">template-go</h1>

<p align="center">A Go service template assembled with tplAIter's existing Go renderer.</p>

<p align="center"><strong>Status: private development preview · local manifest</strong></p>

<p align="center"><a href="https://github.com/tplAIter/tplaiter">core CLI</a> · <a href="https://github.com/tplAIter/template-base">base template</a> · <a href="https://github.com/tplAIter/template-rust">Rust template</a> · <a href="https://github.com/tplAIter/tplaiter/blob/main/docs/template-validation.md">validation workflow</a></p>

`template-go` provides a Go service layout rendered by Go's `text/template` engine. It covers service, repository, transport, worker, migration, OpenAPI, and observability assets; integration and lifecycle work is still in progress.

## Capabilities

- Manifest-driven settings for database, brokers, cache, object storage, and background workflows.
- Service, repository, transport, worker, migration, OpenAPI, and observability render assets.
- A native template contract with no declared external dependencies.

## Verification

The repository workflow runs on pushes, pull requests, and manual dispatch. It first uses the pinned core template-check action to validate the manifest and render every fixture combination. Each rendered Go project is then checked with:

```sh
go test ./...
go build ./...
```

Those commands apply to each rendered fixture. The shared checker validates template files and does not run template hooks or manifest commands; its boundaries are documented in the [core validation contract](https://github.com/tplAIter/tplaiter/blob/main/docs/template-validation.md).

## Status

This is a local `0.0.0-local` preview. Further template validation and lifecycle integration are still in progress; this repository does not claim complete production readiness or a direct `new`/`update` workflow by itself.
