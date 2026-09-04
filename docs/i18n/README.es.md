# Datos de precios de Marbella

Leer en: [Deutsch](README.de.md) · [Français](README.fr.md) · [Nederlands](README.nl.md) · [Svenska](README.sv.md) · [English](../../README.md)

Conjuntos de datos abiertos de [Marbella Wire](https://marbellawire.com), un observatorio de precios independiente de Marbella, España. Cada cifra se publica primero en marbellawire.com; este repositorio es la copia legible por máquina, regenerada a partir de los mismos archivos de datos que generan las páginas. Esta página es una traducción del [README en inglés](../../README.md), que es la versión de referencia.

| Conjunto de datos | Qué es | Filas | Página canónica |
|---|---|---|---|
| [Índice de hamacas de Marbella Wire](../../sunbed-index/) | Precio en temporada alta de dos hamacas más consumo mínimo en los beach clubs de Marbella, temporada 2026 | 8 con precio + 4 sin precio publicado | https://marbellawire.com/prices/sunbeds/ |
| [Censo de chiringuitos de Marbella](../../chiringuitos/) | Censo de chiringuitos de Marbella con precios tomados de la carta, temporada 2026 | 43 locales | https://marbellawire.com/prices/chiringuitos/ |
| [Censo de precios de beach clubs de Marbella](../../beach-clubs/) | Precios de hamaca, cama balinesa y consumo mínimo en los beach clubs, temporada 2026 | 14 clubs | https://marbellawire.com/prices/beach-clubs/ |

Todos los precios están en euros; cada fila del CSV lleva su unidad, su página de origen y su licencia. Una celda vacía significa que el local no publica la cifra, nunca cero y nunca una estimación.

Palabras clave: Marbella, España, precios, turismo, beach clubs, chiringuitos, hamacas, datos abiertos, Costa del Sol, índice de precios, conjunto de datos

## Úsalo en diez segundos

```python
import pandas as pd
df = pd.read_csv('https://raw.githubusercontent.com/shiftdylson1/marbella-price-data/main/beach-clubs/beach-clubs.csv')
```

O sáltate el repositorio y lee el endpoint en vivo, siempre actualizado:

```sh
curl https://marbellawire.com/data/sunbed-index.csv
```

## Glosario de columnas

Los nombres de columna están en inglés en todos los archivos (son identificadores); los significados de abajo están en Español.

### `sunbed-index/sunbed-index.csv`

**Índice de hamacas de Marbella Wire.** Lo que cuesta tumbarse en la Costa del Sol: dos hamacas más consumo mínimo en cada gran beach club de Marbella que publica su tarifa. Verano 2026, 8 con precio + 4 sin precio publicado.

Palabras clave: precio hamaca Marbella, precios hamacas Marbella, índice de hamacas, precio tumbona playa, precios beach club Marbella, hamacas Puerto Banús, precios Costa del Sol, datos abiertos, Marbella, España, precios, turismo, beach clubs, chiringuitos, hamacas, Costa del Sol, índice de precios, conjunto de datos

| columna | tipo | significado |
|---|---|---|
| `venue` | texto | Nombre del local tal y como aparece en el índice |
| `area` | texto | Zona de Marbella |
| `status` | texto | 'priced' (con precio) o 'unpublished' (listado pero sin precio a propósito; ver note) |
| `price_eur` | número | Dos hamacas más consumo mínimo, EUR, temporada alta, día laborable. VACÍO en las filas unpublished, nunca cero |
| `price_unit` | texto | Repite la unidad de la fila |
| `yoy_pct` | número | Variación interanual de price_eur, en porcentaje |
| `note` | texto | Matiz sobre el precio, o el motivo por el que la fila no tiene precio |
| `season` | número | Año de la temporada |
| `measured_on` | fecha | Fecha en que se leyó el precio (filas unpublished: fecha de comprobación) |
| `price_source_url` | URL | Fuente de precio del propio local |
| `price_source_via` | texto | Nombre de la plataforma de reservas cuando la fuente es el sistema de reservas del local |
| `index_average_eur` | número | Media de todo el índice, repetida en cada fila para que la fila viaje con su titular |
| `index_average_unit` | texto | Unidad de la media |
| `index_yoy_pct` | número | Variación interanual de todo el índice, en porcentaje |
| `source_url` | URL | Página canónica del índice |
| `attribution` | texto | Línea de crédito solicitada |
| `license` | URL | CC BY 4.0 |

### `chiringuitos/chiringuitos.csv`

**Censo de chiringuitos de Marbella.** Censo de chiringuitos de Marbella con precios tomados de la carta: espeto de sardinas y plato estrella por local, cada fila enlazada a la carta del propio chiringuito. Temporada 2026, 43 locales.

Palabras clave: precios chiringuitos Marbella, precio espeto Marbella, chiringuitos Marbella datos, precios restaurantes playa Marbella, censo chiringuitos, Costa del Sol, datos abiertos, Marbella, España, precios, turismo, beach clubs, chiringuitos, hamacas, índice de precios, conjunto de datos

| columna | tipo | significado |
|---|---|---|
| `slug` | texto | Identificador estable; también el segmento de ruta de la página del local |
| `name` | texto | Nombre del local |
| `area` | texto | Zona de Marbella |
| `beach` | texto | Playa en la que está el local, cuando consta |
| `business_status` | texto | 'operating' (en funcionamiento) en la última verificación |
| `opened_year` | número | Año de apertura, cuando consta |
| `price_display` | texto | Precio de referencia tal y como se muestra en la página, EUR |
| `price_basis` | texto | De dónde se lee el precio (p. ej. 'menu', la carta) |
| `price_kind` | texto | Qué compra el precio (p. ej. 'espeto', 'mains' = platos principales) |
| `price_note` | texto | Matiz sobre el precio |
| `price_level` | texto | Franja relativa, de € a €€€€ |
| `signature_dish` | texto | Plato estrella, cuando se publica uno |
| `signature_dish_eur` | número | Su precio en carta, EUR |
| `espeto_eur` | número | Precio en carta del espeto de sardinas, EUR, cuando se publica |
| `official_site` | URL | Web del propio local |
| `menu_url` | URL | Carta publicada por el propio local (menu_is_pdf indica si es un PDF) |
| `menu_is_pdf` | booleano | Si menu_url es un PDF |
| `address` | texto | Dirección postal |
| `phone` | texto | Teléfono publicado del negocio |
| `google_place_id` | texto | Google Place ID, para cruzar datos |
| `google_maps_url` | URL | Ficha en Google Maps |
| `verified` | fecha | Última verificación por local |
| `url` | URL | Página del local en marbellawire.com |
| `currency` | texto | Todos los precios están en EUR |
| `source_page` | URL | Página canónica del conjunto de datos |
| `attribution` | texto | Línea de crédito solicitada |
| `license` | URL | CC BY 4.0 |

### `beach-clubs/beach-clubs.csv`

**Censo de precios de beach clubs de Marbella.** Precios verificados de hamaca, cama balinesa y consumo mínimo en los beach clubs de Marbella y Puerto Banús, leídos de la tarifa o la página de reservas de cada club. Temporada 2026, 14 clubs.

Palabras clave: precios beach club Marbella, beach clubs Puerto Banús precios, precio hamaca, consumo mínimo beach club, beach clubs Marbella datos, datos abiertos, Marbella, España, precios, turismo, beach clubs, chiringuitos, hamacas, Costa del Sol, índice de precios, conjunto de datos

| columna | tipo | significado |
|---|---|---|
| `slug` | texto | Identificador estable; también el segmento de ruta de la página del club |
| `name` | texto | Nombre del club |
| `area` | texto | Zona de Marbella |
| `season_year` | número | Temporada a la que corresponden los precios |
| `sunbed_price` | número | Hamaca o tumbona más barata, EUR, temporada alta, por día |
| `bed_or_vip_price` | número | Cama balinesa superior o paquete VIP, EUR, por día, cuando se publica |
| `minimum_spend` | texto | Condiciones de consumo mínimo, tal y como se publican |
| `entry_policy` | texto | Cómo funciona realmente la entrada |
| `price_level` | texto | Franja relativa |
| `source_url` | URL | Tarifa o página de reservas del propio club de la que se leyeron los precios |
| `rates_url` | URL | Documento de tarifas publicado (rates_is_pdf indica si es PDF) |
| `rates_is_pdf` | booleano | Si rates_url es un PDF |
| `booking_url` | URL | Sistema de reservas del club |
| `official_site` | URL | Web del propio club |
| `menu_url` | URL | Carta publicada, cuando existe |
| `menu_is_pdf` | booleano | Si menu_url es un PDF |
| `opened_year` | número | Año de apertura, cuando consta |
| `business_status` | texto | 'operating' (en funcionamiento) en la última verificación |
| `address` | texto | Dirección postal |
| `phone` | texto | Teléfono publicado del negocio |
| `google_place_id` | texto | Google Place ID, para cruzar datos |
| `google_maps_url` | URL | Ficha en Google Maps |
| `verified` | fecha | Última verificación por club |
| `notes` | texto | Nota sobre la fuente |
| `url` | URL | Página del club en marbellawire.com |
| `currency` | texto | Todos los precios están en EUR |
| `source_page` | URL | Página canónica del conjunto de datos |
| `attribution` | texto | Línea de crédito solicitada |
| `license` | URL | CC BY 4.0 |

## Fuentes y método

Los precios proceden de las cartas, tarifas y sistemas de reservas publicados por los propios locales, nunca de agregadores; cada fila enlaza su fuente. Un número en la fuente es un número aquí: nada se estima ni se interpola. El detalle por conjunto de datos está en el PROVENANCE.md de cada carpeta (en inglés).

## Actualizaciones

Se regenera a mano, más o menos una vez por temporada, desde el pipeline de datos de Marbella Wire. Las páginas en vivo se actualizan primero; si una cifra de aquí no coincide con marbellawire.com, la web es la vigente y este repositorio espera su actualización de temporada. Correcciones: [abre un issue](https://github.com/shiftdylson1/marbella-price-data/issues) con una URL de fuente del propio local.

## Licencia y atribución

[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). Uso, distribución y adaptación libres citando a **Marbella Wire** y enlazando la página canónica del conjunto de datos (tabla de arriba). La licencia cubre solo las cifras de este repositorio; el texto editorial de la web y sus fotografías acreditadas no están incluidos ni tienen licencia CC.

## Cómo citar

En texto: "Marbella Wire, https://marbellawire.com/prices/sunbeds/". Identificador permanente para todas las versiones: [10.5281/zenodo.22094846](https://doi.org/10.5281/zenodo.22094846). Ver [CITATION.cff](../../CITATION.cff).

## Otras ubicaciones de este conjunto de datos

- GitHub: https://github.com/shiftdylson1/marbella-price-data
- Zenodo (DOI): https://doi.org/10.5281/zenodo.22094846
- Hugging Face: https://huggingface.co/datasets/editorwire11/marbella-price-data
- Kaggle: https://www.kaggle.com/datasets/marbellawire/marbella-price-data
- OpenML: https://www.openml.org/d/47281 · https://www.openml.org/d/47282 · https://www.openml.org/d/47283
