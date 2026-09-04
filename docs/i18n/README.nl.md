# Marbella prijsdata

Lees dit in het: [Español](README.es.md) · [Deutsch](README.de.md) · [Français](README.fr.md) · [Svenska](README.sv.md) · [English](../../README.md)

Open datasets van [Marbella Wire](https://marbellawire.com), een onafhankelijke prijsmonitor voor Marbella, Spanje. Elk cijfer verschijnt eerst op marbellawire.com; deze repository is de machineleesbare kopie, opnieuw gegenereerd uit dezelfde databestanden waaruit de pagina's worden gebouwd. Deze pagina is een vertaling van de [Engelse README](../../README.md), die de referentieversie is.

| Dataset | Wat het is | Rijen | Canonieke pagina |
|---|---|---|---|
| [Marbella Wire Ligbedindex](../../sunbed-index/) | Hoogseizoenprijs van twee ligbedden plus minimale besteding bij de beachclubs van Marbella, seizoen 2026 | 8 met prijs + 4 zonder gepubliceerde prijs | https://marbellawire.com/prices/sunbeds/ |
| [Marbella chiringuito-census](../../chiringuitos/) | Census van de chiringuitos (strandrestaurants) van Marbella met prijzen uit de menukaart, seizoen 2026 | 43 zaken | https://marbellawire.com/prices/chiringuitos/ |
| [Marbella beachclub-prijscensus](../../beach-clubs/) | Ligbed-, bed- en minimale-bestedingsprijzen van de beachclubs, seizoen 2026 | 14 clubs | https://marbellawire.com/prices/beach-clubs/ |

Alle prijzen zijn in euro; elke CSV-rij draagt zijn eenheid, bronpagina en licentie. Een lege cel betekent dat de zaak geen cijfer publiceert, nooit nul en nooit een schatting.

Trefwoorden: Marbella, Spanje, prijzen, toerisme, beachclubs, chiringuitos, ligbedden, open data, Costa del Sol, prijsindex, dataset

## In tien seconden aan de slag

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

Of sla de repository over en lees het altijd actuele live-eindpunt:

```sh
curl https://marbellawire.com/data/sunbed-index.csv
```

## Kolomwoordenlijst

De kolomnamen zijn in elk bestand Engels (het zijn identifiers); de betekenissen hieronder zijn in het Nederlands.

### `sunbed-index/sunbed-index.csv`

**Marbella Wire Ligbedindex.** Wat een dag liggen aan de Costa del Sol kost: twee ligbedden plus minimale besteding bij elke grote beachclub van Marbella die een tarief publiceert. Zomer 2026, 8 met prijs + 4 zonder gepubliceerde prijs.

Trefwoorden: ligbed prijzen Marbella, ligbedden prijs Marbella, zonnebed prijs Marbella, ligbedindex, beachclub prijzen Marbella, ligbed Puerto Banús, Costa del Sol prijzen, open data, Marbella, Spanje, prijzen, toerisme, beachclubs, chiringuitos, ligbedden, Costa del Sol, prijsindex, dataset

| kolom | type | betekenis |
|---|---|---|
| `venue` | tekst | Naam van de zaak zoals gepubliceerd in de index |
| `area` | tekst | Deel van Marbella |
| `status` | tekst | 'priced' (met prijs) of 'unpublished' (vermeld maar bewust zonder prijs; zie note) |
| `price_eur` | getal | Twee ligbedden plus minimale besteding, EUR, hoogseizoen, doordeweeks. LEEG bij unpublished-rijen, nooit nul |
| `price_unit` | tekst | Herhaalt de eenheid van de rij |
| `yoy_pct` | getal | Verandering van price_eur ten opzichte van vorig jaar, in procent |
| `note` | tekst | Voorbehoud bij de prijs, of de reden waarom een rij geen prijs heeft |
| `season` | getal | Seizoensjaar |
| `measured_on` | datum | Datum waarop de prijs is afgelezen (unpublished-rijen: controledatum) |
| `price_source_url` | URL | Eigen prijsbron van de zaak |
| `price_source_via` | tekst | Naam van het boekingsplatform wanneer de bron het boekingssysteem van de zaak is |
| `index_average_eur` | getal | Gemiddelde van de hele index, op elke rij herhaald zodat een rij zijn kerncijfer meeneemt |
| `index_average_unit` | tekst | Eenheid van het gemiddelde |
| `index_yoy_pct` | getal | Verandering van de hele index ten opzichte van vorig jaar, in procent |
| `source_url` | URL | Canonieke indexpagina |
| `attribution` | tekst | Gevraagde bronvermelding |
| `license` | URL | CC BY 4.0 |

### `chiringuitos/chiringuitos.csv`

**Marbella chiringuito-census.** Census van de chiringuitos (strandrestaurants) van Marbella met prijzen uit de menukaart: espeto de sardinas en signatuurgerecht per zaak, elke rij gekoppeld aan de eigen menukaart van de zaak. Seizoen 2026, 43 zaken.

Trefwoorden: chiringuito prijzen Marbella, strandrestaurant prijzen Marbella, espeto prijs, restaurantprijzen Marbella, chiringuitos dataset, Costa del Sol, open data, Marbella, Spanje, prijzen, toerisme, beachclubs, chiringuitos, ligbedden, prijsindex, dataset

| kolom | type | betekenis |
|---|---|---|
| `slug` | tekst | Stabiele id; tevens het padsegment van de pagina van de zaak |
| `name` | tekst | Naam van de zaak |
| `area` | tekst | Deel van Marbella |
| `beach` | tekst | Strand waar de zaak ligt, indien vastgelegd |
| `business_status` | tekst | 'operating' (in bedrijf) bij de laatste controle |
| `opened_year` | getal | Openingsjaar, indien vastgelegd |
| `price_display` | tekst | Richtprijs zoals getoond op de pagina, EUR |
| `price_basis` | tekst | Waaruit de prijs is afgelezen (bijv. 'menu', de menukaart) |
| `price_kind` | tekst | Wat de prijs koopt (bijv. 'espeto', 'mains' = hoofdgerechten) |
| `price_note` | tekst | Voorbehoud bij de prijs |
| `price_level` | tekst | Relatieve prijsklasse, € tot €€€€ |
| `signature_dish` | tekst | Signatuurgerecht, indien er een is gepubliceerd |
| `signature_dish_eur` | getal | De menuprijs daarvan, EUR |
| `espeto_eur` | getal | Menuprijs van de espeto de sardinas, EUR, indien gepubliceerd |
| `official_site` | URL | Eigen website van de zaak |
| `menu_url` | URL | Door de zaak zelf gepubliceerde menukaart (menu_is_pdf zegt of het een PDF is) |
| `menu_is_pdf` | booleaans | Of menu_url een PDF is |
| `address` | tekst | Straatadres |
| `phone` | tekst | Gepubliceerd zakelijk telefoonnummer |
| `google_place_id` | tekst | Google Place ID, om te koppelen |
| `google_maps_url` | URL | Google Maps-vermelding |
| `verified` | datum | Laatste controle per zaak |
| `url` | URL | Pagina van de zaak op marbellawire.com |
| `currency` | tekst | Alle prijzen zijn in EUR |
| `source_page` | URL | Canonieke pagina van de dataset |
| `attribution` | tekst | Gevraagde bronvermelding |
| `license` | URL | CC BY 4.0 |

### `beach-clubs/beach-clubs.csv`

**Marbella beachclub-prijscensus.** Geverifieerde prijzen voor ligbed, bed en minimale besteding bij de beachclubs van Marbella en Puerto Banús, afgelezen van de tarievenkaart of boekingspagina van elke club. Seizoen 2026, 14 clubs.

Trefwoorden: beachclub prijzen Marbella, beach club Puerto Banús prijzen, ligbed prijs, minimale besteding beachclub, Marbella beachclubs dataset, open data, Marbella, Spanje, prijzen, toerisme, beachclubs, chiringuitos, ligbedden, Costa del Sol, prijsindex, dataset

| kolom | type | betekenis |
|---|---|---|
| `slug` | tekst | Stabiele id; tevens het padsegment van de pagina van de club |
| `name` | tekst | Naam van de club |
| `area` | tekst | Deel van Marbella |
| `season_year` | getal | Seizoen waartoe de prijzen behoren |
| `sunbed_price` | getal | Goedkoopste ligbed of ligstoel, EUR, hoogseizoen, per dag |
| `bed_or_vip_price` | getal | Beste bed of VIP-pakket, EUR, per dag, indien gepubliceerd |
| `minimum_spend` | tekst | Voorwaarden voor minimale besteding, zoals gepubliceerd |
| `entry_policy` | tekst | Hoe de toegang in de praktijk werkt |
| `price_level` | tekst | Relatieve prijsklasse |
| `source_url` | URL | Eigen tarievenkaart of boekingspagina van de club waarvan de prijzen zijn afgelezen |
| `rates_url` | URL | Gepubliceerd tarievendocument (rates_is_pdf zegt of het een PDF is) |
| `rates_is_pdf` | booleaans | Of rates_url een PDF is |
| `booking_url` | URL | Boekingssysteem van de club |
| `official_site` | URL | Eigen website van de club |
| `menu_url` | URL | Gepubliceerde menukaart, indien aanwezig |
| `menu_is_pdf` | booleaans | Of menu_url een PDF is |
| `opened_year` | getal | Openingsjaar, indien vastgelegd |
| `business_status` | tekst | 'operating' (in bedrijf) bij de laatste controle |
| `address` | tekst | Straatadres |
| `phone` | tekst | Gepubliceerd zakelijk telefoonnummer |
| `google_place_id` | tekst | Google Place ID, om te koppelen |
| `google_maps_url` | URL | Google Maps-vermelding |
| `verified` | datum | Laatste controle per club |
| `notes` | tekst | Opmerking over de bron |
| `url` | URL | Pagina van de club op marbellawire.com |
| `currency` | tekst | Alle prijzen zijn in EUR |
| `source_page` | URL | Canonieke pagina van de dataset |
| `attribution` | tekst | Gevraagde bronvermelding |
| `license` | URL | CC BY 4.0 |

## Bronnen en methode

De prijzen komen uit de door de zaken zelf gepubliceerde menukaarten, tarievenkaarten en boekingssystemen, nooit van aggregators; elke rij linkt naar zijn bron. Eén getal in de bron is één getal hier: niets wordt geschat of geïnterpoleerd. Details per dataset staan in de PROVENANCE.md van elke map (Engels).

## Updates

Handmatig opnieuw gegenereerd, ongeveer eens per seizoen, vanuit de datapipeline van Marbella Wire. De live-pagina's worden eerst bijgewerkt; wijkt een cijfer hier af van marbellawire.com, dan is de site actueel en wacht deze repository op zijn seizoensverversing. Correcties: [open een issue](https://github.com/shiftdylson1/marbella-price-data/issues) met een bron-URL van de zaak zelf.

## Licentie en naamsvermelding

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Vrij te gebruiken, delen en bewerken met vermelding van **Marbella Wire** en een link naar de canonieke pagina van de dataset (tabel hierboven). De licentie dekt alleen de cijfers in deze repository; de redactionele tekst van de site en de gecrediteerde foto's zijn niet inbegrepen en vallen niet onder CC.

## Citeren

In lopende tekst: "Marbella Wire, https://marbellawire.com/prices/sunbeds/". Permanente identifier voor alle versies: [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846). Zie [CITATION.cff](../../CITATION.cff).

## Andere vindplaatsen van deze dataset

- GitHub: https://github.com/shiftdylson1/marbella-price-data
- Zenodo (DOI): https://doi.org/10.5281/zenodo.22094846
- Hugging Face: https://huggingface.co/datasets/editorwire11/marbella-price-data
- Kaggle: https://www.kaggle.com/datasets/marbellawire/marbella-price-data
- OpenML: https://www.openml.org/d/47281 · https://www.openml.org/d/47282 · https://www.openml.org/d/47283
