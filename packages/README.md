# Packages

The datasets in this repo, installable:

| Registry | Package | Install | Source |
|---|---|---|---|
| PyPI | `marbella-prices` | `pip install marbella-prices` | [python/](python/) |
| npm | `marbella-prices` | `npm install marbella-prices` | [npm/](npm/) |
| pkg.go.dev | `.../packages/go` | `go get github.com/shiftdylson1/marbella-price-data/packages/go` | [go/](go/) |
| R-universe | `marbellaprices` | `install.packages("marbellaprices", repos = "https://shiftdylson1.r-universe.dev")` | [r/](r/) |

All four embed the three CSV/JSON files from the repo root and expose one call
returning plain rows — `load("sunbed_index")` in Python and JavaScript,
`Load("sunbed_index")` in Go, `mw_load("sunbed_index")` in R. Loader code is
MIT; the data stays CC BY 4.0 (credit "Marbella Wire" + the canonical page).

Versions are CalVer (`2026.9.4` = the day the data was packaged). After each
seasonal export, `node packages/sync-data.mjs --version=YYYY.M.D` copies the
data into all four trees and stamps the version.

Publishing differs by ecosystem:

- **PyPI and npm** are pushed: a `packages-vYYYY.M.D` tag publishes both through
  [`.github/workflows/packages.yml`](../.github/workflows/packages.yml)
  (PyPI Trusted Publishing + npm provenance, no stored tokens).
- **pkg.go.dev** pulls. It indexes the module once a tag exists; because the
  module sits in a subdirectory the tag must be `packages/go/vX.Y.Z`, and the Go
  module keeps its own semver (`Version`) separate from the CalVer data stamp
  (`DataVersion`).
- **R-universe** pulls too, building the default branch about once an hour. It
  needs no tag — only that the package name be listed in the registry repo
  `shiftdylson1/shiftdylson1.r-universe.dev` with `"subdir": "packages/r"`, and
  that the r-universe GitHub App be installed on the account.
