// Marbella Wire open price datasets, embedded as CSV/JSON. Zero dependencies.
//
//   import { load } from 'marbella-prices';
//   const rows = await load('sunbed_index');   // [{ venue, area, status, price_eur, ... }]
//
// Figures are CC BY 4.0 (credit "Marbella Wire" + the canonical page); this
// loader is MIT. Live, always-current copies: https://marbellawire.com/data/

import { readFile } from 'node:fs/promises';
import { readFileSync } from 'node:fs';

export const version = '2026.9.4';
export const SITE = 'https://marbellawire.com';
export const REPOSITORY = 'https://github.com/shiftdylson1/marbella-price-data';
export const CONCEPT_DOI = '10.5281/zenodo.22094846';
export const DATA_LICENSE = 'CC-BY-4.0';
export const DATA_LICENSE_URL = 'https://creativecommons.org/licenses/by/4.0/';
export const LOCALES = ['en', 'es', 'de', 'fr', 'nl', 'sv'];

export const DATASETS = {
  sunbed_index: { file: 'sunbed-index', path: '/prices/sunbeds/', title: 'Marbella Wire Sunbed Index' },
  chiringuitos: { file: 'chiringuitos', path: '/prices/chiringuitos/', title: 'Marbella Wire chiringuito census' },
  beach_clubs: { file: 'beach-clubs', path: '/prices/beach-clubs/', title: 'Marbella Wire beach club price census' },
};

function key(name) {
  const k = String(name).trim().toLowerCase().replace(/-/g, '_');
  if (!(k in DATASETS)) throw new Error(`unknown dataset "${name}"; choose from ${Object.keys(DATASETS).join(', ')}`);
  return k;
}

const dataUrl = (stem, ext) => new URL(`./data/${stem}.${ext}`, import.meta.url);

// CSV cells are text; numbers, booleans and blanks get their obvious type.
// Empty = not published (never zero, never a guess) -> null. Text that merely
// starts with digits (phones, "€8.90") stays text.
function coerce(v) {
  if (v === '') return null;
  if (v === 'true') return true;
  if (v === 'false') return false;
  return /^-?\d+(\.\d+)?$/.test(v) ? Number(v) : v;
}

/** RFC 4180 parser: quoted fields, doubled quotes, CRLF or LF, newlines inside quotes. */
export function parseCsv(text) {
  const rows = [];
  let row = [];
  let field = '';
  let quoted = false;
  for (let i = 0; i < text.length; i++) {
    const c = text[i];
    if (quoted) {
      if (c === '"') {
        if (text[i + 1] === '"') { field += '"'; i++; } else quoted = false;
      } else field += c;
    } else if (c === '"') quoted = true;
    else if (c === ',') { row.push(field); field = ''; }
    else if (c === '\n' || c === '\r') {
      if (c === '\r' && text[i + 1] === '\n') i++;
      row.push(field); field = ''; rows.push(row); row = [];
    } else field += c;
  }
  if (field !== '' || row.length) { row.push(field); rows.push(row); }
  const [header, ...body] = rows;
  return body
    .filter((r) => r.length > 1 || r[0] !== '')
    .map((r) => Object.fromEntries(header.map((h, j) => [h, coerce(r[j] ?? '')])));
}

/** Names accepted by load(). */
export function datasets() {
  return Object.keys(DATASETS);
}

/** Load one dataset as an array of row objects (one per CSV row). */
export async function load(name) {
  return parseCsv(await readFile(dataUrl(DATASETS[key(name)].file, 'csv'), 'utf8'));
}

/** Synchronous twin of load(). */
export function loadSync(name) {
  return parseCsv(readFileSync(dataUrl(DATASETS[key(name)].file, 'csv'), 'utf8'));
}

/** The dataset's JSON envelope (season, updated, method, venues ...). */
export async function loadJson(name) {
  return JSON.parse(await readFile(dataUrl(DATASETS[key(name)].file, 'json'), 'utf8'));
}

/** Canonical pages in all six locales, licence, row count and version. */
export async function info(name) {
  const k = key(name);
  const meta = DATASETS[k];
  const env = await loadJson(k);
  const pages = Object.fromEntries(LOCALES.map((l) => [l, `${SITE}${l === 'en' ? '' : '/' + l}${meta.path}`]));
  return {
    name: k,
    title: meta.title,
    canonical: pages.en,
    pages,
    rows: (await load(k)).length,
    season: env.season ?? null,
    updated: env.updated ?? null,
    verified: env.verified ?? env.updated ?? null,
    license: DATA_LICENSE,
    license_url: DATA_LICENSE_URL,
    attribution: `Marbella Wire, ${pages.en}`,
    repository: REPOSITORY,
    doi: CONCEPT_DOI,
    package_version: version,
  };
}

export default { load, loadSync, loadJson, info, datasets, parseCsv, version, DATASETS };
