export type DatasetName = 'sunbed_index' | 'chiringuitos' | 'beach_clubs';
export type Locale = 'en' | 'es' | 'de' | 'fr' | 'nl' | 'sv';
export type Cell = string | number | boolean | null;
export type Row = Record<string, Cell>;

export interface SunbedIndexRow extends Row {
  venue: string; area: string; status: 'priced' | 'unpublished';
  price_eur: number | null; price_unit: string; yoy_pct: number | null; note: string | null;
  season: number; measured_on: string; price_source_url: string | null; price_source_via: string | null;
  index_average_eur: number; index_average_unit: string; index_yoy_pct: number | null;
  source_url: string; attribution: string; license: string;
}

export interface DatasetInfo {
  name: DatasetName; title: string; canonical: string; pages: Record<Locale, string>;
  rows: number; season: number | null; updated: string | null; verified: string | null;
  license: 'CC-BY-4.0'; license_url: string; attribution: string; repository: string;
  doi: string; package_version: string;
}

export const version: string;
export const SITE: string;
export const REPOSITORY: string;
export const CONCEPT_DOI: string;
export const DATA_LICENSE: 'CC-BY-4.0';
export const DATA_LICENSE_URL: string;
export const LOCALES: readonly Locale[];
export const DATASETS: Record<DatasetName, { file: string; path: string; title: string }>;

export function datasets(): DatasetName[];
export function parseCsv(text: string): Row[];
export function load(name: 'sunbed_index' | 'sunbed-index'): Promise<SunbedIndexRow[]>;
export function load(name: DatasetName | string): Promise<Row[]>;
export function loadSync(name: 'sunbed_index' | 'sunbed-index'): SunbedIndexRow[];
export function loadSync(name: DatasetName | string): Row[];
export function loadJson(name: DatasetName | string): Promise<Record<string, unknown>>;
export function info(name: DatasetName | string): Promise<DatasetInfo>;

declare const _default: {
  load: typeof load; loadSync: typeof loadSync; loadJson: typeof loadJson; info: typeof info;
  datasets: typeof datasets; parseCsv: typeof parseCsv; version: string; DATASETS: typeof DATASETS;
};
export default _default;
