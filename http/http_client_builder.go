package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/Azure/go-ntlmssp"
	"github.com/launchdarkly/go-ntlm-proxy-auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
)

const DefaultContentType = "application/json"

// TokenAuthentication represents the authentication credentials using a bearer token.
type TokenAuthentication struct {
	BearerToken string
}

// BasicAuthentication represents the authentication credentials using a username and password.
type BasicAuthentication struct {
	Username string
	Password string
}

// NtlmAuthentication represents the NTLM authentication credentials using a username, password, and domain.
type NtlmAuthentication struct {
	Username string
	Password string
	Domain   string
}

// OAuth2Authentication represents the OAuth2 authentication credentials using a client ID, client secret, scopes, and token URL.
type OAuth2Authentication struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	TokenURL     string
}

// ClientCertificateAuthentication represents the authentication credentials using a client certificate.
type ClientCertificateAuthentication struct {
	CertData []byte
	KeyData  []byte
	CertFile string
	KeyFile  string
}

// CertificateAuthority represents the certificate authority information.
type CertificateAuthority struct {
	CAData []byte
	CAFile string
}

// HttpClientConfig represents the configuration options for the HttpClient.
type HttpClientConfig struct {
	Host                            string
	Insecure                        bool
	BasicAuthentication             *BasicAuthentication
	TokenAuthentication             *TokenAuthentication
	NtlmAuthentication              *NtlmAuthentication
	OAuth2Authentication            *OAuth2Authentication
	ClientCertificateAuthentication *ClientCertificateAuthentication
	CertificateAuthority            *CertificateAuthority
	ProxyHost                       string
	ProxyPort                       string
	ProxyUsername                   string
	ProxyPassword                   string
	ProxyDomain                     string
}

// fetchTokenFunc represents the function for fetching authentication token.
type fetchTokenFunc func(*HttpClientBuilder) (string, error)

// HttpClientBuilder represents a builder for creating an HttpClient instance.
type HttpClientBuilder struct {
	config         *rest.Config
	headers        map[string][]string
	oAuth2Config   *clientcredentials.Config
	tokenPath      string
	lazyFetchToken fetchTokenFunc
}

// NewHttpClientBuilder creates a new instance of HttpClientBuilder.
func NewHttpClientBuilder() *HttpClientBuilder {
	return &HttpClientBuilder{
		config: &rest.Config{
			TLSClientConfig: rest.TLSClientConfig{},
		},
		lazyFetchToken: voidToken,
	}
}

// voidToken is a function that returns empty token.
func voidToken(_ *HttpClientBuilder) (string, error) {
	return "", nil
}

// WithTokenFetch sets the token path and configures the lazy token fetch function for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithTokenFetch(tokenPath string) *HttpClientBuilder {
	b.tokenPath = tokenPath
	b.lazyFetchToken = func(builder *HttpClientBuilder) (string, error) {
		if len(builder.config.Username) > 0 && len(builder.config.Password) > 0 {
			values := map[string]string{"username": builder.config.Username, "password": builder.config.Password}
			jsonValue, _ := json.Marshal(values)
			httpClient, err := b.buildClient()

			result, err := httpClient.Post(context.Background(), builder.tokenPath, jsonValue)
			if err != nil {
				newErr := fmt.Errorf("cannot fetch token for given credentials: %s", err)
				klog.Errorln(newErr)
				return "", newErr
			}

			var tokenMap = map[string]string{}
			marshalError := json.Unmarshal(result, &tokenMap)
			if marshalError != nil {
				newErr := fmt.Errorf("cannot unmarshal token for given credentials: %s", marshalError)
				klog.Errorln(newErr)
				return "", newErr
			}
			return tokenMap["token"], nil
		} else {
			newErr := fmt.Errorf("missing token and credentials for Argo Connection")
			klog.Errorln(newErr)
			return "", newErr
		}
	}
	return b
}

// ForHost sets the host for the HttpClientBuilder instance.
func (b *HttpClientBuilder) ForHost(host string) *HttpClientBuilder {
	b.config.Host = host
	return b
}

// SetInsecure sets the insecure flag for the HttpClientBuilder instance.
func (b *HttpClientBuilder) SetInsecure(insecure bool) *HttpClientBuilder {
	b.config.Insecure = insecure
	return b
}

// WithCertificateAuthority sets the certificate authority data for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithCertificateAuthority(caData []byte) *HttpClientBuilder {
	b.config.TLSClientConfig.CAData = caData
	return b
}

// WithCertificateAuthorityFile sets the certificate authority file for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithCertificateAuthorityFile(caFile string) *HttpClientBuilder {
	b.config.TLSClientConfig.CAFile = caFile
	return b
}

// WithBasicAuth sets the basic authentication credentials for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithBasicAuth(username string, password string) *HttpClientBuilder {
	b.config.Username = username
	b.config.Password = password
	return b
}

// WithToken sets the bearer token for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithToken(token string) *HttpClientBuilder {
	b.config.BearerToken = token
	return b
}

// WithClientCertAuth sets the client certificate and key data for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithClientCertAuth(certData []byte, keyData []byte) *HttpClientBuilder {
	b.config.TLSClientConfig.CertData = certData
	b.config.TLSClientConfig.KeyData = keyData
	return b
}

// WithClientCertAuthFiles sets the client certificate and key files for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithClientCertAuthFiles(certFile string, keyFile string) *HttpClientBuilder {
	b.config.TLSClientConfig.CertFile = certFile
	b.config.TLSClientConfig.KeyFile = keyFile
	return b
}

// WithNtlmAuth sets the NTLM authentication credentials for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithNtlmAuth(username string, password string, domain string) *HttpClientBuilder {
	b.config.Username = fmt.Sprintf("%s\\%s", domain, username)
	b.config.Password = password
	b.config.Transport = ntlmssp.Negotiator{
		RoundTripper: b.config.Transport,
	}
	return b
}

