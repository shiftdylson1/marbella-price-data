import { test } from 'node:test';
import assert from 'node:assert/strict';
import { load, loadSync, loadJson, info, datasets, parseCsv, version } from '../index.js';
import { readFileSync } from 'node:fs';

test('datasets listed', () => {
  assert.deepEqual(datasets(), ['sunbed_index', 'chiringuitos', 'beach_clubs']);
});

test('sunbed index rows are typed and licensed', async () => {
  const rows = await load('sunbed_index');
  assert.ok(rows.length > 5);
  for (const r of rows) {
    assert.ok(r.license.startsWith('https://creativecommons.org/licenses/by/4.0'));
    if (r.status === 'priced') assert.equal(typeof r.price_eur, 'number');
    else assert.equal(r.price_eur, null);
  }
});

test('hyphen alias, sync twin, unknown name', async () => {
  assert.equal(loadSync('beach-clubs').length, (await load('beach_clubs')).length);
  assert.throws(() => loadSync('hotels'), /unknown dataset/);
});

test('json envelope count matches csv rows', async () => {
  const env = await loadJson('chiringuitos');
  assert.equal(env.count, (await load('chiringuitos')).length);
});

test('info covers six locales and matches package version', async () => {
  const i = await info('beach_clubs');
  assert.deepEqual(Object.keys(i.pages).sort(), ['de', 'en', 'es', 'fr', 'nl', 'sv']);
  assert.equal(i.canonical, 'https://marbellawire.com/prices/beach-clubs/');
  assert.equal(i.package_version, version);
  assert.equal(JSON.parse(readFileSync(new URL('../package.json', import.meta.url), 'utf8')).version, version);
});

test('csv parser: quotes, doubled quotes, embedded newline, CRLF', () => {
  const rows = parseCsv('a,b,c\r\n1,"x ""y""","line1\nline2"\r\n,true,\r\n');
  assert.deepEqual(rows, [
    { a: 1, b: 'x "y"', c: 'line1\nline2' },
    { a: null, b: true, c: null },
  ]);
});

test('phones stay text', async () => {
  const row = (await load('chiringuitos')).find((r) => r.phone);
  assert.equal(typeof row.phone, 'string');
});
