\name{marbellaprices-package}
\alias{marbellaprices-package}
\alias{marbellaprices}
\docType{package}
\title{Marbella Wire Open Price Datasets}
\description{
Three open price datasets for Marbella, on the Spanish Costa del Sol: the
Sunbed Index, a census of chiringuitos (beach restaurants) with menu prices,
and a census of beach clubs with their published day rates.
}
\details{
Use \code{\link{mw_datasets}} to list the datasets, \code{\link{mw_load}} to
read one into a data frame, \code{\link{mw_json}} for the published JSON
envelope, and \code{\link{mw_info}} for the provenance a citation needs.

Every figure is read from a venue's own price list or menu and carries the
source URL and the date it was read. A blank is always "the venue does not
publish this", never a zero and never an estimate.
}
\section{Licence}{
The figures are released under CC BY 4.0: credit "Marbella Wire" and link the
canonical page. \code{\link{mw_info}} returns the exact attribution string.
See the LICENSE-DATA file. The R code in this package is MIT.
}
\references{
Canonical pages: \url{https://marbellawire.com/prices/sunbeds/},
\url{https://marbellawire.com/prices/chiringuitos/},
\url{https://marbellawire.com/prices/beach-clubs/}.

Archived dataset: \doi{10.5281/zenodo.22094846}.
}
\seealso{
Repository: \url{https://github.com/shiftdylson1/marbella-price-data}
}
