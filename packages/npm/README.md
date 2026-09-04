# marbella-prices

[![npm](https://img.shields.io/npm/v/marbella-prices)](https://www.npmjs.com/package/marbella-prices)
[![node](https://img.shields.io/node/v/marbella-prices)](https://www.npmjs.com/package/marbella-prices)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.22094846.svg)](https://doi.org/10.5281/zenodo.22094846)
![Code: MIT](https://img.shields.io/badge/code-MIT-green)
![Data: CC BY 4.0](https://img.shields.io/badge/data-CC_BY_4.0-blue)
![Season](https://img.shields.io/badge/season-2026-informational)

Open price datasets from [Marbella Wire](https://marbellawire.com/prices/), an independent price
tracker for Marbella, Spain, packaged for Node: the **Sunbed Index**, a **chiringuito census**
and a **beach club price census**. ESM, zero dependencies, typed; CSV and JSON are embedded,
nothing is fetched at runtime.

```sh
npm install marbella-prices
```

```js
import { load, info } from 'marbella-prices';

const rows = await load('sunbed_index');   // [{ venue, area, status, price_eur, ... }]
const priced = rows.filter((r) => r.status === 'priced');
const cheapest = priced.reduce((a, b) => (a.price_eur < b.price_eur ? a : b));

(await info('chiringuitos')).pages.es;     // 'https://marbellawire.com/es/prices/chiringuitos/'
```

Also exported: `loadSync(name)`, `loadJson(name)` (the dataset envelope: season, update date,
method, venues), `datasets()`, `parseCsv(text)`, `version`. Raw files are reachable as
`marbella-prices/data/sunbed-index.csv` etc. Cells are typed: numbers become numbers,
`true`/`false` become booleans, an empty cell (price not published) becomes `null`.

## Datasets and canonical pages

| Dataset | `load()` name | en | es | de | fr | nl | sv |
|---|---|---|---|---|---|---|---|
| Sunbed Index | `sunbed_index` | [en](https://marbellawire.com/prices/sunbeds/) | [es](https://marbellawire.com/es/prices/sunbeds/) | [de](https://marbellawire.com/de/prices/sunbeds/) | [fr](https://marbellawire.com/fr/prices/sunbeds/) | [nl](https://marbellawire.com/nl/prices/sunbeds/) | [sv](https://marbellawire.com/sv/prices/sunbeds/) |
| Chiringuito census | `chiringuitos` | [en](https://marbellawire.com/prices/chiringuitos/) | [es](https://marbellawire.com/es/prices/chiringuitos/) | [de](https://marbellawire.com/de/prices/chiringuitos/) | [fr](https://marbellawire.com/fr/prices/chiringuitos/) | [nl](https://marbellawire.com/nl/prices/chiringuitos/) | [sv](https://marbellawire.com/sv/prices/chiringuitos/) |
| Beach club census | `beach_clubs` | [en](https://marbellawire.com/prices/beach-clubs/) | [es](https://marbellawire.com/es/prices/beach-clubs/) | [de](https://marbellawire.com/de/prices/beach-clubs/) | [fr](https://marbellawire.com/fr/prices/beach-clubs/) | [nl](https://marbellawire.com/nl/prices/beach-clubs/) | [sv](https://marbellawire.com/sv/prices/beach-clubs/) |

All prices hub: [en](https://marbellawire.com/prices/) · [es](https://marbellawire.com/es/prices/) · [de](https://marbellawire.com/de/prices/) · [fr](https://marbellawire.com/fr/prices/) · [nl](https://marbellawire.com/nl/prices/) · [sv](https://marbellawire.com/sv/prices/).
Always-current raw endpoints: [sunbed-index.csv](https://marbellawire.com/data/sunbed-index.csv) · [sunbed-index.json](https://marbellawire.com/data/sunbed-index.json) · [sunbed-index.svg](https://marbellawire.com/data/sunbed-index.svg) (embeddable chart).

Column dictionaries: [sunbed-index](https://github.com/shiftdylson1/marbella-price-data/blob/main/sunbed-index/COLUMNS.md) ·
[chiringuitos](https://github.com/shiftdylson1/marbella-price-data/blob/main/chiringuitos/COLUMNS.md) ·
[beach-clubs](https://github.com/shiftdylson1/marbella-price-data/blob/main/beach-clubs/COLUMNS.md).

## Citing

Version-independent DOI (Zenodo): [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846).

```bibtex
@dataset{marbellawire_prices,
  author    = {{Marbella Wire}},
  title     = {Marbella price data},
  year      = {2026},
  publisher = {Zenodo},
  doi       = {10.5281/zenodo.22094846},
  url       = {https://marbellawire.com/prices/}
}
```

In prose: "Marbella Wire, https://marbellawire.com/prices/sunbeds/".

## Licence

- Loader code: [MIT](LICENSE).
- Data (`data/*.csv`, `data/*.json`): [CC BY 4.0](LICENSE-DATA), copyright Marbella Wire. Credit "Marbella Wire" and link the canonical page above. Figures only: no editorial text, no photographs, no Google Places ratings or coordinates are included.

## Method

Prices are read from each venue's own menu, rate card or booking system, never from aggregators; every row carries its source URL, the verification date, the attribution line and the licence. One number in the source is one number here: an empty cell means "not published", never zero and never an estimate. Refreshed once a season from the [source repository](https://github.com/shiftdylson1/marbella-price-data); the live pages update first.

## Corrections

Wrong price, closed venue, dead link: [open a correction issue](https://github.com/shiftdylson1/marbella-price-data/issues/new?template=price-correction.yml) with a source URL from the venue itself. Contact: editor@marbellawire.com.
