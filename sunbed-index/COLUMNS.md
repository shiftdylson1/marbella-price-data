# Columns: sunbed-index.csv

| column | type | meaning |
|---|---|---|
| `venue` | string | Venue name as published on the index |
| `area` | string | Area of Marbella |
| `status` | string | 'priced' or 'unpublished' (listed but deliberately not priced; see note) |
| `price_eur` | number | Two sunbeds plus minimum spend, EUR, high season, weekday. EMPTY for unpublished rows, never zero |
| `price_unit` | string | Restates the unit for the row |
| `yoy_pct` | number | Year-on-year change of price_eur, percent |
| `note` | string | Pricing caveat, or the reason a row is unpublished |
| `season` | number | Season year |
| `measured_on` | date | When the price was read (unpublished rows: when checked) |
| `price_source_url` | url | The venue's own price source |
| `price_source_via` | string | Booking platform name when the source is the venue's booking system |
| `index_average_eur` | number | Index-wide average, repeated on every row so a row travels with its headline |
| `index_average_unit` | string | Unit of the average |
| `index_yoy_pct` | number | Index-wide year-on-year change, percent |
| `source_url` | url | Canonical index page |
| `attribution` | string | Requested credit line |
| `license` | url | CC BY 4.0 |

Empty cell = not published by the venue, never zero and never a guess.
