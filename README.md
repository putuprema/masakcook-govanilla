# MasakCook - Recipe Sharing Web App

A modern, crowd-sourced recipe web application built with Golang, Vanilla Web Stack (HTML/CSS/TS), and Tailwind CSS v4.

This is a recreation of the same app previously built with Next.js. Check out the repo of the Next.js version [here](https://github.com/putuprema/masakcook).

## Features

### 🏠 Homepage

- **Hero Section** - Showcases the "Recipe of the Day" with full-width image banner
- **Trending Now** - Horizontal scrollable carousel of trending recipes
- **Search Section** - Real-time recipe search with keyword and category filters

### 🔖 Bookmark System

- Save/unsave recipes with persistent localStorage
- Navbar badge showing saved recipe count (displays "99+" for 100+)
- Modal popup to view all saved recipes with thumbnails
- Empty state with helpful message when no recipes saved
- Click-outside-to-close and Escape key support for modal
- Individual unsave buttons for each saved recipe

### 🎨 UI/UX

- Glassmorphism navbar with backdrop blur effect
- Responsive mobile-first design
- Suspense-like behavior on page loading (loading static element first, and then later streaming in dynamic elements)
- Smooth animations and transitions
- Accessibility-focused (ARIA labels, keyboard navigation)

## Tech Stack

- **Server-side:** Golang 1.25 with net/http, Chi (routing engine), and templ (templating engine)
- **Client-side scripting and styling:** Typescript 5 and Tailwind CSS v4 with Bun as package manager and bundler
- **UI Icons:** Heroicons SVG as Go templ components
- **Linter/Formatter:** Biome (client-side), golangci-lint (server-side)
- **Development Environment**: Air for Go live reload

## Project Structure

```
masakcook/
├── cmd/
│   └── server/
│       └── main.go             # Main entry point for the HTTP web server
├── frontend/                   # Client-side scripting and styling all goes here
│   ├── scripts/                # Client-side Typescript files
│   ├── styles/                 # Client-side CSS files
│   └── build.ts                # Entry point for the Typescript and CSS bundler (output of this file is saved in projectdir/static/)
├── internal/
│   ├── business/               # Server-side business logic
│   ├── config/                 # Config related
│   └── viewmodel/              # Page view model
├── views/
│   ├── layouts/                # Shared layout files
│   ├── pages/                  # Page specific files
│   └── sharedcomponents/       # Shared templ components
├── .air.toml                   # Air config for live reload
├── Makefile                    # Collection of commands for building, dev server, linting, etc (uses make tool)
└── go.mod                      # Main go module file
```

## Getting Started

### Prerequisites

- Go 1.25+
- templ CLI (`go install github.com/a-h/templ/cmd/templ@latest`)
- air CLI (`go install github.com/air-verse/air@latest`)
- Bun 1.3+

### Available Scripts

Refer to the `Makefile` and `frontend/package.json` files

## Key Features Implementation

### Recipe Search API

The search functionality uses a dedicated API route (`/api/recipes/search`) that:

- Accepts `keyword` and `category` query parameters
- Validates input and returns proper HTTP status codes
- Simulates database query with artificial delays
- Returns filtered recipes in JSON format

### Glassmorphism Navbar

Modern glass effect using:

- `backdrop-blur-md` for blur effect
- `bg-white/70` for semi-transparency
- Sticky positioning that stays on top while scrolling
- Bookmark button with count badge (red pill badge showing saved recipes)
- Opens SavedRecipesModal on click with smooth animations

### SavedRecipesModal Component

A dropdown-style modal for managing saved recipes:

- Positioned at top-right below navbar (`right-4 top-20`)
- Displays recipe thumbnails, titles, categories, and total time
- Empty state with HeartIcon when no recipes saved
- Individual unsave buttons (red heart icons) for each recipe
- Closes via Escape key, backdrop click, or X button
- Scroll support for long lists (max-height with overflow-y-auto)
- Uses event listeners for click-outside detection

## Best Practices

- ✅ **Golang, Chi, Templ** - Latest patterns and features, HTML streaming for dynamic data
- ✅ **TypeScript** - Full type safety
- ✅ **Accessibility** - ARIA labels, semantic HTML
- ✅ **Code Quality** - Biome & Golang linting, consistent formatting
- ✅ **Responsive** - Mobile-first design
