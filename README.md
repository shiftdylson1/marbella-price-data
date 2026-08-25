# Marbella price data

![License: CC BY 4.0](https://img.shields.io/badge/license-CC_BY_4.0-blue)
![Verified](https://img.shields.io/badge/data_verified-2026--08--17-brightgreen)
![Season](https://img.shields.io/badge/season-2026-informational)

Open datasets from [Marbella Wire](https://marbellawire.com), an independent price tracker
for Marbella, Spain. Every figure here is published on marbellawire.com
first; this repo is the machine-readable copy, regenerated from the same
data files that render the pages.

| Dataset | What it is | Rows | Canonical page |
|---|---|---|---|
| [sunbed-index](sunbed-index/) | High-season price of two sunbeds plus minimum spend at Marbella beach venues, season 2026 | 8 priced + 4 unpublished | https://marbellawire.com/prices/sunbeds/ |
| [chiringuitos](chiringuitos/) | Census of Marbella chiringuitos (beach restaurants) with menu-sourced prices, season 2026 | 43 venues | https://marbellawire.com/prices/chiringuitos/ |
| [beach-clubs](beach-clubs/) | Beach club sunbed, bed and minimum-spend prices, season 2026 | 14 clubs | https://marbellawire.com/prices/beach-clubs/ |

Each folder holds the data as JSON and CSV plus a PROVENANCE.md stating how
it was collected, when it was last verified, and what was deliberately left
out, a COLUMNS.md data dictionary, and (for the censuses) a DATA.md with one
anchored section per venue so you can deep-link a single row. All prices are
euros; every CSV row carries its unit, source page and licence so a row
quoted on its own stays traceable.

## Use it in ten seconds

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

```sql
-- DuckDB
SELECT name, sunbed_price FROM 'https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv' ORDER BY 2 DESC;
```

```sh
# Or skip the repo and hit the always-current live endpoint
curl https://marbellawire.com/data/sunbed-index.csv
```

## Embed the Sunbed Index

The live index is also served as an image that updates with the data:

```html
<a href="https://marbellawire.com/prices/sunbeds/"><img src="https://marbellawire.com/data/sunbed-index.svg" alt="Marbella Wire Sunbed Index" width="600" height="315"></a>
```

## Citing

See [CITATION.cff](CITATION.cff) (GitHub's "Cite this repository" button
uses it), or in prose: "Marbella Wire, https://marbellawire.com/prices/sunbeds/".

## License

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) (see [LICENSE](LICENSE)). Free to use, share
and adapt with credit to **Marbella Wire** and a link to the dataset's
canonical page (table above). The grant covers the figures in this repo
only; the site's editorial text and its credited photographs are not
included here and are not CC-licensed.

## Sources and method

Prices come from venues' own published menus, rate cards and booking
systems, never from third-party aggregators; each row links its source.
One number in the source is one number here: nothing is estimated or
interpolated. Per-dataset detail is in each PROVENANCE.md.

## Updates

Regenerated manually, roughly once a season, from the Marbella Wire data
pipeline. The live pages update first; if a figure here disagrees with
marbellawire.com, the site is current and this repo is awaiting its
seasonal refresh.
