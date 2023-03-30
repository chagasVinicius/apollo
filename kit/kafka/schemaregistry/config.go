package schemaregistry

import (
	"fmt"
	"time"
)

// Config is used to pass multiple configuration options to the Schema Registry client.
type Config struct {
	URL                            string
	BasicAuthUserInfo              string
	BasicAuthCredentialsSource     string
	SASLMechanism                  string
	SASLUsername                   string
	SASLPassword                   string
	SSLCertificateLocation         string
	SSLKeyLocation                 string
	SSLCaLocation                  string
	SSLDisableEndpointVerification bool
	ConnectionTimeout              time.Duration
	RequestTimeout                 time.Duration
	CacheCapacity                  int
}

// NewConfig returns a new configuration instance with sane defaults.
func NewConfig(url string, options ...Option) Config {
	c := Config{
		URL:                        url,
		BasicAuthCredentialsSource: "URL",
		SASLMechanism:              "GSSAPI",
		ConnectionTimeout:          10000,
		RequestTimeout:             10000,
	}

	for _, option := range options {
		option(&c)
	}

	return c
}

// ====================================================================================================================
// Options

// Option is used to pass configuration options to the Schema registry client.
type Option func(*Config)

// WithURL sets the URL of the Schema Registry client.
func WithURL(url string) Option {
	return func(c *Config) {
		c.URL = url
	}
}

// WithBasictAuth sets the username and password for basic authentication.
// Default source: USER_INFO.
func WithBasicAuth(username, password string) Option {
	return func(c *Config) {
		c.BasicAuthUserInfo = fmt.Sprintf("%s:%s", username, password)
		c.BasicAuthCredentialsSource = "USER_INFO"
	}
}

// WithSASLAuth sets the SASL mechanism.
func WithSASLAuth(mechanism, username, password string) Option {
	return func(c *Config) {
		c.SASLMechanism = mechanism
		c.SASLUsername = username
		c.SASLPassword = password
		c.BasicAuthCredentialsSource = "SASL_INHERIT"
	}
}

// WithSSL sets the SSL configuration.
func WithSSL(certLocation, caLocation, keyLocation string, disableVerify bool) Option {
	return func(c *Config) {
		c.SSLCertificateLocation = certLocation
		c.SSLKeyLocation = keyLocation
		c.SSLCaLocation = caLocation
		c.SSLDisableEndpointVerification = disableVerify
	}
}

// WithConnectionTimeout sets the connection timeout.
func WithConnectionTimeout(d time.Duration) Option {
	return func(c *Config) {
		c.ConnectionTimeout = d
	}
}

// WithRequestTimeout sets the request timeout.
func WithRequestTimeout(d time.Duration) Option {
	return func(c *Config) {
		c.RequestTimeout = d
	}
}
