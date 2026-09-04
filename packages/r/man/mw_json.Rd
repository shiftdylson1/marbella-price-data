\name{mw_json}
\alias{mw_json}
\title{Load a dataset's JSON envelope}
\usage{
mw_json(dataset)
}
\arguments{
\item{dataset}{Dataset name; see \code{\link{mw_datasets}}.}
}
\value{
A list, as returned by \code{jsonlite::fromJSON} with
\code{simplifyVector = FALSE}.
}
\description{
The published JSON file carries the season, the date the figures were
verified, the method, and the venues themselves.
}
\note{
Requires the \pkg{jsonlite} package, which is only suggested: the rest
of this package has no dependencies beyond base R.
}
\examples{
if (requireNamespace("jsonlite", quietly = TRUE)) {
  env <- mw_json("chiringuitos")
  env$season
}
}
