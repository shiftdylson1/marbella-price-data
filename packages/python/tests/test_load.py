import unittest

import marbella_prices as mp


class LoadTests(unittest.TestCase):
    def test_datasets_listed(self):
        self.assertEqual(mp.datasets(), ["sunbed_index", "chiringuitos", "beach_clubs"])

    def test_sunbed_index_rows_are_typed(self):
        rows = mp.load("sunbed_index")
        self.assertGreater(len(rows), 5)
        priced = [r for r in rows if r["status"] == "priced"]
        self.assertTrue(all(isinstance(r["price_eur"], (int, float)) for r in priced))
        self.assertTrue(all(r["price_eur"] is None for r in rows if r["status"] == "unpublished"))
        self.assertTrue(all(r["license"].startswith("https://creativecommons.org/licenses/by/4.0") for r in rows))

    def test_hyphen_alias_and_unknown(self):
        self.assertEqual(len(mp.load("beach-clubs")), len(mp.load("beach_clubs")))
        with self.assertRaises(KeyError):
            mp.load("hotels")

    def test_json_envelope_matches_csv(self):
        env = mp.load_json("chiringuitos")
        self.assertEqual(env["count"], len(mp.load("chiringuitos")))
        self.assertTrue(all(r["menu_is_pdf"] in (True, False, None) for r in mp.load("chiringuitos")))

    def test_info_pages_cover_six_locales(self):
        i = mp.info("beach_clubs")
        self.assertEqual(sorted(i["pages"]), ["de", "en", "es", "fr", "nl", "sv"])
        self.assertEqual(i["canonical"], "https://marbellawire.com/prices/beach-clubs/")
        self.assertEqual(i["package_version"], mp.__version__)

    def test_text_that_starts_with_digits_stays_text(self):
        row = next(r for r in mp.load("chiringuitos") if r["phone"])
        self.assertIsInstance(row["phone"], str)


if __name__ == "__main__":
    unittest.main()
