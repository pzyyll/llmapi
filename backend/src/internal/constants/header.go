package constants

type HeaderType string

const (
	TurnstileTokenHeader HeaderType = "X-Turnstile-Token"
	AuthorizationHeader  HeaderType = "Authorization"
)
