## Dependency-free checks that run under `R CMD check`.
library(marbellaprices)

expect <- function(ok, what) if (!isTRUE(ok)) stop("FAILED: ", what, call. = FALSE)

## Names and name resolution -------------------------------------------------
expect(identical(mw_datasets(), c("sunbed_index", "chiringuitos", "beach_clubs")),
       "mw_datasets() lists the three datasets")
for (alias in c("beach-clubs", "BEACH_CLUBS", "  beach_clubs  ")) {
  expect(identical(mw_info(alias)$name, "beach_clubs"),
         paste("alias resolves:", alias))
}
expect(inherits(try(mw_load("sunbeds"), silent = TRUE), "try-error"),
       "an unknown dataset is an error")

## Every dataset loads, and every row states its own licence ------------------
for (ds in mw_datasets()) {
  d <- mw_load(ds)
  expect(is.data.frame(d) && nrow(d) > 0L, paste(ds, "loads rows"))
  expect(all(d$license == "https://creativecommons.org/licenses/by/4.0/"),
         paste(ds, "rows carry the CC BY 4.0 licence"))
  expect(!any(vapply(d, is.factor, logical(1))), paste(ds, "has no factor columns"))
}

## A blank is NA, never zero -------------------------------------------------
s <- mw_load("sunbed_index")
expect(is.numeric(s$price_eur), "price_eur is numeric")
expect(all(is.na(s$price_eur[s$status == "unpublished"])),
       "unpublished rows have no price")
expect(!any(is.na(s$price_eur[s$status == "priced"])),
       "priced rows all carry a figure")
expect(all(nzchar(s$note[s$status == "unpublished"])),
       "unpublished rows say why")
expect(all(s$price_eur[s$status == "priced"] >= 0),
       "no negative prices")
expect(any(s$price_eur[s$status == "priced"] == 0),
       "the free public beach is carried as a real zero")

## Censuses ------------------------------------------------------------------
chi <- mw_load("chiringuitos")
expect(all(nzchar(chi$slug)) && !anyDuplicated(chi$slug), "chiringuito slugs are unique")
expect(all(grepl("^https://marbellawire\\.com/prices/chiringuitos/", chi$url)),
       "chiringuito rows link their canonical venue page")

bc <- mw_load("beach_clubs")
expect(all(bc$season_year >= 2020), "beach club season years are plausible")

## Provenance ----------------------------------------------------------------
info <- mw_info("sunbed-index")
expect(identical(info$canonical, "https://marbellawire.com/prices/sunbeds/"),
       "canonical page")
expect(identical(info$pages$es, "https://marbellawire.com/es/prices/sunbeds/"),
       "Spanish page")
expect(length(info$pages) == 6L, "six locales")
expect(identical(info$rows, nrow(s)), "row count matches")
expect(identical(info$license, "CC-BY-4.0"), "licence")
expect(grepl("Marbella Wire", info$attribution, fixed = TRUE), "attribution names the entity")

## File access ---------------------------------------------------------------
expect(file.exists(mw_file("chiringuitos", "csv")), "csv path exists")
expect(file.exists(mw_file("chiringuitos", "json")), "json path exists")
expect(inherits(try(mw_file("chiringuitos", "xlsx"), silent = TRUE), "try-error"),
       "an unknown extension is an error")

if (requireNamespace("jsonlite", quietly = TRUE)) {
  env <- mw_json("beach_clubs")
  expect(!is.null(env$season) && !is.null(env$clubs), "JSON envelope has season and clubs")
}

cat("all marbellaprices checks passed\n")
