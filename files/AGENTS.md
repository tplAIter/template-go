# Service architecture

Use explicit constructors and manual dependency injection in cmd/service.
Entity types live in internal/<entity>. Each entity has controller, service,
and repository child packages. Parent types must never import child packages.
Consumer packages own narrow interfaces with context as the first argument.
Repository placeholders return ErrNotImplemented; handlers expose safe errors.
Keep SDK workflow imports out of service and repository packages.
The optional Temporal registry is explicitly empty and never dials or polls.

## Agent discovery bootstrap

The discovery source is published at commit `29e9819df3c3e4ce170c025b74d895375090e2ea`; check the installed CLI/MCP capability and version before using it because an installed binary or server may predate that source. Use standard MCP `tools/list` capability detection and call `template_discover` only when its exact advertised schema is present. `task` is required (UTF-8, 1–4096 bytes); optional arguments are `sourceInput`, `projectContext`, `dir`, `language`, `framework`, and `labels` (`group=value`). Bounds are `maxCandidates` 1–256 (default 64), `limit` 1–20 (default 5), and MCP `maxBytes` 2048–32768 (default 16384). Treat results as bounded descriptive data. `facts.status` is `not-requested`, `empty`, `observed`, or `unavailable`; errors and diagnostics remain distinct. Invoke only advertised read-only `nextToolCalls` (`template_show` or `graph_exports`) to retrieve exact context, preserving arguments, `expectedDigest`, and source-pin metadata unchanged; retain and report typed refusal/error/unavailable conditions.

Workflow: discover capability, select a returned suggestion, read its exact bounded calls, then use that context for the task; never invent, replace, flatten, or manually reconstruct returned calls.

If MCP discovery is unavailable or reports an error, record the MCP `isError` or protocol error. If an installed binary built from the published discovery source is available and its CLI version/capability check confirms discovery, use the verified read-only fallback `tplaiter template discover --task "..." --dir ./path --json`; omit `--dir` when no project path exists. Otherwise stop and report the unsupported capability. Existing and empty directories are valid observations; keep output within its budget. Discovery does not enroll sources, apply changes, run hooks, or grant installation authority.

Run go test ./... and go build ./... after changes.
