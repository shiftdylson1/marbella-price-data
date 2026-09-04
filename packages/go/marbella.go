// Package marbellaprices serves the Marbella Wire open price datasets — the
// Sunbed Index, the chiringuito census and the beach club census — from data
// embedded in the binary. No network, no dependencies outside the standard
// library.
//
//	rows, err := marbellaprices.Sunbeds()
//	for _, r := range rows {
//		if r.PriceEUR != nil {
//			fmt.Println(r.Venue, *r.PriceEUR)
//		}
//	}
//
// Untyped access mirrors the Python and npm packages, which return plain rows:
//
//	rows, err := marbellaprices.Load("sunbed_index") // []Row, a map per CSV row
//
// An empty cell means the venue does not publish that figure. It is never zero
// and never a guess, so every optional number is a pointer and stays nil.
//
// Every figure was published first on https://marbellawire.com/prices/ and is
// CC BY 4.0: credit "Marbella Wire" and link the canonical page (Info gives you
// both). This loader code is MIT. The live endpoints at
// https://marbellawire.com/data/ are always current; the embedded copy is a
// seasonal snapshot stamped in DataVersion.
package marbellaprices

import (
	"embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//go:embed data/*.csv data/*.json
var files embed.FS

// Version is the version of this Go module. DataVersion is the CalVer stamp of
// the embedded snapshot, shared with the PyPI and npm packages.
const (
	Version     = "0.1.0"
	DataVersion = "2026.9.4"

	Site           = "https://marbellawire.com"
	Repository     = "https://github.com/shiftdylson1/marbella-price-data"
	ConceptDOI     = "10.5281/zenodo.22094846"
	DataLicense    = "CC-BY-4.0"
	DataLicenseURL = "https://creativecommons.org/licenses/by/4.0/"
)

// Locales are the languages the canonical pages are published in.
var Locales = []string{"en", "es", "de", "fr", "nl", "sv"}

// Dataset describes one published dataset.
type Dataset struct {
	Name  string // key accepted by Load: "sunbed_index", "chiringuitos", "beach_clubs"
	File  string // file stem inside data/
	Path  string // canonical path on marbellawire.com
	Title string
}

var datasets = map[string]Dataset{
	"sunbed_index": {"sunbed_index", "sunbed-index", "/prices/sunbeds/", "Marbella Wire Sunbed Index"},
	"chiringuitos": {"chiringuitos", "chiringuitos", "/prices/chiringuitos/", "Marbella Wire chiringuito census"},
	"beach_clubs":  {"beach_clubs", "beach-clubs", "/prices/beach-clubs/", "Marbella Wire beach club price census"},
}

// Names returns the dataset keys accepted by Load, in a stable order.
func Names() []string {
	out := make([]string, 0, len(datasets))
	for k := range datasets {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// Lookup resolves a dataset name. Hyphens and case are accepted, so
// "beach-clubs", "Beach_Clubs" and "beach_clubs" are the same dataset.
func Lookup(name string) (Dataset, error) {
	key := strings.ReplaceAll(strings.ToLower(strings.TrimSpace(name)), "-", "_")
	d, ok := datasets[key]
	if !ok {
		return Dataset{}, fmt.Errorf("marbellaprices: unknown dataset %q; choose from %s",
			name, strings.Join(Names(), ", "))
	}
	return d, nil
}

// Row is one CSV row. Values are string, float64, bool or nil; nil means the
// venue publishes no figure for that column.
type Row map[string]any

// coerce gives a CSV cell its obvious type. Text that merely starts with
// digits (phone numbers, "€8.90", "3-4, P.º Marítimo") stays a string.
func coerce(v string) any {
	switch v {
	case "":
		return nil
	case "true":
		return true
	case "false":
		return false
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil && isPlainNumber(v) {
		return f
	}
	return v
}

// isPlainNumber reports whether v is only an optional sign, digits and at most
// one decimal point — excluding the exponent and hex forms ParseFloat accepts.
func isPlainNumber(v string) bool {
	s := strings.TrimPrefix(v, "-")
	if s == "" {
		return false
	}
	dots := 0
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
		case r == '.':
			dots++
			if dots > 1 {
				return false
			}
		default:
			return false
		}
	}
	return s != "."
}

// readCSV returns the header and the data records of an embedded CSV.
func readCSV(stem string) ([]string, [][]string, error) {
	f, err := files.Open("data/" + stem + ".csv")
	if err != nil {
		return nil, nil, fmt.Errorf("marbellaprices: %w", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	header, err := r.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("marbellaprices: reading %s.csv header: %w", stem, err)
	}
	var records [][]string
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("marbellaprices: reading %s.csv: %w", stem, err)
		}
		records = append(records, rec)
	}
	return header, records, nil
}

// Load returns a dataset as plain rows, one map per CSV row — the same shape
// the Python and npm packages return. Use Sunbeds, Chiringuitos or BeachClubs
// when you want static types.
func Load(name string) ([]Row, error) {
	d, err := Lookup(name)
	if err != nil {
		return nil, err
	}
	header, records, err := readCSV(d.File)
	if err != nil {
		return nil, err
	}
	rows := make([]Row, 0, len(records))
	for _, rec := range records {
		row := make(Row, len(header))
		for i, col := range header {
			if i < len(rec) {
				row[col] = coerce(rec[i])
			} else {
				row[col] = nil
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// LoadJSON unmarshals the dataset's JSON envelope (season, updated, method,
// venues ...) into v, which is typically *map[string]any or your own struct.
func LoadJSON(name string, v any) error {
	d, err := Lookup(name)
	if err != nil {
		return err
	}
	b, err := files.ReadFile("data/" + d.File + ".json")
	if err != nil {
		return fmt.Errorf("marbellaprices: %w", err)
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("marbellaprices: decoding %s.json: %w", d.File, err)
	}
	return nil
}

// Raw returns the embedded bytes of one file, e.g. Raw("sunbed_index", "csv").
// Useful for piping the dataset straight to disk or to a CSV reader of your own.
func Raw(name, ext string) ([]byte, error) {
	d, err := Lookup(name)
	if err != nil {
		return nil, err
	}
	if ext != "csv" && ext != "json" {
		return nil, fmt.Errorf("marbellaprices: ext must be \"csv\" or \"json\", got %q", ext)
	}
	b, err := files.ReadFile("data/" + d.File + "." + ext)
	if err != nil {
		return nil, fmt.Errorf("marbellaprices: %w", err)
	}
	return b, nil
}

// DatasetInfo is the provenance a citation needs: where the figures were
// published, when they were verified, and the licence terms.
type DatasetInfo struct {
	Name           string            `json:"name"`
	Title          string            `json:"title"`
	Canonical      string            `json:"canonical"`
	Pages          map[string]string `json:"pages"`
	Rows           int               `json:"rows"`
	Season         any               `json:"season"`
	Updated        any               `json:"updated"`
	Verified       any               `json:"verified"`
	License        string            `json:"license"`
	LicenseURL     string            `json:"license_url"`
	Attribution    string            `json:"attribution"`
	Repository     string            `json:"repository"`
	DOI            string            `json:"doi"`
	PackageVersion string            `json:"package_version"`
	DataVersion    string            `json:"data_version"`
}

// Info returns the canonical pages in all six locales, the licence, the row
// count and the season for one dataset.
func Info(name string) (DatasetInfo, error) {
	d, err := Lookup(name)
	if err != nil {
		return DatasetInfo{}, err
	}
	var env map[string]any
	if err := LoadJSON(d.Name, &env); err != nil {
		return DatasetInfo{}, err
	}
	rows, err := Load(d.Name)
	if err != nil {
		return DatasetInfo{}, err
	}
	pages := make(map[string]string, len(Locales))
	for _, l := range Locales {
		prefix := ""
		if l != "en" {
			prefix = "/" + l
		}
		pages[l] = Site + prefix + d.Path
	}
	verified, ok := env["verified"]
	if !ok {
		verified = env["updated"]
	}
	return DatasetInfo{
		Name:           d.Name,
		Title:          d.Title,
		Canonical:      pages["en"],
		Pages:          pages,
		Rows:           len(rows),
		Season:         env["season"],
		Updated:        env["updated"],
		Verified:       verified,
		License:        DataLicense,
		LicenseURL:     DataLicenseURL,
		Attribution:    "Marbella Wire, " + pages["en"],
		Repository:     Repository,
		DOI:            ConceptDOI,
		PackageVersion: Version,
		DataVersion:    DataVersion,
	}, nil
}

// --- Typed rows ----------------------------------------------------------
//
// A pointer field is nil when the venue publishes no figure. Reading it as
// zero would invent a price, which the dataset never does.

// SunbedRow is one row of the Sunbed Index: the cost of putting two people on
// sunbeds for a weekday in high season.
type SunbedRow struct {
	Venue            string   `csv:"venue"`
	Area             string   `csv:"area"`
	Status           string   `csv:"status"` // "priced" or "unpublished"
	PriceEUR         *float64 `csv:"price_eur"`
	PriceUnit        string   `csv:"price_unit"`
	YoYPct           *float64 `csv:"yoy_pct"`
	Note             string   `csv:"note"`
	Season           int      `csv:"season"`
	MeasuredOn       string   `csv:"measured_on"` // YYYY-MM-DD
	PriceSourceURL   string   `csv:"price_source_url"`
	PriceSourceVia   string   `csv:"price_source_via"`
	IndexAverageEUR  *float64 `csv:"index_average_eur"`
	IndexAverageUnit string   `csv:"index_average_unit"`
	IndexYoYPct      *float64 `csv:"index_yoy_pct"`
	SourceURL        string   `csv:"source_url"`
	Attribution      string   `csv:"attribution"`
	License          string   `csv:"license"`
}

// ChiringuitoRow is one beach restaurant in the chiringuito census.
type ChiringuitoRow struct {
	Slug             string   `csv:"slug"`
	Name             string   `csv:"name"`
	Area             string   `csv:"area"`
	Beach            string   `csv:"beach"`
	BusinessStatus   string   `csv:"business_status"`
	OpenedYear       *int     `csv:"opened_year"`
	PriceDisplay     string   `csv:"price_display"`
	PriceBasis       string   `csv:"price_basis"`
	PriceKind        string   `csv:"price_kind"`
	PriceNote        string   `csv:"price_note"`
	PriceLevel       string   `csv:"price_level"`
	SignatureDish    string   `csv:"signature_dish"`
	SignatureDishEUR *float64 `csv:"signature_dish_eur"`
	EspetoEUR        *float64 `csv:"espeto_eur"`
	OfficialSite     string   `csv:"official_site"`
	MenuURL          string   `csv:"menu_url"`
	MenuIsPDF        *bool    `csv:"menu_is_pdf"`
	Address          string   `csv:"address"`
	Phone            string   `csv:"phone"`
	GooglePlaceID    string   `csv:"google_place_id"`
	GoogleMapsURL    string   `csv:"google_maps_url"`
	Verified         string   `csv:"verified"`
	URL              string   `csv:"url"`
	Currency         string   `csv:"currency"`
	SourcePage       string   `csv:"source_page"`
	Attribution      string   `csv:"attribution"`
	License          string   `csv:"license"`
}

// BeachClubRow is one beach club and its published day rates.
type BeachClubRow struct {
	Slug           string   `csv:"slug"`
	Name           string   `csv:"name"`
	Area           string   `csv:"area"`
	SeasonYear     int      `csv:"season_year"`
	SunbedPrice    *float64 `csv:"sunbed_price"`
	BedOrVIPPrice  *float64 `csv:"bed_or_vip_price"`
	MinimumSpend   string   `csv:"minimum_spend"`
	EntryPolicy    string   `csv:"entry_policy"`
	PriceLevel     string   `csv:"price_level"`
	SourceURL      string   `csv:"source_url"`
	RatesURL       string   `csv:"rates_url"`
	RatesIsPDF     *bool    `csv:"rates_is_pdf"`
	BookingURL     string   `csv:"booking_url"`
	OfficialSite   string   `csv:"official_site"`
	MenuURL        string   `csv:"menu_url"`
	MenuIsPDF      *bool    `csv:"menu_is_pdf"`
	OpenedYear     *int     `csv:"opened_year"`
	BusinessStatus string   `csv:"business_status"`
	Address        string   `csv:"address"`
	Phone          string   `csv:"phone"`
	GooglePlaceID  string   `csv:"google_place_id"`
	GoogleMapsURL  string   `csv:"google_maps_url"`
	Verified       string   `csv:"verified"`
	Notes          string   `csv:"notes"`
	URL            string   `csv:"url"`
	Currency       string   `csv:"currency"`
	SourcePage     string   `csv:"source_page"`
	Attribution    string   `csv:"attribution"`
	License        string   `csv:"license"`
}

// Sunbeds returns the Sunbed Index as typed rows.
func Sunbeds() ([]SunbedRow, error) { return decode[SunbedRow]("sunbed_index") }

// Chiringuitos returns the chiringuito census as typed rows.
func Chiringuitos() ([]ChiringuitoRow, error) { return decode[ChiringuitoRow]("chiringuitos") }

// BeachClubs returns the beach club census as typed rows.
func BeachClubs() ([]BeachClubRow, error) { return decode[BeachClubRow]("beach_clubs") }

// decode maps CSV columns onto struct fields by their `csv` tag. A column with
// no matching field is ignored; a tagged field with no matching column is an
// error, because it means the struct and the published schema have drifted.
func decode[T any](name string) ([]T, error) {
	d, err := Lookup(name)
	if err != nil {
		return nil, err
	}
	header, records, err := readCSV(d.File)
	if err != nil {
		return nil, err
	}
	index := make(map[string]int, len(header))
	for i, col := range header {
		index[col] = i
	}

	var zero T
	typ := reflect.TypeOf(zero)
	type binding struct {
		field int
		col   int
	}
	var bindings []binding
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("csv")
		if tag == "" || tag == "-" {
			continue
		}
		col, ok := index[tag]
		if !ok {
			return nil, fmt.Errorf("marbellaprices: %s.csv has no column %q for %s.%s",
				d.File, tag, typ.Name(), typ.Field(i).Name)
		}
		bindings = append(bindings, binding{i, col})
	}

	out := make([]T, 0, len(records))
	for n, rec := range records {
		var row T
		v := reflect.ValueOf(&row).Elem()
		for _, b := range bindings {
			cell := ""
			if b.col < len(rec) {
				cell = rec[b.col]
			}
			if err := setField(v.Field(b.field), cell); err != nil {
				return nil, fmt.Errorf("marbellaprices: %s.csv row %d, column %q: %w",
					d.File, n+2, header[b.col], err)
			}
		}
		out = append(out, row)
	}
	return out, nil
}

// setField writes one CSV cell into one struct field. An empty cell leaves a
// pointer nil (the figure is not published) and a value type at its zero value.
func setField(f reflect.Value, cell string) error {
	if f.Kind() == reflect.Pointer {
		if cell == "" {
			f.SetZero()
			return nil
		}
		p := reflect.New(f.Type().Elem())
		if err := setField(p.Elem(), cell); err != nil {
			return err
		}
		f.Set(p)
		return nil
	}
	switch f.Kind() {
	case reflect.String:
		f.SetString(cell)
	case reflect.Float64, reflect.Float32:
		if cell == "" {
			f.SetFloat(0)
			return nil
		}
		n, err := strconv.ParseFloat(cell, 64)
		if err != nil {
			return fmt.Errorf("%q is not a number", cell)
		}
		f.SetFloat(n)
	case reflect.Int, reflect.Int64, reflect.Int32:
		if cell == "" {
			f.SetInt(0)
			return nil
		}
		n, err := strconv.ParseInt(cell, 10, 64)
		if err != nil {
			return fmt.Errorf("%q is not an integer", cell)
		}
		f.SetInt(n)
	case reflect.Bool:
		if cell == "" {
			f.SetBool(false)
			return nil
		}
		b, err := strconv.ParseBool(cell)
		if err != nil {
			return fmt.Errorf("%q is not a boolean", cell)
		}
		f.SetBool(b)
	default:
		return fmt.Errorf("unsupported field kind %s", f.Kind())
	}
	return nil
}
