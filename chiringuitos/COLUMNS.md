# Columns: chiringuitos.csv

| column | type | meaning |
|---|---|---|
| `slug` | string | Stable id; also the venue page path segment |
| `name` | string | Venue name |
| `area` | string | Area of Marbella |
| `beach` | string | Beach the venue sits on, where recorded |
| `business_status` | string | 'operating' at last verification |
| `opened_year` | number | Opening year, where recorded |
| `price_display` | string | Headline price as shown on the page, EUR |
| `price_basis` | string | What the price is read from (e.g. 'menu') |
| `price_kind` | string | What the price buys (e.g. 'espeto', 'mains') |
| `price_note` | string | Caveat on the price |
| `price_level` | string | Relative bracket, € to €€€€ |
| `signature_dish` | string | Signature dish, where one is published |
| `signature_dish_eur` | number | Its menu price, EUR |
| `espeto_eur` | number | Espeto de sardinas menu price, EUR, where published |
| `official_site` | url | Venue's own website |
| `menu_url` | url | Venue's own published menu (menu_is_pdf says if it is a PDF) |
| `menu_is_pdf` | boolean | Whether menu_url is a PDF |
| `address` | string | Street address |
| `phone` | string | Published business phone |
| `google_place_id` | string | Google Place ID, for joining |
| `google_maps_url` | url | Google Maps listing |
| `verified` | date | Last per-venue verification |
| `url` | url | The venue's page on marbellawire.com |
| `currency` | string | All prices are EUR |
| `source_page` | url | Canonical dataset page |
| `attribution` | string | Requested credit line |
| `license` | url | CC BY 4.0 |

Empty cell = not published by the venue, never zero and never a guess.
