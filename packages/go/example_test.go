package marbellaprices_test

import (
	"fmt"
	"sort"

	marbellaprices "github.com/shiftdylson1/marbella-price-data/packages/go"
)

// The cheapest and dearest way to put two people on sunbeds for a weekday in
// high season. Venues that publish no rate are carried with a nil price and a
// note, never as zero.
func ExampleSunbeds() {
	rows, err := marbellaprices.Sunbeds()
	if err != nil {
		panic(err)
	}

	var priced []marbellaprices.SunbedRow
	for _, r := range rows {
		if r.PriceEUR != nil {
			priced = append(priced, r)
		}
	}
	sort.Slice(priced, func(i, j int) bool { return *priced[i].PriceEUR < *priced[j].PriceEUR })

	cheapest, dearest := priced[0], priced[len(priced)-1]
	fmt.Printf("cheapest: %s, %.0f EUR\n", cheapest.Venue, *cheapest.PriceEUR)
	fmt.Printf("dearest:  %s, %.0f EUR\n", dearest.Venue, *dearest.PriceEUR)
	fmt.Printf("%d of %d venues publish a rate\n", len(priced), len(rows))

	// Output:
	// cheapest: Public beach, 0 EUR
	// dearest:  Ocean Club, 300 EUR
	// 8 of 12 venues publish a rate
}

// Info carries the provenance a citation needs.
func ExampleInfo() {
	info, err := marbellaprices.Info("beach-clubs")
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Title)
	fmt.Println(info.Canonical)
	fmt.Println(info.Attribution)
	fmt.Println(info.License)

	// Output:
	// Marbella Wire beach club price census
	// https://marbellawire.com/prices/beach-clubs/
	// Marbella Wire, https://marbellawire.com/prices/beach-clubs/
	// CC-BY-4.0
}

// Load returns plain rows, the same shape as the Python and npm packages.
func ExampleLoad() {
	rows, err := marbellaprices.Load("chiringuitos")
	if err != nil {
		panic(err)
	}
	withEspeto := 0
	for _, r := range rows {
		if r["espeto_eur"] != nil {
			withEspeto++
		}
	}
	fmt.Printf("datasets: %v\n", marbellaprices.Names())
	fmt.Printf("%d chiringuitos, %d publish an espeto price\n", len(rows), withEspeto)

	// Output:
	// datasets: [beach_clubs chiringuitos sunbed_index]
	// 43 chiringuitos, 11 publish an espeto price
}