// WithContentType sets the accept and content type for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithContentType(accept string, contentType string) *HttpClientBuilder {
	b.config.AcceptContentTypes = accept
	b.config.ContentType = contentType
	return b
}

// WithHeaders sets the additional headers for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithHeaders(headers map[string][]string) *HttpClientBuilder {
	b.headers = headers
	return b
}

// WithHttpClientConfig sets the HttpClientConfig for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithHttpClientConfig(config *HttpClientConfig) *HttpClientBuilder {
	b.config.Host = config.Host
	b.config.Insecure = config.Insecure
	if !b.config.Insecure {
		rootCAs, _ := x509.SystemCertPool()
		newTlsConfig := &tls.Config{}
		newTlsConfig.RootCAs = rootCAs

		defaultTransport := http.DefaultTransport.(*http.Transport)
		defaultTransport.TLSClientConfig = newTlsConfig
		b.config.Transport = defaultTransport
	}
	if config.CertificateAuthority != nil {
		b.config.TLSClientConfig.CAFile = config.CertificateAuthority.CAFile
		b.config.TLSClientConfig.CAData = config.CertificateAuthority.CAData
	}
	if config.BasicAuthentication != nil {
		b.config.Username = config.BasicAuthentication.Username
		b.config.Password = config.BasicAuthentication.Password
	}
	if config.TokenAuthentication != nil {
		b.config.BearerToken = config.TokenAuthentication.BearerToken
	}
	if config.ClientCertificateAuthentication != nil {
		b.config.TLSClientConfig.CertData = config.ClientCertificateAuthentication.CertData
		b.config.TLSClientConfig.KeyData = config.ClientCertificateAuthentication.KeyData
		b.config.TLSClientConfig.CertFile = config.ClientCertificateAuthentication.CertFile
		b.config.TLSClientConfig.KeyFile = config.ClientCertificateAuthentication.KeyFile
	}
	if config.NtlmAuthentication != nil {
		b.config.Username = fmt.Sprintf("%s\\%s", config.NtlmAuthentication.Domain, config.NtlmAuthentication.Username)
		b.config.Password = config.NtlmAuthentication.Password
		b.config.Transport = ntlmssp.Negotiator{
			RoundTripper: b.config.Transport,
		}
	}
	if config.OAuth2Authentication != nil {
		b.oAuth2Config = &clientcredentials.Config{
			ClientID:     config.OAuth2Authentication.ClientID,
			ClientSecret: config.OAuth2Authentication.ClientSecret,
			Scopes:       config.OAuth2Authentication.Scopes,
			TokenURL:     config.OAuth2Authentication.TokenURL,
		}
	}
	if config.ProxyHost != "" {
		proxyUrl, proxyErr := url.Parse(fmt.Sprintf("%s:%s", config.ProxyHost, config.ProxyPort))
		proxyUrlMethod := func(*http.Request) (*url.URL, error) {
			if proxyErr != nil {
				return nil, proxyErr
			}
			proxyUrl.User = url.UserPassword(config.ProxyUsername, config.ProxyPassword)
			return proxyUrl, nil
		}

		if config.NtlmAuthentication != nil {
			transport := b.config.Transport.(ntlmssp.Negotiator).RoundTripper.(*http.Transport)
			if config.ProxyDomain != "" {
				ntlmDialContext := ntlm.NewNTLMProxyDialContext(nil, *proxyUrl, config.ProxyUsername, config.ProxyPassword, config.ProxyDomain, nil)
				transport.DialContext = ntlmDialContext
			} else {
				transport.Proxy = proxyUrlMethod
			}
		} else {
			transport := b.config.Transport.(*http.Transport)
			transport.Proxy = proxyUrlMethod
		}
	}
	b.lazyFetchToken = voidToken
	return b
}

// WithOAuth2Config sets the OAuth2 configuration for the HttpClientBuilder instance.
func (b *HttpClientBuilder) WithOAuth2Config(tokenUrl string, clientId string, clientSecret string, scopes []string) *HttpClientBuilder {
	b.oAuth2Config = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		TokenURL:     tokenUrl,
	}
	return b
}

// Build creates and returns an instance of HttpClient based on the current configuration.
// It returns an error if there is any issue during the creation of the client.
func (b *HttpClientBuilder) Build() (*HttpClient, error) {
	if b.oAuth2Config != nil {
		ctx := context.Background()
		httpClient, err := b.buildClient()
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient.client)
		httpClient.client = b.oAuth2Config.Client(ctx)
		return httpClient, nil
	} else if b.lazyFetchToken != nil {
		token, err := b.lazyFetchToken(b)
		if err != nil {
			return nil, err
		}
		if token != "" {
			b.config.BearerToken = token
			b.config.Username = ""
			b.config.Password = ""
		}
	}

	return b.buildClient()
}

// buildClient creates and returns an instance of HttpClient based on the current configuration.
// It returns an error if there is any issue during the creation of the client.
func (b *HttpClientBuilder) buildClient() (*HttpClient, error) {
	client, err := rest.HTTPClientFor(b.config)
	if err != nil {
		return nil, err
	}
	if b.config.AcceptContentTypes == "" {
		b.config.AcceptContentTypes = DefaultContentType
	}
	if b.config.ContentType == "" {
		b.config.ContentType = DefaultContentType
	}
	if b.headers == nil {
		b.headers = make(map[string][]string)
	}
	b.headers["Accept"] = []string{b.config.AcceptContentTypes}
	b.headers["Content-Type"] = []string{b.config.ContentType}
	return &HttpClient{
		baseUrl: b.config.Host,
		client:  client,
		headers: b.headers,
	}, nil
}
