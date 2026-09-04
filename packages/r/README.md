# marbellaprices (R)

The Marbella Wire open price datasets — the [Sunbed Index][sunbeds], the
[chiringuito census][chi] and the [beach club census][bc] — as data frames.
No hard dependencies.

```r
install.packages("marbellaprices", repos = "https://shiftdylson1.r-universe.dev")
```

```r
library(marbellaprices)

sunbeds <- mw_load("sunbed_index")
sunbeds[!is.na(sunbeds$price_eur), c("venue", "area", "price_eur")]
#>                venue                       area price_eur
#> 1         Ocean Club              Puerto Banús       300
#> 2        Playa Padre Marbella centre (El Cable)       240
#> ...
```

## Functions

| Call | Returns |
|---|---|
| `mw_datasets()` | the three dataset names |
| `mw_load(dataset)` | a data frame, one row per venue |
| `mw_json(dataset)` | the published JSON envelope: season, method, venues (needs **jsonlite**) |
| `mw_info(dataset)` | provenance: canonical page in all six locales, licence, row count, DOI |
| `mw_file(dataset, ext)` | the path to the installed CSV or JSON, to read with a tool of your own |

`dataset` is `"sunbed_index"`, `"chiringuitos"` or `"beach_clubs"`; hyphens and
capitals are accepted, so `"beach-clubs"` works too.

## An empty cell is `NA`, not a zero

A venue that publishes no rate is carried with the figure absent and a `note`
saying why. It is never filled in with zero and never estimated:

```r
sunbeds[is.na(sunbeds$price_eur), c("venue", "note")]
```

A zero that *is* present is a real price — the public beach costs nothing.

## Licence and credit

The R code is MIT ([LICENSE](LICENSE)). The figures are CC BY 4.0
([LICENSE.note](LICENSE.note)): credit "Marbella Wire" and link the canonical
page. `mw_info()` returns the exact attribution string:

```r
mw_info("beach_clubs")$attribution
#> [1] "Marbella Wire, https://marbellawire.com/prices/beach-clubs/"
```

Cite the archived dataset as DOI [10.5281/zenodo.22094846][doi].

## Notes for maintainers

The CSV and JSON files in `inst/extdata/` are copied from the repository root
by `node packages/sync-data.mjs`; they are never edited here. There is no
`data/*.rda` twin, so there is only ever one copy to keep current. The `man/`
pages are written by hand and kept in step with the roxygen comments in
`R/marbellaprices.R`.

[sunbeds]: https://marbellawire.com/prices/sunbeds/
[chi]: https://marbellawire.com/prices/chiringuitos/
[bc]: https://marbellawire.com/prices/beach-clubs/
[doi]: https://doi.org/10.5281/zenodo.22094846
