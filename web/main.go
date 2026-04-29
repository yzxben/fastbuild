package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/a-h/templ"
	"github.com/yzxben/fastbuild/web/templates"
)

//go:embed static
var staticFS embed.FS

func main() {
	var clicks atomic.Int64

	static, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))

	mux.Handle("GET /{$}", templ.Handler(templates.Index()))

	mux.HandleFunc("POST /click", func(w http.ResponseWriter, r *http.Request) {
		n := clicks.Add(1)
		templates.Count(int(n)).Render(r.Context(), w)
	})

	addr := ":" + env("PORT", "8080")
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
