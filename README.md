# 🛠️ YesFix – The Gamusa of Services

YesFix is a service company aimed at providing all home and office services under a single name — electricians, plumbers, repair professionals, and cleaners.

Built with 💛 in Assam.

---

## Live Website

> Coming soon at [https://yesfix.in](https://yesfix.in)

---

## Tech Stack

- **Frontend**: [Templ](https://templ.guide) (Go templating), Tailwind CSS, HTMX
- **Backend**: Go, Chi router, `net/http`
- **Database**: PostgreSQL with SQLC
- **Email**: Mailtrap API

---

## Folder Structure

```txt
.
├── bin/               # Compiled binaries (usually ignored in Git)
├── dbConfig/          # DB connection setup, queries, migrations
├── handlers/          # all route handlers 
├── initializers/      # env, logging
├── public/            # Static assets (CSS, JS, fonts, images)
├── routes/            # HTTP route definitions using chi
├── tmp/               # Temporary files (often .air.tmp etc.)
├── types/             # Custom types and structs (e.g., App, FormData)
├── utils/             # Helper functions (validation, rendering)
├── views/             # `.templ` templates (HTML layout and pages)
│
├── .air.toml          # Air config for live reload
├── .gitignore         # Git ignored files(ignore the stuff inside)
├── Makefile           # Dev & build tasks
├── README.md          # You’re here
├── Sitemap.txt        # SEO sitemap
├── go.mod / go.sum    # Go dependencies
├── main.go            # Entry point
│
├── postcss.config.js  # PostCSS config for Tailwind
├── static_dev.go      # Static file handler (dev)
├── static_prod.go     # Static file handler (prod)
├── tailwind.config.js # Tailwind customization
├── package.json       # Node dependencies for Tailwind
├── package-lock.json  # Node lock file

---

