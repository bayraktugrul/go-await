package poll

type Strategy string

const (
	Double Strategy = "DoublePollInterval"
	Fixed  Strategy = "FixedPollInterval"
)
