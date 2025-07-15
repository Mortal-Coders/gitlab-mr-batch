package httpclient

type TokenType string

const (
	Bearer TokenType = "Bearer"
	Basic  TokenType = "Basic"
	APIKey TokenType = "ApiKey"
	Custom TokenType = "Custom"
)

type TokenConfig struct {
	Type       TokenType
	Value      string
	HeaderName string // for customHeader
}
