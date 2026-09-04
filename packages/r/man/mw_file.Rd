\name{mw_file}
\alias{mw_file}
\title{Path to a packaged data file}
\usage{
mw_file(dataset, ext = c("csv", "json"))
}
\arguments{
\item{dataset}{Dataset name; see \code{\link{mw_datasets}}.}

\item{ext}{Either \code{"csv"} or \code{"json"}.}
}
\value{
A single file path.
}
\description{
The location on disk of one of the installed CSV or JSON files, for reading
with a tool of your own.
}
\examples{
mw_file("sunbed_index", "csv")
}
