# Prisdata för Marbella

Läs på: [Español](README.es.md) · [Deutsch](README.de.md) · [Français](README.fr.md) · [Nederlands](README.nl.md) · [English](../../README.md)

Öppna dataset från [Marbella Wire](https://marbellawire.com), en oberoende prisbevakare för Marbella, Spanien. Varje siffra publiceras först på marbellawire.com; det här arkivet är den maskinläsbara kopian, återskapad från samma datafiler som bygger sidorna. Den här sidan är en översättning av den [engelska README-filen](../../README.md), som är referensversionen.

| Dataset | Vad det är | Rader | Kanonisk sida |
|---|---|---|---|
| [Marbella Wire Solstolsindex](../../sunbed-index/) | Högsäsongspris för två solstolar plus minimikonsumtion på Marbellas beachklubbar, säsong 2026 | 8 med pris + 4 utan publicerat pris | https://marbellawire.com/prices/sunbeds/ |
| [Marbella chiringuito-inventering](../../chiringuitos/) | Inventering av Marbellas chiringuitos (strandrestauranger) med priser hämtade från menyn, säsong 2026 | 43 ställen | https://marbellawire.com/prices/chiringuitos/ |
| [Prisinventering av Marbellas beachklubbar](../../beach-clubs/) | Priser för solstol, bädd och minimikonsumtion på beachklubbarna, säsong 2026 | 14 klubbar | https://marbellawire.com/prices/beach-clubs/ |

Alla priser är i euro; varje CSV-rad bär sin enhet, källsida och licens. En tom cell betyder att stället inte publicerar någon siffra, aldrig noll och aldrig en gissning.

Nyckelord: Marbella, Spanien, priser, turism, beachklubbar, chiringuitos, solstolar, öppna data, Costa del Sol, prisindex, dataset

## Kom igång på tio sekunder

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

Eller hoppa över arkivet och läs den alltid aktuella live-ändpunkten:

```sh
curl https://marbellawire.com/data/sunbed-index.csv
```

## Kolumnordlista

Kolumnnamnen är på engelska i alla filer (de är identifierare); betydelserna nedan är på Svenska.

### `sunbed-index/sunbed-index.csv`

**Marbella Wire Solstolsindex.** Vad en dag på solstolen kostar på Costa del Sol: två solstolar plus minimikonsumtion på varje stor beachklubb i Marbella som publicerar ett pris. Sommaren 2026, 8 med pris + 4 utan publicerat pris.

Nyckelord: solstol priser Marbella, solstolar pris Marbella, solsäng pris Marbella, solstolsindex, beachklubb priser Marbella, solstol Puerto Banús, Costa del Sol priser, öppna data, Marbella, Spanien, priser, turism, beachklubbar, chiringuitos, solstolar, Costa del Sol, prisindex, dataset

| kolumn | typ | betydelse |
|---|---|---|
| `venue` | text | Ställets namn som det publiceras i indexet |
| `area` | text | Område i Marbella |
| `status` | text | 'priced' (med pris) eller 'unpublished' (listat men medvetet utan pris; se note) |
| `price_eur` | tal | Två solstolar plus minimikonsumtion, EUR, högsäsong, vardag. TOM för unpublished-rader, aldrig noll |
| `price_unit` | text | Upprepar radens enhet |
| `yoy_pct` | tal | Förändring av price_eur jämfört med föregående år, i procent |
| `note` | text | Förbehåll om priset, eller skälet till att en rad saknar pris |
| `season` | tal | Säsongsår |
| `measured_on` | datum | Datum då priset lästes av (unpublished-rader: kontrolldatum) |
| `price_source_url` | URL | Ställets egen priskälla |
| `price_source_via` | text | Bokningsplattformens namn när källan är ställets bokningssystem |
| `index_average_eur` | tal | Genomsnitt för hela indexet, upprepat på varje rad så att en rad bär med sig sin rubriksiffra |
| `index_average_unit` | text | Genomsnittets enhet |
| `index_yoy_pct` | tal | Förändring för hela indexet jämfört med föregående år, i procent |
| `source_url` | URL | Kanonisk indexsida |
| `attribution` | text | Önskad källhänvisning |
| `license` | URL | CC BY 4.0 |

### `chiringuitos/chiringuitos.csv`

**Marbella chiringuito-inventering.** Inventering av Marbellas chiringuitos (strandrestauranger) med priser hämtade från menyn: espeto de sardinas och signaturrätt per ställe, varje rad länkad till ställets egen meny. Säsong 2026, 43 ställen.

Nyckelord: chiringuito priser Marbella, strandrestaurang priser Marbella, espeto pris, restaurangpriser Marbella, chiringuitos dataset, Costa del Sol, öppna data, Marbella, Spanien, priser, turism, beachklubbar, chiringuitos, solstolar, prisindex, dataset

| kolumn | typ | betydelse |
|---|---|---|
| `slug` | text | Stabilt id; även sökvägssegmentet för ställets sida |
| `name` | text | Ställets namn |
| `area` | text | Område i Marbella |
| `beach` | text | Strand där stället ligger, där det är registrerat |
| `business_status` | text | 'operating' (i drift) vid senaste kontrollen |
| `opened_year` | tal | Öppningsår, där det är registrerat |
| `price_display` | text | Riktpris som det visas på sidan, EUR |
| `price_basis` | text | Vad priset läses från (t.ex. 'menu', menyn) |
| `price_kind` | text | Vad priset köper (t.ex. 'espeto', 'mains' = huvudrätter) |
| `price_note` | text | Förbehåll om priset |
| `price_level` | text | Relativ prisklass, € till €€€€ |
| `signature_dish` | text | Signaturrätt, där en är publicerad |
| `signature_dish_eur` | tal | Dess menypris, EUR |
| `espeto_eur` | tal | Menypris för espeto de sardinas, EUR, där det är publicerat |
| `official_site` | URL | Ställets egen webbplats |
| `menu_url` | URL | Meny publicerad av stället självt (menu_is_pdf anger om den är en PDF) |
| `menu_is_pdf` | boolesk | Om menu_url är en PDF |
| `address` | text | Gatuadress |
| `phone` | text | Publicerat telefonnummer till verksamheten |
| `google_place_id` | text | Google Place ID, för sammankoppling |
| `google_maps_url` | URL | Google Maps-post |
| `verified` | datum | Senaste kontroll per ställe |
| `url` | URL | Ställets sida på marbellawire.com |
| `currency` | text | Alla priser är i EUR |
| `source_page` | URL | Datasetets kanoniska sida |
| `attribution` | text | Önskad källhänvisning |
| `license` | URL | CC BY 4.0 |

### `beach-clubs/beach-clubs.csv`

**Prisinventering av Marbellas beachklubbar.** Verifierade priser för solstol, bädd och minimikonsumtion på beachklubbarna i Marbella och Puerto Banús, avlästa från varje klubbs egen prislista eller bokningssida. Säsong 2026, 14 klubbar.

Nyckelord: beachklubb priser Marbella, beach club Puerto Banús priser, solstol pris, minimikonsumtion beachklubb, Marbella beachklubbar dataset, öppna data, Marbella, Spanien, priser, turism, beachklubbar, chiringuitos, solstolar, Costa del Sol, prisindex, dataset

| kolumn | typ | betydelse |
|---|---|---|
| `slug` | text | Stabilt id; även sökvägssegmentet för klubbens sida |
| `name` | text | Klubbens namn |
| `area` | text | Område i Marbella |
| `season_year` | tal | Säsong som priserna gäller |
| `sunbed_price` | tal | Billigaste solstol eller solsäng, EUR, högsäsong, per dag |
| `bed_or_vip_price` | tal | Bästa bädd eller VIP-paket, EUR, per dag, där det är publicerat |
| `minimum_spend` | text | Villkor för minimikonsumtion, som de publiceras |
| `entry_policy` | text | Hur inträdet fungerar i praktiken |
| `price_level` | text | Relativ prisklass |
| `source_url` | URL | Klubbens egen prislista eller bokningssida som priserna lästes från |
| `rates_url` | URL | Publicerat prisdokument (rates_is_pdf anger om PDF) |
| `rates_is_pdf` | boolesk | Om rates_url är en PDF |
| `booking_url` | URL | Klubbens bokningssystem |
| `official_site` | URL | Klubbens egen webbplats |
| `menu_url` | URL | Publicerad meny, där en finns |
| `menu_is_pdf` | boolesk | Om menu_url är en PDF |
| `opened_year` | tal | Öppningsår, där det är registrerat |
| `business_status` | text | 'operating' (i drift) vid senaste kontrollen |
| `address` | text | Gatuadress |
| `phone` | text | Publicerat telefonnummer till verksamheten |
| `google_place_id` | text | Google Place ID, för sammankoppling |
| `google_maps_url` | URL | Google Maps-post |
| `verified` | datum | Senaste kontroll per klubb |
| `notes` | text | Anteckning om källan |
| `url` | URL | Klubbens sida på marbellawire.com |
| `currency` | text | Alla priser är i EUR |
| `source_page` | URL | Datasetets kanoniska sida |
| `attribution` | text | Önskad källhänvisning |
| `license` | URL | CC BY 4.0 |

## Källor och metod

Priserna kommer från ställenas egna publicerade menyer, prislistor och bokningssystem, aldrig från aggregatorer; varje rad länkar till sin källa. En siffra i källan är en siffra här: inget uppskattas eller interpoleras. Detaljer per dataset finns i varje mapps PROVENANCE.md (på engelska).

## Uppdateringar

Återskapas manuellt, ungefär en gång per säsong, från Marbella Wires datapipeline. Live-sidorna uppdateras först; om en siffra här skiljer sig från marbellawire.com är webbplatsen aktuell och det här arkivet väntar på sin säsongsuppdatering. Rättelser: [öppna ett ärende](https://github.com/shiftdylson1/marbella-price-data/issues) med en käll-URL från stället självt.

## Licens och erkännande

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Fritt att använda, dela och bearbeta med erkännande av **Marbella Wire** och en länk till datasetets kanoniska sida (tabellen ovan). Licensen täcker bara siffrorna i det här arkivet; webbplatsens redaktionella text och dess krediterade fotografier ingår inte och är inte CC-licensierade.

## Citering

I löpande text: "Marbella Wire, https://marbellawire.com/prices/sunbeds/". Beständig identifierare för alla versioner: [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846). Se [CITATION.cff](../../CITATION.cff).

## Andra platser där datasetet finns

- GitHub: https://github.com/shiftdylson1/marbella-price-data
- Zenodo (DOI): https://doi.org/10.5281/zenodo.22094846
- Hugging Face: https://huggingface.co/datasets/editorwire11/marbella-price-data
- Kaggle: https://www.kaggle.com/datasets/marbellawire/marbella-price-data
- OpenML: https://www.openml.org/d/47281 · https://www.openml.org/d/47282 · https://www.openml.org/d/47283
