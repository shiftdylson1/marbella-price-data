package marbellaprices

import (
	"reflect"
	"strings"
	"testing"
)

// tagsOf returns the `csv` tags of T in declaration order.
func tagsOf(t *testing.T, v any) []string {
	t.Helper()
	typ := reflect.TypeOf(v)
	var out []string
	for i := 0; i < typ.NumField(); i++ {
		if tag := typ.Field(i).Tag.Get("csv"); tag != "" && tag != "-" {
			out = append(out, tag)
		}
	}
	return out
}

func headerOf(t *testing.T, name string) []string {
	t.Helper()
	d, err := Lookup(name)
	if err != nil {
		t.Fatal(err)
	}
	header, _, err := readCSV(d.File)
	if err != nil {
		t.Fatal(err)
	}
	return header
}

// The typed structs must describe the published CSV exactly. A seasonal export
// that adds or renames a column has to be reflected here, not silently dropped.
func TestStructsMatchPublishedColumns(t *testing.T) {
	for _, tc := range []struct {
		name string
		zero any
	}{
		{"sunbed_index", SunbedRow{}},
		{"chiringuitos", ChiringuitoRow{}},
		{"beach_clubs", BeachClubRow{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tagsOf(t, tc.zero), headerOf(t, tc.name)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("csv tags do not match the published header\n tags:   %v\n header: %v", got, want)
			}
		})
	}
}

func TestNamesAndLookup(t *testing.T) {
	want := []string{"beach_clubs", "chiringuitos", "sunbed_index"}
	if got := Names(); !reflect.DeepEqual(got, want) {
		t.Errorf("Names() = %v, want %v", got, want)
	}
	// Hyphens, spaces and case all resolve to the same dataset.
	for _, alias := range []string{"beach-clubs", "Beach_Clubs", "  beach_clubs  ", "BEACH-CLUBS"} {
		d, err := Lookup(alias)
		if err != nil {
			t.Fatalf("Lookup(%q): %v", alias, err)
		}
		if d.Name != "beach_clubs" {
			t.Errorf("Lookup(%q).Name = %q, want beach_clubs", alias, d.Name)
		}
	}
	if _, err := Lookup("sunbeds"); err == nil {
		t.Error("Lookup of an unknown dataset should fail")
	} else if !strings.Contains(err.Error(), "beach_clubs") {
		t.Errorf("the error should list the valid names, got %v", err)
	}
}

func TestLoadRows(t *testing.T) {
	for _, name := range Names() {
		rows, err := Load(name)
		if err != nil {
			t.Fatalf("Load(%q): %v", name, err)
		}
		if len(rows) == 0 {
			t.Fatalf("Load(%q) returned no rows", name)
		}
		header := headerOf(t, name)
		for i, r := range rows {
			if len(r) != len(header) {
				t.Fatalf("%s row %d has %d keys, want %d", name, i, len(r), len(header))
			}
			// Every row carries its own licence, so a row that escapes the
			// repo still states its terms.
			if got := r["license"]; got != DataLicenseURL {
				t.Errorf("%s row %d license = %v, want %s", name, i, got, DataLicenseURL)
			}
		}
	}
}

// An empty cell means "not published". It must never arrive as zero.
func TestUnpublishedIsNilNeverZero(t *testing.T) {
	rows, err := Sunbeds()
	if err != nil {
		t.Fatal(err)
	}
	var priced, unpublished int
	for _, r := range rows {
		switch r.Status {
		case "priced":
			priced++
			// A priced row always carries a figure, and that figure may
			// legitimately be zero: the public beach costs nothing to lie on.
			// Zero is only forbidden as a stand-in for "not published", which
			// is what the nil check on unpublished rows below guards.
			if r.PriceEUR == nil {
				t.Errorf("%s is priced but PriceEUR is nil", r.Venue)
			} else if *r.PriceEUR < 0 {
				t.Errorf("%s has a negative price %v", r.Venue, *r.PriceEUR)
			}
		case "unpublished":
			unpublished++
			if r.PriceEUR != nil {
				t.Errorf("%s is unpublished but carries a price %v", r.Venue, *r.PriceEUR)
			}
			if r.Note == "" {
				t.Errorf("%s is unpublished with no note explaining why", r.Venue)
			}
		default:
			t.Errorf("%s has unexpected status %q", r.Venue, r.Status)
		}
		if r.Season == 0 {
			t.Errorf("%s has no season", r.Venue)
		}
	}
	if priced == 0 || unpublished == 0 {
		t.Fatalf("expected both priced and unpublished rows, got %d/%d", priced, unpublished)
	}
}

