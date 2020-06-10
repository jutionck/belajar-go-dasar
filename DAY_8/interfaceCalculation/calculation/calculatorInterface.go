package calculation

type Calculator interface {
	GetCalculation() (int, error)
}
