package fastbuild

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("api", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	switch path {
	case "/hello":
		hello(w, r)
	case "/visit":
		visit(w, r)
	default:
		http.NotFound(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<p>hello from go — %s</p>`, time.Now().UTC().Format(time.RFC3339))
}

func visit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c, err := fsClient(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doc := c.Collection("counters").Doc("visits")
	if _, err := doc.Set(ctx, map[string]any{
		"count":     firestore.Increment(1),
		"updatedAt": firestore.ServerTimestamp,
	}, firestore.MergeAll); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	snap, err := doc.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, _ := snap.DataAt("count")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<p>visits: %v</p>`, count)
}

var (
	fsOnce sync.Once
	fs     *firestore.Client
	fsErr  error
)

func fsClient(ctx context.Context) (*firestore.Client, error) {
	fsOnce.Do(func() {
		projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if projectID == "" {
			projectID = os.Getenv("GCP_PROJECT")
		}
		if projectID == "" {
			projectID = firestore.DetectProjectID
		}
		fs, fsErr = firestore.NewClient(ctx, projectID)
	})
	return fs, fsErr
}
