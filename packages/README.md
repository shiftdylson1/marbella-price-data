# Packages

The datasets in this repo, installable:

| Registry | Package | Install | Source |
|---|---|---|---|
| PyPI | `marbella-prices` | `pip install marbella-prices` | [python/](python/) |
| npm | `marbella-prices` | `npm install marbella-prices` | [npm/](npm/) |

Both embed the three CSV/JSON files from the repo root and expose one call,
`load("sunbed_index")`, returning plain rows. Loader code is MIT; the data
stays CC BY 4.0 (credit "Marbella Wire" + the canonical page).

Versions are CalVer (`2026.9.4` = the day the data was packaged). After each
seasonal export, `node packages/sync-data.mjs --version=YYYY.M.D` copies the
data in and stamps the version; a `packages-vYYYY.M.D` tag then publishes both
through [`.github/workflows/packages.yml`](../.github/workflows/packages.yml)
(PyPI Trusted Publishing + npm provenance, no stored tokens).
