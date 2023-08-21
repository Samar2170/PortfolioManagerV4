package security

type Security interface {
	GetLatestPrice() (float64, error)
}
