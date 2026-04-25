# fastbuild

- **Project ID:** `fast-bld`
- **Live:** https://fast-bld.web.app · https://fast-bld.firebaseapp.com
- **Stack:** htmx (frontend) · Go Cloud Function Gen 2 (`/api/**`) · Firestore · Firebase Hosting
- **Region:** `us-central1` (function + Firestore, single-region)
- **Go module path:** `example.com/fastbuild` (placeholder — change in `go.mod` if you want)
- **Endpoints:** `GET /api/hello` · `POST /api/visit`

## Layout

- `public/` — static HTML/CSS served by Hosting
- `functions/` — Go function (`function.go`, `go.mod`)
- `firebase.json` — hosting config + `/api/**` rewrite to function
- `.firebaserc` — default project

## Deploy

Firebase CLI doesn't ship a Go runtime, so the function goes via `gcloud`; Hosting rewrites `/api/**` to it.

```sh
# Function
gcloud functions deploy api \
  --gen2 --runtime=go126 --region=us-central1 \
  --source=./functions --entry-point=api \
  --trigger-http --allow-unauthenticated

# Hosting
firebase deploy --only hosting
```

## One-time setup

```sh
npm i -g firebase-tools
brew install --cask google-cloud-sdk
firebase login && gcloud auth login
gcloud config set project fast-bld
cd functions && go mod tidy
```

## Local dev

- Hosting + rewrites: `firebase emulators:start --only hosting`
- Firestore emulator: set `FIRESTORE_EMULATOR_HOST=localhost:8080` before running the function
