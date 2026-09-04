# marbella-prices (Go)

The Marbella Wire open price datasets — the [Sunbed Index][sunbeds], the
[chiringuito census][chi] and the [beach club census][bc] — embedded in your
binary with `go:embed`. No network at run time, no dependencies outside the
standard library.

```sh
go get github.com/shiftdylson1/marbella-price-data/packages/go
```

```go
import marbellaprices "github.com/shiftdylson1/marbella-price-data/packages/go"

rows, err := marbellaprices.Sunbeds()
for _, r := range rows {
    if r.PriceEUR != nil {
        fmt.Printf("%-24s %6.0f EUR  (%s)\n", r.Venue, *r.PriceEUR, r.Area)
    }
}
```

## API

| Call | Returns |
|---|---|
| `Sunbeds()` | `[]SunbedRow` — the Sunbed Index |
| `Chiringuitos()` | `[]ChiringuitoRow` — the chiringuito census |
| `BeachClubs()` | `[]BeachClubRow` — the beach club census |
| `Load(name)` | `[]Row` (a `map[string]any` per row), the shape the Python and npm packages return |
| `LoadJSON(name, &v)` | the dataset's JSON envelope: season, updated, method, venues |
| `Raw(name, "csv"\|"json")` | the embedded bytes, verbatim |
| `Info(name)` | `DatasetInfo`: canonical page in all six locales, licence, row count, season, DOI |
| `Names()`, `Lookup(name)` | the dataset keys, and name resolution |

`name` is `"sunbed_index"`, `"chiringuitos"` or `"beach_clubs"`; hyphens, spaces
and capitals are accepted, so `"beach-clubs"` works too.

## An empty cell is not a zero

A venue that publishes no rate is carried with the figure absent and a `note`
saying why — it is never filled in with zero and never guessed. Every optional
number is therefore a pointer and stays `nil`:

```go
for _, r := range rows {
    if r.PriceEUR == nil {
        fmt.Printf("%s publishes no rate: %s\n", r.Venue, r.Note)
    }
}
```

A zero that *is* present is a real price: the public beach costs nothing.

## Licence and credit

The loader code is MIT ([LICENSE](LICENSE)). The figures are CC BY 4.0
([LICENSE-DATA](LICENSE-DATA)): credit "Marbella Wire" and link the canonical
page. `Info(name)` hands you the exact attribution string and URL to use.

Cite the archived dataset as DOI [10.5281/zenodo.22094846][doi].

## Versions

The module is versioned with semver (`Version`), while `DataVersion` carries
the CalVer stamp of the embedded snapshot — `2026.9.4` is the day the data was
packaged, and matches the PyPI and npm releases of the same snapshot. The
embedded copy is a seasonal snapshot; the live endpoints at
<https://marbellawire.com/data/> are always current.

Because this module lives in a subdirectory, its release tags carry the
directory prefix: `packages/go/v0.1.0`.

[sunbeds]: https://marbellawire.com/prices/sunbeds/
[chi]: https://marbellawire.com/prices/chiringuitos/
[bc]: https://marbellawire.com/prices/beach-clubs/
[doi]: https://doi.org/10.5281/zenodo.22094846
