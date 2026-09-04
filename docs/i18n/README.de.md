# Marbella Preisdaten

Diese Seite auf: [Español](README.es.md) · [Français](README.fr.md) · [Nederlands](README.nl.md) · [Svenska](README.sv.md) · [English](../../README.md)

Offene Datensätze von [Marbella Wire](https://marbellawire.com), einem unabhängigen Preisbeobachter für Marbella, Spanien. Jede Zahl erscheint zuerst auf marbellawire.com; dieses Repository ist die maschinenlesbare Kopie, erzeugt aus denselben Datendateien, aus denen die Seiten gebaut werden. Diese Seite ist eine Übersetzung der [englischen README](../../README.md), die die Referenzfassung ist.

| Datensatz | Was es ist | Zeilen | Kanonische Seite |
|---|---|---|---|
| [Marbella Wire Liegestuhl-Index](../../sunbed-index/) | Hochsaisonpreis für zwei Liegestühle plus Mindestverzehr in den Beachclubs von Marbella, Saison 2026 | 8 mit Preis + 4 ohne veröffentlichten Preis | https://marbellawire.com/prices/sunbeds/ |
| [Marbella Chiringuito-Zensus](../../chiringuitos/) | Zensus der Chiringuitos (Strandrestaurants) von Marbella mit Preisen aus der Speisekarte, Saison 2026 | 43 Lokale | https://marbellawire.com/prices/chiringuitos/ |
| [Marbella Beachclub-Preiszensus](../../beach-clubs/) | Liegestuhl-, Bett- und Mindestverzehrpreise der Beachclubs, Saison 2026 | 14 Clubs | https://marbellawire.com/prices/beach-clubs/ |

Alle Preise in Euro; jede CSV-Zeile trägt ihre Einheit, ihre Quellseite und ihre Lizenz. Eine leere Zelle bedeutet, dass das Lokal keine Zahl veröffentlicht, nie null und nie eine Schätzung.

Schlagwörter: Marbella, Spanien, Preise, Tourismus, Beachclubs, Chiringuitos, Liegestühle, offene Daten, Costa del Sol, Preisindex, Datensatz

## In zehn Sekunden nutzen

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

Oder das Repository überspringen und den immer aktuellen Live-Endpunkt lesen:

```sh
curl https://marbellawire.com/data/sunbed-index.csv
```

## Spaltenglossar

Die Spaltennamen sind in allen Dateien englisch (sie sind Bezeichner); die Bedeutungen unten sind auf Deutsch.

### `sunbed-index/sunbed-index.csv`

**Marbella Wire Liegestuhl-Index.** Was Liegen an der Costa del Sol kostet: zwei Liegestühle plus Mindestverzehr in jedem großen Beachclub von Marbella, der einen Tarif veröffentlicht. Sommer 2026, 8 mit Preis + 4 ohne veröffentlichten Preis.

Schlagwörter: Liegestuhl Preise Marbella, Liegestuhl Preise Datensatz, Sonnenliege Preis Marbella, Liegestuhl-Index, Beachclub Preise Marbella, Puerto Banús Liegestuhl, Costa del Sol Preise, offene Daten, Marbella, Spanien, Preise, Tourismus, Beachclubs, Chiringuitos, Liegestühle, Costa del Sol, Preisindex, Datensatz

| Spalte | Typ | Bedeutung |
|---|---|---|
| `venue` | Text | Name des Lokals, wie im Index veröffentlicht |
| `area` | Text | Gegend von Marbella |
| `status` | Text | 'priced' (mit Preis) oder 'unpublished' (gelistet, aber bewusst ohne Preis; siehe note) |
| `price_eur` | Zahl | Zwei Liegestühle plus Mindestverzehr, EUR, Hochsaison, Wochentag. LEER bei unpublished-Zeilen, nie null |
| `price_unit` | Text | Wiederholt die Einheit der Zeile |
| `yoy_pct` | Zahl | Veränderung von price_eur zum Vorjahr, in Prozent |
| `note` | Text | Preisvorbehalt oder der Grund, warum eine Zeile keinen Preis hat |
| `season` | Zahl | Saisonjahr |
| `measured_on` | Datum | Datum, an dem der Preis abgelesen wurde (unpublished-Zeilen: Prüfdatum) |
| `price_source_url` | URL | Eigene Preisquelle des Lokals |
| `price_source_via` | Text | Name der Buchungsplattform, wenn die Quelle das Buchungssystem des Lokals ist |
| `index_average_eur` | Zahl | Durchschnitt des gesamten Index, in jeder Zeile wiederholt, damit eine Zeile ihre Schlagzahl mitführt |
| `index_average_unit` | Text | Einheit des Durchschnitts |
| `index_yoy_pct` | Zahl | Veränderung des gesamten Index zum Vorjahr, in Prozent |
| `source_url` | URL | Kanonische Indexseite |
| `attribution` | Text | Gewünschte Quellenangabe |
| `license` | URL | CC BY 4.0 |

### `chiringuitos/chiringuitos.csv`

**Marbella Chiringuito-Zensus.** Zensus der Chiringuitos (Strandrestaurants) von Marbella mit Preisen aus der Speisekarte: Espeto de sardinas und Signature-Gericht je Lokal, jede Zeile mit Link auf die eigene Karte des Lokals. Saison 2026, 43 Lokale.

Schlagwörter: Chiringuito Preise Marbella, Strandrestaurant Preise Marbella, Espeto Preis, Restaurant Preise Marbella, Chiringuitos Datensatz, Costa del Sol, offene Daten, Marbella, Spanien, Preise, Tourismus, Beachclubs, Chiringuitos, Liegestühle, Preisindex, Datensatz

| Spalte | Typ | Bedeutung |
|---|---|---|
| `slug` | Text | Stabile Kennung; zugleich Pfadsegment der Seite des Lokals |
| `name` | Text | Name des Lokals |
| `area` | Text | Gegend von Marbella |
| `beach` | Text | Strand, an dem das Lokal liegt, sofern erfasst |
| `business_status` | Text | 'operating' (in Betrieb) bei der letzten Prüfung |
| `opened_year` | Zahl | Eröffnungsjahr, sofern erfasst |
| `price_display` | Text | Richtpreis, wie auf der Seite angezeigt, EUR |
| `price_basis` | Text | Woraus der Preis abgelesen wird (z. B. 'menu', die Speisekarte) |
| `price_kind` | Text | Was der Preis kauft (z. B. 'espeto', 'mains' = Hauptgerichte) |
| `price_note` | Text | Vorbehalt zum Preis |
| `price_level` | Text | Relative Preisklasse, € bis €€€€ |
| `signature_dish` | Text | Signature-Gericht, sofern eines veröffentlicht ist |
| `signature_dish_eur` | Zahl | Dessen Preis laut Karte, EUR |
| `espeto_eur` | Zahl | Preis des Espeto de sardinas laut Karte, EUR, sofern veröffentlicht |
| `official_site` | URL | Eigene Website des Lokals |
| `menu_url` | URL | Vom Lokal selbst veröffentlichte Speisekarte (menu_is_pdf sagt, ob es ein PDF ist) |
| `menu_is_pdf` | Wahrheitswert | Ob menu_url ein PDF ist |
| `address` | Text | Straßenadresse |
| `phone` | Text | Veröffentlichte Geschäftstelefonnummer |
| `google_place_id` | Text | Google Place ID, zum Verknüpfen |
| `google_maps_url` | URL | Google-Maps-Eintrag |
| `verified` | Datum | Letzte Prüfung je Lokal |
| `url` | URL | Seite des Lokals auf marbellawire.com |
| `currency` | Text | Alle Preise in EUR |
| `source_page` | URL | Kanonische Seite des Datensatzes |
| `attribution` | Text | Gewünschte Quellenangabe |
| `license` | URL | CC BY 4.0 |

### `beach-clubs/beach-clubs.csv`

**Marbella Beachclub-Preiszensus.** Geprüfte Preise für Liegestuhl, Bett und Mindestverzehr in den Beachclubs von Marbella und Puerto Banús, abgelesen von der Preisliste oder Buchungsseite jedes Clubs. Saison 2026, 14 Clubs.

Schlagwörter: Beachclub Preise Marbella, Beach Club Puerto Banús Preise, Liegestuhl Preis, Mindestverzehr Beachclub, Marbella Beachclubs Datensatz, offene Daten, Marbella, Spanien, Preise, Tourismus, Beachclubs, Chiringuitos, Liegestühle, Costa del Sol, Preisindex, Datensatz

| Spalte | Typ | Bedeutung |
|---|---|---|
| `slug` | Text | Stabile Kennung; zugleich Pfadsegment der Seite des Clubs |
| `name` | Text | Name des Clubs |
| `area` | Text | Gegend von Marbella |
| `season_year` | Zahl | Saison, zu der die Preise gehören |
| `sunbed_price` | Zahl | Günstigster Liegestuhl bzw. günstigste Liege, EUR, Hochsaison, pro Tag |
| `bed_or_vip_price` | Zahl | Bestes Bett oder VIP-Paket, EUR, pro Tag, sofern veröffentlicht |
| `minimum_spend` | Text | Mindestverzehr-Bedingungen, wie veröffentlicht |
| `entry_policy` | Text | Wie der Eintritt tatsächlich funktioniert |
| `price_level` | Text | Relative Preisklasse |
| `source_url` | URL | Eigene Preisliste oder Buchungsseite des Clubs, von der die Preise abgelesen wurden |
| `rates_url` | URL | Veröffentlichtes Tarifdokument (rates_is_pdf sagt, ob PDF) |
| `rates_is_pdf` | Wahrheitswert | Ob rates_url ein PDF ist |
| `booking_url` | URL | Buchungssystem des Clubs |
| `official_site` | URL | Eigene Website des Clubs |
| `menu_url` | URL | Veröffentlichte Speisekarte, sofern vorhanden |
| `menu_is_pdf` | Wahrheitswert | Ob menu_url ein PDF ist |
| `opened_year` | Zahl | Eröffnungsjahr, sofern erfasst |
| `business_status` | Text | 'operating' (in Betrieb) bei der letzten Prüfung |
| `address` | Text | Straßenadresse |
| `phone` | Text | Veröffentlichte Geschäftstelefonnummer |
| `google_place_id` | Text | Google Place ID, zum Verknüpfen |
| `google_maps_url` | URL | Google-Maps-Eintrag |
| `verified` | Datum | Letzte Prüfung je Club |
| `notes` | Text | Hinweis zur Quelle |
| `url` | URL | Seite des Clubs auf marbellawire.com |
| `currency` | Text | Alle Preise in EUR |
| `source_page` | URL | Kanonische Seite des Datensatzes |
| `attribution` | Text | Gewünschte Quellenangabe |
| `license` | URL | CC BY 4.0 |

## Quellen und Methode

Die Preise stammen aus den von den Lokalen selbst veröffentlichten Speisekarten, Preislisten und Buchungssystemen, nie von Aggregatoren; jede Zeile verlinkt ihre Quelle. Eine Zahl in der Quelle ist eine Zahl hier: nichts wird geschätzt oder interpoliert. Details je Datensatz stehen in der PROVENANCE.md jedes Ordners (englisch).

## Aktualisierungen

Wird von Hand etwa einmal pro Saison aus der Datenpipeline von Marbella Wire neu erzeugt. Die Live-Seiten werden zuerst aktualisiert; weicht eine Zahl hier von marbellawire.com ab, ist die Website aktuell und dieses Repository wartet auf seine Saisonaktualisierung. Korrekturen: [Issue eröffnen](https://github.com/shiftdylson1/marbella-price-data/issues) mit einer Quell-URL des Lokals selbst.

## Lizenz und Namensnennung

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Frei nutzbar, teilbar und anpassbar mit Nennung von **Marbella Wire** und Link auf die kanonische Seite des Datensatzes (Tabelle oben). Die Lizenz deckt nur die Zahlen in diesem Repository; der redaktionelle Text der Website und ihre Fotos mit Bildnachweis sind nicht enthalten und nicht CC-lizenziert.

## Zitieren

Im Fließtext: "Marbella Wire, https://marbellawire.com/prices/sunbeds/". Dauerhafte Kennung für alle Versionen: [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846). Siehe [CITATION.cff](../../CITATION.cff).

## Weitere Fundorte dieses Datensatzes

- GitHub: https://github.com/shiftdylson1/marbella-price-data
- Zenodo (DOI): https://doi.org/10.5281/zenodo.22094846
- Hugging Face: https://huggingface.co/datasets/editorwire11/marbella-price-data
- Kaggle: https://www.kaggle.com/datasets/marbellawire/marbella-price-data
- OpenML: https://www.openml.org/d/47281 · https://www.openml.org/d/47282 · https://www.openml.org/d/47283
