# Données de prix de Marbella

Lire en: [Español](README.es.md) · [Deutsch](README.de.md) · [Nederlands](README.nl.md) · [Svenska](README.sv.md) · [English](../../README.md)

Jeux de données ouverts de [Marbella Wire](https://marbellawire.com), un observatoire des prix indépendant pour Marbella, Espagne. Chaque chiffre est d'abord publié sur marbellawire.com ; ce dépôt en est la copie lisible par machine, régénérée à partir des mêmes fichiers de données qui produisent les pages. Cette page est une traduction du [README en anglais](../../README.md), qui fait référence.

| Jeu de données | Ce que c'est | Lignes | Page canonique |
|---|---|---|---|
| [Indice des transats Marbella Wire](../../sunbed-index/) | Prix en haute saison de deux transats plus consommation minimum dans les beach clubs de Marbella, saison 2026 | 8 avec prix + 4 sans prix publié | https://marbellawire.com/prices/sunbeds/ |
| [Recensement des chiringuitos de Marbella](../../chiringuitos/) | Recensement des chiringuitos (restaurants de plage) de Marbella avec prix relevés sur la carte, saison 2026 | 43 établissements | https://marbellawire.com/prices/chiringuitos/ |
| [Recensement des prix des beach clubs de Marbella](../../beach-clubs/) | Prix des transats, lits et consommation minimum dans les beach clubs, saison 2026 | 14 clubs | https://marbellawire.com/prices/beach-clubs/ |

Tous les prix sont en euros ; chaque ligne CSV porte son unité, sa page source et sa licence. Une cellule vide signifie que l'établissement ne publie pas le chiffre, jamais zéro et jamais une estimation.

Mots-clés: Marbella, Espagne, prix, tourisme, beach clubs, chiringuitos, transats, données ouvertes, Costa del Sol, indice des prix, jeu de données

## À utiliser en dix secondes

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

Ou sauter le dépôt et lire le point d'accès en direct, toujours à jour :

```sh
curl https://marbellawire.com/data/sunbed-index.csv
```

## Glossaire des colonnes

Les noms de colonnes sont en anglais dans tous les fichiers (ce sont des identifiants) ; les significations ci-dessous sont en Français.

### `sunbed-index/sunbed-index.csv`

**Indice des transats Marbella Wire.** Ce que coûte une journée allongé sur la Costa del Sol : deux transats plus consommation minimum dans chaque grand beach club de Marbella qui publie un tarif. Été 2026, 8 avec prix + 4 sans prix publié.

Mots-clés: prix transat Marbella, prix transats Marbella, indice des transats, prix chaise longue plage, prix beach club Marbella, transat Puerto Banús, prix Costa del Sol, données ouvertes, Marbella, Espagne, prix, tourisme, beach clubs, chiringuitos, transats, Costa del Sol, indice des prix, jeu de données

| colonne | type | signification |
|---|---|---|
| `venue` | texte | Nom de l'établissement tel que publié dans l'indice |
| `area` | texte | Quartier de Marbella |
| `status` | texte | 'priced' (avec prix) ou 'unpublished' (listé mais volontairement sans prix ; voir note) |
| `price_eur` | nombre | Deux transats plus consommation minimum, EUR, haute saison, jour de semaine. VIDE pour les lignes unpublished, jamais zéro |
| `price_unit` | texte | Rappelle l'unité de la ligne |
| `yoy_pct` | nombre | Variation de price_eur sur un an, en pourcentage |
| `note` | texte | Réserve sur le prix, ou raison pour laquelle la ligne est sans prix |
| `season` | nombre | Année de la saison |
| `measured_on` | date | Date de relevé du prix (lignes unpublished : date de vérification) |
| `price_source_url` | URL | Source de prix propre à l'établissement |
| `price_source_via` | texte | Nom de la plateforme de réservation quand la source est le système de réservation de l'établissement |
| `index_average_eur` | nombre | Moyenne de tout l'indice, répétée sur chaque ligne pour qu'une ligne voyage avec son chiffre clé |
| `index_average_unit` | texte | Unité de la moyenne |
| `index_yoy_pct` | nombre | Variation de tout l'indice sur un an, en pourcentage |
| `source_url` | URL | Page canonique de l'indice |
| `attribution` | texte | Mention de crédit demandée |
| `license` | URL | CC BY 4.0 |

### `chiringuitos/chiringuitos.csv`

**Recensement des chiringuitos de Marbella.** Recensement des chiringuitos (restaurants de plage) de Marbella avec prix relevés sur la carte : espeto de sardinas et plat signature par établissement, chaque ligne reliée à la carte de l'établissement. Saison 2026, 43 établissements.

Mots-clés: prix chiringuitos Marbella, prix restaurant de plage Marbella, prix espeto, prix restaurants Marbella, chiringuitos jeu de données, Costa del Sol, données ouvertes, Marbella, Espagne, prix, tourisme, beach clubs, chiringuitos, transats, indice des prix, jeu de données

| colonne | type | signification |
|---|---|---|
| `slug` | texte | Identifiant stable ; aussi le segment de chemin de la page de l'établissement |
| `name` | texte | Nom de l'établissement |
| `area` | texte | Quartier de Marbella |
| `beach` | texte | Plage où se trouve l'établissement, lorsqu'elle est renseignée |
| `business_status` | texte | 'operating' (en activité) à la dernière vérification |
| `opened_year` | nombre | Année d'ouverture, lorsqu'elle est renseignée |
| `price_display` | texte | Prix de référence tel qu'affiché sur la page, EUR |
| `price_basis` | texte | Source de lecture du prix (par ex. 'menu', la carte) |
| `price_kind` | texte | Ce que le prix achète (par ex. 'espeto', 'mains' = plats principaux) |
| `price_note` | texte | Réserve sur le prix |
| `price_level` | texte | Tranche relative, de € à €€€€ |
| `signature_dish` | texte | Plat signature, lorsqu'il est publié |
| `signature_dish_eur` | nombre | Son prix à la carte, EUR |
| `espeto_eur` | nombre | Prix à la carte de l'espeto de sardinas, EUR, lorsqu'il est publié |
| `official_site` | URL | Site web de l'établissement |
| `menu_url` | URL | Carte publiée par l'établissement lui-même (menu_is_pdf indique s'il s'agit d'un PDF) |
| `menu_is_pdf` | booléen | Si menu_url est un PDF |
| `address` | texte | Adresse postale |
| `phone` | texte | Téléphone professionnel publié |
| `google_place_id` | texte | Google Place ID, pour les jointures |
| `google_maps_url` | URL | Fiche Google Maps |
| `verified` | date | Dernière vérification par établissement |
| `url` | URL | Page de l'établissement sur marbellawire.com |
| `currency` | texte | Tous les prix sont en EUR |
| `source_page` | URL | Page canonique du jeu de données |
| `attribution` | texte | Mention de crédit demandée |
| `license` | URL | CC BY 4.0 |