func TestTypedCensuses(t *testing.T) {
	chi, err := Chiringuitos()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range chi {
		if c.Slug == "" || c.Name == "" {
			t.Errorf("chiringuito row with no slug/name: %+v", c)
		}
		if !strings.HasPrefix(c.URL, Site+"/prices/chiringuitos/") {
			t.Errorf("%s: URL %q is not a canonical venue page", c.Slug, c.URL)
		}
		if c.EspetoEUR != nil && *c.EspetoEUR <= 0 {
			t.Errorf("%s: espeto price %v", c.Slug, *c.EspetoEUR)
		}
	}

	bc, err := BeachClubs()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range bc {
		if c.SeasonYear < 2020 {
			t.Errorf("%s: implausible season_year %d", c.Slug, c.SeasonYear)
		}
		// rates_is_pdf is false whenever no rates document is published at
		// all, so only a true flag implies a URL to point at.
		if c.RatesIsPDF != nil && *c.RatesIsPDF && c.RatesURL == "" {
			t.Errorf("%s: rates_is_pdf is true with no rates_url", c.Slug)
		}
		if c.MenuIsPDF != nil && *c.MenuIsPDF && c.MenuURL == "" {
			t.Errorf("%s: menu_is_pdf is true with no menu_url", c.Slug)
		}
	}
}

// Text that merely starts with digits (phones, addresses) must stay text.
func TestCoerce(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want any
	}{
		{"", nil},
		{"true", true},
		{"false", false},
		{"300", 300.0},
		{"8.9", 8.9},
		{"-3.5", -3.5},
		{"952 85 74 48", "952 85 74 48"},
		{"3-4, P.º Marítimo", "3-4, P.º Marítimo"},
		{"€8.90", "€8.90"},
		{"2026-08-11", "2026-08-11"},
		{"1e5", "1e5"},
		{"0x10", "0x10"},
		{"Inf", "Inf"},
		{"NaN", "NaN"},
		{".", "."},
	} {
		if got := coerce(tc.in); got != tc.want {
			t.Errorf("coerce(%q) = %#v, want %#v", tc.in, got, tc.want)
		}
	}
}

func TestInfo(t *testing.T) {
	info, err := Info("sunbed-index")
	if err != nil {
		t.Fatal(err)
	}
	if info.Name != "sunbed_index" {
		t.Errorf("Name = %q", info.Name)
	}
	if info.Canonical != "https://marbellawire.com/prices/sunbeds/" {
		t.Errorf("Canonical = %q", info.Canonical)
	}
	if info.Pages["es"] != "https://marbellawire.com/es/prices/sunbeds/" {
		t.Errorf("Pages[es] = %q", info.Pages["es"])
	}
	if len(info.Pages) != len(Locales) {
		t.Errorf("got %d pages, want %d", len(info.Pages), len(Locales))
	}
	rows, _ := Load("sunbed_index")
	if info.Rows != len(rows) {
		t.Errorf("Rows = %d, want %d", info.Rows, len(rows))
	}
	if info.License != "CC-BY-4.0" || info.DOI != ConceptDOI {
		t.Errorf("licence/DOI wrong: %q %q", info.License, info.DOI)
	}
	if !strings.Contains(info.Attribution, "Marbella Wire") {
		t.Errorf("Attribution = %q", info.Attribution)
	}
	if info.Season == nil || info.Updated == nil {
		t.Errorf("season/updated missing: %v %v", info.Season, info.Updated)
	}
}

func TestLoadJSONAndRaw(t *testing.T) {
	var env map[string]any
	if err := LoadJSON("chiringuitos", &env); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"name", "season", "updated", "venues", "license"} {
		if _, ok := env[k]; !ok {
			t.Errorf("envelope is missing %q", k)
		}
	}

	csvBytes, err := Raw("beach_clubs", "csv")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(string(csvBytes), "slug,name,area,season_year,") {
		t.Errorf("raw CSV does not start with the published header")
	}
	if _, err := Raw("beach_clubs", "xlsx"); err == nil {
		t.Error("Raw should reject an unknown extension")
	}
	if err := LoadJSON("nope", &env); err == nil {
		t.Error("LoadJSON should reject an unknown dataset")
	}
}

func TestErrorsAreNamespaced(t *testing.T) {
	_, err := Load("nope")
	if err == nil || !strings.HasPrefix(err.Error(), "marbellaprices: ") {
		t.Errorf("error should be namespaced, got %v", err)
	}
}
