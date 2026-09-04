#!/usr/bin/env node
// Copy the repo's datasets into the four package trees and stamp one CalVer
// version into every version field. Run after each seasonal export:
//
//   node packages/sync-data.mjs --version=2026.9.4      # write
//   node packages/sync-data.mjs --check                 # CI: fail if stale
//
// The repo root (sunbed-index/, chiringuitos/, beach-clubs/) is the only
// source of truth; the package copies are build inputs, never edited by hand.
import { readFileSync, writeFileSync, copyFileSync, existsSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.join(HERE, '..');
const DATASETS = ['sunbed-index', 'chiringuitos', 'beach-clubs'];
// Go embeds data/ with go:embed; R installs inst/extdata. Both expect the same
// flat `<dataset>.<ext>` layout the Python and npm packages use.
const TARGETS = [
  path.join(HERE, 'python/src/marbella_prices/data'),
  path.join(HERE, 'npm/data'),
  path.join(HERE, 'go/data'),
  path.join(HERE, 'r/inst/extdata'),
];
const check = process.argv.includes('--check');
const versionArg = process.argv.find((a) => a.startsWith('--version='))?.slice(10);
if (versionArg && !/^\d{4}\.\d{1,2}\.\d{1,2}$/.test(versionArg)) {
  console.error(`--version must be CalVer YYYY.M.D, got ${versionArg}`);
  process.exit(2);
}

let stale = 0;
for (const ds of DATASETS) {
  for (const ext of ['csv', 'json']) {
    const src = path.join(ROOT, ds, `${ds}.${ext}`);
    for (const dir of TARGETS) {
      const dst = path.join(dir, `${ds}.${ext}`);
      const same = existsSync(dst) && readFileSync(src).equals(readFileSync(dst));
      if (same) continue;
      if (check) { console.error(`stale: ${path.relative(ROOT, dst)}`); stale++; }
      else { copyFileSync(src, dst); console.log(`copied ${path.relative(ROOT, dst)}`); }
    }
  }
}

// One version, three files. Stamped only when --version is given.
const VERSION_FILES = [
  [path.join(HERE, 'python/pyproject.toml'), /^version = "([^"]+)"/m],
  [path.join(HERE, 'python/src/marbella_prices/__init__.py'), /^__version__ = "([^"]+)"/m],
  [path.join(HERE, 'npm/package.json'), /"version": "([^"]+)"/],
  // Go: the module itself is semver (Version), so the CalVer data stamp lives
  // in DataVersion. R: DESCRIPTION Version is unquoted, hence the capture group.
  [path.join(HERE, 'go/marbella.go'), /^\s*DataVersion\s*=\s*"([^"]+)"/m],
  [path.join(HERE, 'r/DESCRIPTION'), /^Version: (\S+)$/m],
];
// Each regex captures the version itself, so one rule stamps every format:
// rewrite the captured text inside the matched line, leaving quotes, keys and
// gofmt's alignment untouched.
const versions = new Set();
for (const [file, re] of VERSION_FILES) {
  const text = readFileSync(file, 'utf8');
  const m = text.match(re);
  if (!m) { console.error(`no version field in ${path.relative(ROOT, file)}`); process.exit(2); }
  if (versionArg && !check) {
    writeFileSync(file, text.replace(re, (line, current) => line.replace(current, versionArg)));
    console.log(`stamped ${versionArg} into ${path.relative(ROOT, file)}`);
  } else {
    versions.add(m[1]);
  }
}
if (check && versions.size > 1) { console.error(`version mismatch: ${[...versions].join(' vs ')}`); stale++; }
if (check && stale) process.exit(1);
if (check) console.log(`ok: package data in sync, version ${[...versions][0]}`);
