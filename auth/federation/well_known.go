package federation

const (
	authEP  = "/oauth2/authorize"
	tokenEP = "/oauth2/token" //nolint:gosec // false positive, it's not credentials
)
