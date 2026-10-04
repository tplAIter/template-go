# Service architecture

Use explicit constructors and manual dependency injection in cmd/service.
Entity types live in internal/<entity>. Each entity has controller, service,
and repository child packages. Parent types must never import child packages.
Consumer packages own narrow interfaces with context as the first argument.
Repository placeholders return ErrNotImplemented; handlers expose safe errors.
Keep SDK workflow imports out of service and repository packages.
The optional Temporal registry is explicitly empty and never dials or polls.

Run go test ./... and go build ./... after changes.
