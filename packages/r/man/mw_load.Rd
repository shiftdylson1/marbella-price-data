\name{mw_load}
\alias{mw_load}
\title{Load a dataset}
\usage{
mw_load(dataset)
}
\arguments{
\item{dataset}{Dataset name: \code{"sunbed_index"}, \code{"chiringuitos"} or
\code{"beach_clubs"}. Hyphens and capitals are accepted.}
}
\value{
A data frame, one row per venue, with character columns left as
character (\code{stringsAsFactors} is off).
}
\description{
Reads one published dataset into a data frame.
}
\details{
An empty cell means the venue publishes no figure for that column. It is
read as \code{NA}, never as zero: the datasets never guess and never fill a
gap with a placeholder. A zero that is actually present is a real price --
the public beach costs nothing to lie on.
}
\examples{
sunbeds <- mw_load("sunbed_index")
sunbeds[!is.na(sunbeds$price_eur), c("venue", "area", "price_eur")]

# Venues that publish no rate carry a note saying why.
sunbeds[is.na(sunbeds$price_eur), c("venue", "note")]
}
\seealso{
\code{\link{mw_info}} for the provenance a citation needs.
}
