\name{mw_info}
\alias{mw_info}
\title{Provenance for a dataset}
\usage{
mw_info(dataset)
}
\arguments{
\item{dataset}{Dataset name; see \code{\link{mw_datasets}}.}
}
\value{
A list with the canonical page, the same page in all six published
languages, the row count, the season, the licence and the DOI.
}
\description{
Where the figures were published, when they were verified, how many rows
there are, and the exact credit line the licence asks for.
}
\examples{
info <- mw_info("beach_clubs")
info$canonical
info$attribution
}
