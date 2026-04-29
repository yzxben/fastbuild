# fastbuild

Go + [templ](https://templ.guide) + [htmx](https://htmx.org), single binary, deployed to GCE.

- GCP project ID: `fast-bld`
- Go module: `github.com/yzxben/fastbuild` (single module at repo root; subfolders are packages)

## Layout
- `web/` — web app (HTTP server, templates, static)
  - `web/main.go` — server, routes
  - `web/templates/*.templ` — components; generated `*_templ.go` is committed
  - `web/static/` — vendored assets (htmx); embedded into the binary via `//go:embed`
- additional top-level folders TBD (other services, infra, etc.)

## Design
- Philosophy: simple design with lots of animations.
- Fonts: Reddit Sans for headlines; system font for everything else.

## Dev
- `templ generate` after editing `.templ` files (install: `go install github.com/a-h/templ/cmd/templ@latest`)
- `go run ./web` from repo root — serves on `:8080` (override with `PORT`)
- `go build -o fastbuild ./web` — produces a self-contained binary
