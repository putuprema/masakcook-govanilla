# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

MasakCook is a crowd-sourced recipe web application built with:
- **Backend**: Go 1.25 with Chi router and templ templating engine
- **Frontend**: TypeScript 5 and Tailwind CSS v4, bundled with Bun
- **Dev Tools**: Air (Go live reload), Biome (frontend linting/formatting)

## Development Commands

### Start Development Server
```bash
make dev
```
This runs Air which:
1. Builds frontend (TypeScript + CSS) via `bun run build`
2. Generates templ templates via `templ generate`
3. Compiles Go server to `./tmp/main`
4. Watches for changes in `.go`, `.templ`, `.ts`, `.css` files

### Frontend Operations
All frontend commands run from the `frontend/` directory:
```bash
cd frontend
bun run build      # Build TypeScript and CSS to ../static/
bun run lint       # Lint with Biome
bun run format     # Format with Biome
bun run check      # Lint + format + apply fixes
```

### Templ Generation
When editing `.templ` files, run to regenerate Go code:
```bash
templ generate
```
This creates `*_templ.go` files which are excluded from git.

## Architecture Overview

### Request Flow
1. Chi router (`cmd/server/main.go`) handles HTTP requests
2. Routes render templ components from `views/pages/`
3. Templ components use `views/layouts/Base()` wrapper
4. Base layout includes `/static/scripts/global.js` and `/static/styles/global.css`
5. Client-side TypeScript hydrates interactive features

### Directory Structure

**Backend (Go)**
- `cmd/server/main.go` - HTTP server entrypoint, Chi router setup
- `internal/business/` - Domain models (e.g., `recipe.Recipe`)
- `internal/config/` - Environment config loading (.env support)
- `internal/viewmodel/` - View models and metadata for pages
- `views/layouts/` - Shared templ layouts (e.g., `Base()`)
- `views/pages/` - Page-specific templ components
- `views/sharedcomponents/` - Reusable templ components (e.g., icons)

**Frontend (TypeScript/CSS)**
- `frontend/scripts/global.ts` - Loaded on all pages
- `frontend/scripts/pages/*.ts` - Page-specific scripts
- `frontend/styles/global.css` - Global Tailwind CSS
- `frontend/styles/pages/*.css` - Page-specific styles
- `frontend/build.ts` - Bun build script that bundles all scripts/styles to `static/`

**Build Output**
- `static/` - Generated frontend assets (scripts, styles)
- `tmp/` - Compiled Go binary during development

### Templ Components

Templ is a Go templating engine with syntax like:
```go
package pages

import "masakcook/views/layouts"

templ Index() {
    @layouts.Base(nil) {
        <div>Content</div>
    }
}
```

Key patterns:
- Components are Go functions returning `templ.Component`
- Use `@` to call other templ components
- Use `{ children... }` to accept child content
- Generated `*_templ.go` files are ignored in git
- Templ components can accept Go structs (viewmodels)

### Frontend Build System

The `frontend/build.ts` script:
- Auto-discovers `scripts/pages/*.ts` and `styles/pages/*.css`
- Bundles with code splitting and ESM format
- Outputs to `../static/` with structure matching source
- Uses `bun-plugin-tailwind` for CSS processing
- Controlled by `ENVIRONMENT` env var (development/production)

### Configuration

Server config in `internal/config/config.go` loads from environment variables:
- `PORT` - Server port (default: 8000)
- `ENVIRONMENT` - "development" or "production" (default: development)
- `STATIC_DIR` - Static files directory (default: ./static)

No `.env` file in repo; create one locally if needed.

## Development Patterns

### Adding a New Page
1. Create templ file in `views/pages/newpage.templ`
2. Add route in `cmd/server/main.go`:
   ```go
   r.Get("/newpage", func(w http.ResponseWriter, r *http.Request) {
       templ.Handler(pages.NewPage()).ServeHTTP(w, r)
   })
   ```
3. (Optional) Create `frontend/scripts/pages/newpage.ts` for client-side code
4. (Optional) Create `frontend/styles/pages/newpage.css` for page-specific styles
5. Run `make build` to make sure the build compiles successfully


### Working with Templ
- Always run `templ generate` after editing `.templ` files
- Import other packages in templ with standard Go imports
- Pass data to templates using viewmodel structs
- Use `@` prefix to render child components

### Frontend Assets
- Global scripts/styles are bundled separately from page-specific ones
- Bun automatically handles code splitting and tree-shaking
- Page-specific assets are auto-discovered by filename convention

### HTML Streaming Pattern
The app uses "suspense-like" behavior:
- Static content renders immediately
- Dynamic content streams in afterward
- This is implemented via templ's native streaming support

## Important Notes

- `*_templ.go` files are auto-generated and git-ignored
- Air watches exclude `_test.go`, `_templ.go`, and `.d.ts` files
- Static assets are rebuilt on every Air reload via the build command
- Templ components must be in a package matching their directory name
