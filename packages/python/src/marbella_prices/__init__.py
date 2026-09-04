"""Marbella Wire open price datasets, embedded as CSV/JSON, no dependencies.

    >>> from marbella_prices import load
    >>> rows = load("sunbed_index")
    >>> rows[0]["venue"], rows[0]["price_eur"]
    ('Ocean Club', 300)

Every figure was published first on https://marbellawire.com/prices/ and is
CC BY 4.0 (credit "Marbella Wire" + the canonical page). The loader code is
MIT. Data is regenerated once a season; the live endpoints at
https://marbellawire.com/data/ are always current.
"""

from __future__ import annotations

import csv
import json
from importlib import resources
from typing import Any, Dict, List

__version__ = "2026.9.4"
__all__ = ["load", "load_json", "datasets", "info", "Dataset", "DATASETS", "__version__"]

SITE = "https://marbellawire.com"
REPOSITORY = "https://github.com/shiftdylson1/marbella-price-data"
CONCEPT_DOI = "10.5281/zenodo.22094846"
DATA_LICENSE = "CC-BY-4.0"
DATA_LICENSE_URL = "https://creativecommons.org/licenses/by/4.0/"
LOCALES = ("en", "es", "de", "fr", "nl", "sv")

# name -> (file stem, canonical path, short title)
DATASETS: Dict[str, Dict[str, str]] = {
    "sunbed_index": {"file": "sunbed-index", "path": "/prices/sunbeds/", "title": "Marbella Wire Sunbed Index"},
    "chiringuitos": {"file": "chiringuitos", "path": "/prices/chiringuitos/", "title": "Marbella Wire chiringuito census"},
    "beach_clubs": {"file": "beach-clubs", "path": "/prices/beach-clubs/", "title": "Marbella Wire beach club price census"},
}


class Dataset(List[Dict[str, Any]]):
    """A list of row dicts that knows where it came from."""

    name: str = ""

    def to_dataframe(self):  # type: ignore[no-untyped-def]
        """Return the rows as a pandas DataFrame (pandas is optional, imported lazily)."""
        try:
            import pandas as pd
        except ImportError as e:  # pragma: no cover
            raise ImportError("pandas is not installed; pip install 'marbella-prices[pandas]'") from e
        return pd.DataFrame(list(self))

    def __repr__(self) -> str:
        return f"<Dataset {self.name}: {len(self)} rows>"


def _key(name: str) -> str:
    key = name.strip().lower().replace("-", "_")
    if key not in DATASETS:
        raise KeyError(f"unknown dataset {name!r}; choose from {', '.join(DATASETS)}")
    return key


def _coerce(value: str) -> Any:
    """CSV cells are text; give numbers, booleans and blanks their obvious type.

    Empty = not published (never zero, never a guess), so it becomes None.
    Text that merely starts with digits (phones, addresses, "€8.90") stays text.
    """
    if value == "":
        return None
    low = value.lower()
    if low == "true":
        return True
    if low == "false":
        return False
    try:
        return int(value)
    except ValueError:
        pass
    try:
        f = float(value)
    except ValueError:
        return value
    return f if value.replace("-", "").replace(".", "").isdigit() else value


def _read_text(stem: str, ext: str) -> str:
    return resources.files(__name__).joinpath("data", f"{stem}.{ext}").read_text(encoding="utf-8")


def datasets() -> List[str]:
    """Names accepted by load(): ['sunbed_index', 'chiringuitos', 'beach_clubs']."""
    return list(DATASETS)


def load(name: str) -> Dataset:
    """Load one dataset as a list of dicts (one per row of the CSV).

    Args:
        name: 'sunbed_index', 'chiringuitos' or 'beach_clubs' (hyphens accepted).
    """
    key = _key(name)
    reader = csv.DictReader(_read_text(DATASETS[key]["file"], "csv").splitlines())
    ds = Dataset({k: _coerce(v) for k, v in row.items()} for row in reader)
    ds.name = key
    return ds


def load_json(name: str) -> Dict[str, Any]:
    """The dataset's JSON envelope (season, updated, method/methodology, venues ...)."""
    return json.loads(_read_text(DATASETS[_key(name)]["file"], "json"))


def info(name: str) -> Dict[str, Any]:
    """Canonical pages (all six locales), licence, row count and version for a dataset."""
    key = _key(name)
    meta = DATASETS[key]
    env = load_json(key)
    pages = {loc: f"{SITE}{'' if loc == 'en' else '/' + loc}{meta['path']}" for loc in LOCALES}
    return {
        "name": key,
        "title": meta["title"],
        "canonical": pages["en"],
        "pages": pages,
        "rows": len(load(key)),
        "season": env.get("season"),
        "updated": env.get("updated"),
        "verified": env.get("verified", env.get("updated")),
        "license": DATA_LICENSE,
        "license_url": DATA_LICENSE_URL,
        "attribution": f"Marbella Wire, {pages['en']}",
        "repository": REPOSITORY,
        "doi": CONCEPT_DOI,
        "package_version": __version__,
    }

