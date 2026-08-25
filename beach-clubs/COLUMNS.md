# Columns: beach-clubs.csv

| column | type | meaning |
|---|---|---|
| `slug` | string | Stable id; also the club page path segment |
| `name` | string | Club name |
| `area` | string | Area of Marbella |
| `season_year` | number | Season the prices belong to |
| `sunbed_price` | number | Cheapest sunbed/lounger, EUR, high season, per day |
| `bed_or_vip_price` | number | Top bed or VIP package, EUR, per day, where published |
| `minimum_spend` | string | Minimum-consumption terms, as published |
| `entry_policy` | string | How entry actually works |
| `price_level` | string | Relative bracket |
| `source_url` | url | The club's own rate card or booking page the prices were read from |
| `rates_url` | url | Published rates document (rates_is_pdf says if PDF) |
| `rates_is_pdf` | boolean | Whether rates_url is a PDF |
| `booking_url` | url | Club's booking system |
| `official_site` | url | Club's own website |
| `menu_url` | url | Published menu, where one exists |
| `menu_is_pdf` | boolean | Whether menu_url is a PDF |
| `opened_year` | number | Opening year, where recorded |
| `business_status` | string | 'operating' at last verification |
| `address` | string | Street address |
| `phone` | string | Published business phone |
| `google_place_id` | string | Google Place ID, for joining |
| `google_maps_url` | url | Google Maps listing |
| `verified` | date | Last per-venue verification |
| `notes` | string | Sourcing note |
| `url` | url | The club's page on marbellawire.com |
| `currency` | string | All prices are EUR |
| `source_page` | url | Canonical dataset page |
| `attribution` | string | Requested credit line |
| `license` | url | CC BY 4.0 |

Empty cell = not published by the venue, never zero and never a guess.
