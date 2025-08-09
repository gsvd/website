# website

My website, based on [go-website-skeleton](https://github.com/gsvd/go-website-skeleton).

## Dependencies

- [mise](https://github.com/jdx/mise) for managing your Go and Node versions.
- [air](https://github.com/cosmtrek/air) for live reloading the Go server.

## Project Setup

Clone the repository, navigate to its root, then run:
```bash
make setup
```

Create a `.env` file at the project root with the following content:
```txt
ENV=local
HOST=127.0.0.1
PORT=8080
```

## Run the Project

During development use two terminals:

### 1) Terminal 1 — Watch Tailwind CSS
```bash
make watch-css
```

### 2) Terminal 2 — Live-reload Go server
```bash
make dev
```