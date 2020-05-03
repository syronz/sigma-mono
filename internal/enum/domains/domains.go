package domains

// Domain is used for type of event
type Domain string

// Domain enums
const (
	Accounting     Domain = "accounting"
	Administration Domain = "administration"
	Central        Domain = "central"
	Deal           Domain = "deal"
	Facture        Domain = "facture"
	Product        Domain = "product"
	Stockpile      Domain = "stockpile"
	Sync           Domain = "sync"
)