### `beach-clubs/beach-clubs.csv`

**Recensement des prix des beach clubs de Marbella.** Prix vérifiés des transats, lits et consommation minimum dans les beach clubs de Marbella et Puerto Banús, relevés sur la grille tarifaire ou la page de réservation de chaque club. Saison 2026, 14 clubs.

Mots-clés: prix beach club Marbella, prix beach clubs Puerto Banús, prix transat, consommation minimum beach club, beach clubs Marbella jeu de données, données ouvertes, Marbella, Espagne, prix, tourisme, beach clubs, chiringuitos, transats, Costa del Sol, indice des prix, jeu de données

| colonne | type | signification |
|---|---|---|
| `slug` | texte | Identifiant stable ; aussi le segment de chemin de la page du club |
| `name` | texte | Nom du club |
| `area` | texte | Quartier de Marbella |
| `season_year` | nombre | Saison à laquelle les prix se rapportent |
| `sunbed_price` | nombre | Transat ou chaise longue le moins cher, EUR, haute saison, par jour |
| `bed_or_vip_price` | nombre | Lit haut de gamme ou forfait VIP, EUR, par jour, lorsqu'il est publié |
| `minimum_spend` | texte | Conditions de consommation minimum, telles que publiées |
| `entry_policy` | texte | Comment l'entrée fonctionne réellement |
| `price_level` | texte | Tranche relative |
| `source_url` | URL | Grille tarifaire ou page de réservation du club où les prix ont été relevés |
| `rates_url` | URL | Document tarifaire publié (rates_is_pdf indique si PDF) |
| `rates_is_pdf` | booléen | Si rates_url est un PDF |
| `booking_url` | URL | Système de réservation du club |
| `official_site` | URL | Site web du club |
| `menu_url` | URL | Carte publiée, lorsqu'il en existe une |
| `menu_is_pdf` | booléen | Si menu_url est un PDF |
| `opened_year` | nombre | Année d'ouverture, lorsqu'elle est renseignée |
| `business_status` | texte | 'operating' (en activité) à la dernière vérification |
| `address` | texte | Adresse postale |
| `phone` | texte | Téléphone professionnel publié |
| `google_place_id` | texte | Google Place ID, pour les jointures |
| `google_maps_url` | URL | Fiche Google Maps |
| `verified` | date | Dernière vérification par club |
| `notes` | texte | Note sur la source |
| `url` | URL | Page du club sur marbellawire.com |
| `currency` | texte | Tous les prix sont en EUR |
| `source_page` | URL | Page canonique du jeu de données |
| `attribution` | texte | Mention de crédit demandée |
| `license` | URL | CC BY 4.0 |

## Sources et méthode

Les prix proviennent des cartes, grilles tarifaires et systèmes de réservation publiés par les établissements eux-mêmes, jamais d'agrégateurs ; chaque ligne relie sa source. Un chiffre dans la source est un chiffre ici : rien n'est estimé ni interpolé. Le détail par jeu de données est dans le PROVENANCE.md de chaque dossier (en anglais).

## Mises à jour

Régénéré à la main, environ une fois par saison, depuis le pipeline de données de Marbella Wire. Les pages en direct sont mises à jour en premier ; si un chiffre ici diffère de marbellawire.com, le site est à jour et ce dépôt attend son rafraîchissement saisonnier. Corrections : [ouvrir un ticket](https://github.com/shiftdylson1/marbella-price-data/issues) avec une URL source de l'établissement lui-même.

## Licence et attribution

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Libre d'utilisation, de partage et d'adaptation en créditant **Marbella Wire** et en reliant la page canonique du jeu de données (tableau ci-dessus). La licence ne couvre que les chiffres de ce dépôt ; le texte éditorial du site et ses photographies créditées ne sont pas inclus et ne sont pas sous licence CC.

## Citation

En texte : "Marbella Wire, https://marbellawire.com/prices/sunbeds/". Identifiant permanent pour toutes les versions : [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846). Voir [CITATION.cff](../../CITATION.cff).

## Autres emplacements de ce jeu de données

- GitHub: https://github.com/shiftdylson1/marbella-price-data
- Zenodo (DOI): https://doi.org/10.5281/zenodo.22094846
- Hugging Face: https://huggingface.co/datasets/editorwire11/marbella-price-data
- Kaggle: https://www.kaggle.com/datasets/marbellawire/marbella-price-data
- OpenML: https://www.openml.org/d/47281 · https://www.openml.org/d/47282 · https://www.openml.org/d/47283
