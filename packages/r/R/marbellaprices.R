## Marbella Wire open price datasets.
##
## The CSV and JSON files under inst/extdata are the published datasets,
## copied verbatim from the repository root. They are the single copy in this
## package: there is no data/*.rda twin to drift away from them.

MW_SITE <- "https://marbellawire.com"
MW_REPOSITORY <- "https://github.com/shiftdylson1/marbella-price-data"
MW_DOI <- "10.5281/zenodo.22094846"
MW_LICENSE <- "CC-BY-4.0"
MW_LICENSE_URL <- "https://creativecommons.org/licenses/by/4.0/"
MW_LOCALES <- c("en", "es", "de", "fr", "nl", "sv")

.mw_meta <- list(
  sunbed_index = list(
    file = "sunbed-index",
    path = "/prices/sunbeds/",
    title = "Marbella Wire Sunbed Index"
  ),
  chiringuitos = list(
    file = "chiringuitos",
    path = "/prices/chiringuitos/",
    title = "Marbella Wire chiringuito census"
  ),
  beach_clubs = list(
    file = "beach-clubs",
    path = "/prices/beach-clubs/",
    title = "Marbella Wire beach club price census"
  )
)

## Resolve a dataset name. Hyphens, spaces and capitals are all accepted.
.mw_key <- function(dataset) {
  if (!is.character(dataset) || length(dataset) != 1L || is.na(dataset)) {
    stop("`dataset` must be a single dataset name", call. = FALSE)
  }
  key <- gsub("-", "_", tolower(trimws(dataset)), fixed = TRUE)
  if (!key %in% names(.mw_meta)) {
    stop(
      sprintf(
        "unknown dataset \"%s\"; choose from %s",
        dataset, paste(names(.mw_meta), collapse = ", ")
      ),
      call. = FALSE
    )
  }
  key
}

#' Dataset names
#'
#' The dataset names accepted by \code{\link{mw_load}}.
#'
#' @return A character vector: \code{"sunbed_index"}, \code{"chiringuitos"},
#'   \code{"beach_clubs"}.
#' @examples
#' mw_datasets()
#' @export
mw_datasets <- function() names(.mw_meta)

#' Path to a packaged data file
#'
#' The location on disk of one of the installed CSV or JSON files, for reading
#' with a tool of your own.
#'
#' @param dataset Dataset name; see \code{\link{mw_datasets}}.
#' @param ext Either \code{"csv"} or \code{"json"}.
#' @return A single file path.
#' @examples
#' mw_file("sunbed_index", "csv")
#' @export
mw_file <- function(dataset, ext = c("csv", "json")) {
  key <- .mw_key(dataset)
  ext <- match.arg(ext)
  path <- system.file(
    "extdata", paste0(.mw_meta[[key]]$file, ".", ext),
    package = "marbellaprices"
  )
  if (!nzchar(path)) {
    stop("the marbellaprices data files are not installed", call. = FALSE)
  }
  path
}

#' Load a dataset
#'
#' Reads one published dataset into a data frame.
#'
#' An empty cell means the venue publishes no figure for that column. It is
#' read as \code{NA}, never as zero: the datasets never guess and never fill a
#' gap with a placeholder. A zero that is actually present is a real price --
#' the public beach costs nothing to lie on.
#'
#' @param dataset Dataset name: \code{"sunbed_index"}, \code{"chiringuitos"} or
#'   \code{"beach_clubs"}. Hyphens and capitals are accepted.
#' @return A data frame, one row per venue, with character columns left as
#'   character (\code{stringsAsFactors} is off).
#' @seealso \code{\link{mw_info}} for the provenance a citation needs.
#' @examples
#' sunbeds <- mw_load("sunbed_index")
#' sunbeds[!is.na(sunbeds$price_eur), c("venue", "area", "price_eur")]
#'
#' # Venues that publish no rate carry a note saying why.
#' sunbeds[is.na(sunbeds$price_eur), c("venue", "note")]
#' @export
mw_load <- function(dataset) {
  read.csv(
    mw_file(dataset, "csv"),
    stringsAsFactors = FALSE,
    check.names = FALSE,
    na.strings = "",
    encoding = "UTF-8"
  )
}

#' Load a dataset's JSON envelope
#'
#' The published JSON file carries the season, the date the figures were
#' verified, the method, and the venues themselves.
#'
#' @param dataset Dataset name; see \code{\link{mw_datasets}}.
#' @return A list, as returned by \code{jsonlite::fromJSON} with
#'   \code{simplifyVector = FALSE}.
#' @note Requires the \pkg{jsonlite} package, which is only suggested: the rest
#'   of this package has no dependencies beyond base R.
#' @examples
#' if (requireNamespace("jsonlite", quietly = TRUE)) {
#'   env <- mw_json("chiringuitos")
#'   env$season
#' }
#' @export
mw_json <- function(dataset) {
  if (!requireNamespace("jsonlite", quietly = TRUE)) {
    stop(
      "mw_json() needs the jsonlite package: install.packages(\"jsonlite\")",
      call. = FALSE
    )
  }
  jsonlite::fromJSON(mw_file(dataset, "json"), simplifyVector = FALSE)
}

#' Provenance for a dataset
#'
#' Where the figures were published, when they were verified, how many rows
#' there are, and the exact credit line the licence asks for.
#'
#' @param dataset Dataset name; see \code{\link{mw_datasets}}.
#' @return A list with the canonical page, the same page in all six published
#'   languages, the row count, the season, the licence and the DOI.
#' @examples
#' info <- mw_info("beach_clubs")
#' info$canonical
#' info$attribution
#' @export
mw_info <- function(dataset) {
  key <- .mw_key(dataset)
  meta <- .mw_meta[[key]]
  pages <- paste0(
    MW_SITE, ifelse(MW_LOCALES == "en", "", paste0("/", MW_LOCALES)), meta$path
  )
  names(pages) <- MW_LOCALES
  rows <- nrow(mw_load(key))
  env <- if (requireNamespace("jsonlite", quietly = TRUE)) mw_json(key) else list()
  list(
    name = key,
    title = meta$title,
    canonical = unname(pages[["en"]]),
    pages = as.list(pages),
    rows = rows,
    season = env$season,
    updated = env$updated,
    verified = if (is.null(env$verified)) env$updated else env$verified,
    license = MW_LICENSE,
    license_url = MW_LICENSE_URL,
    attribution = paste0("Marbella Wire, ", unname(pages[["en"]])),
    repository = MW_REPOSITORY,
    doi = MW_DOI,
    package_version = as.character(utils::packageVersion("marbellaprices"))
  )
}
