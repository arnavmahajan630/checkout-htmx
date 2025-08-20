# GOTTH Template 🚀

A minimal production-ready starter template for building web apps with:

- **Go** – fast, simple, and efficient backend  
- **Templ** – type-safe HTML templating for Go  
- **TailwindCSS** – utility-first CSS for styling  
- **HTMX** – modern, declarative AJAX and interactivity  
- **Alpine.js** – lightweight frontend reactivity

This project is designed to help you quickly bootstrap a Go web project with a clean structure and modern tooling.

---

## ✨ Features

- ✅ Go HTTP server with clean project structure  
- ✅ Integrated **Templ** layouts & components  
- ✅ Preconfigured **TailwindCSS** build pipeline  
- ✅ Static assets served in both **dev** and **prod** modes  
- ✅ Minimal **HTMX + Alpine.js** setup for interactivity  
- ✅ Single binary deployment (Go `embed` for static files)  

---

## 📂 Project Structure

```
.
├── layouts/       # Base templates
├── views/         # Page-level templates
├── public/        # Static assets (CSS, JS, images)
├── cmd/app        # Go server code
└── tailwind.config.js
```

---

## ⚡ Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/yourusername/gotth-template.git
cd gotth-template
```

### 2. Install dependencies
Make sure you have **Go 1.21+**, **Node.js**, and **npm** installed.

```bash
npm install
```

### 3. Run Tailwind in watch mode
```bash
npx tailwindcss -i ./public/input.css -o ./public/styles.css --watch
```

### 4. Run the Go server
```bash
go run ./src/main.go
```

Visit **http://localhost:8080** 🚀

---

## 🛠 Production Build

Bundle Tailwind before building:
```bash
npx tailwindcss -i ./public/input.css -o ./public/styles.css --minify
```

Build the Go binary (with static assets embedded):
```bash
go build -o app ./src
./app
```

---

## 📌 Notes

- Use **Templ** layouts for reusability (`@layout.base { ... }`)  
- `embed.FS` is used in production to serve `public/` files from a single binary  
- In dev mode, files are served directly from disk  

---

## 🤝 Contributing

Feel free to fork this repo, open issues, or submit PRs.  
This is just a starter – you can extend it with DB integration, authentication, or APIs.

---

## 📜 License

MIT
