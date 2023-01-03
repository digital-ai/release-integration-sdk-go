package http

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"net/http"
)

type TokenAuthentication struct {
	BearerToken string
}

type BasicAuthentication struct {
	Username string
	Password string
}

type ClientCertificateAuthentication struct {
	CertData []byte
	KeyData  []byte
	CertFile string
	KeyFile  string
}

type CertificateAuthority struct {
	CAData []byte
	CAFile string
}

type HttpClientConfig struct {
	Host                            string
	Insecure                        bool
	BasicAuthentication             *BasicAuthentication
	TokenAuthentication             *TokenAuthentication
	ClientCertificateAuthentication *ClientCertificateAuthentication
	CertificateAuthority            *CertificateAuthority
}

type fetchTokenFunc func(*HttpClientBuilder) (string, error)

type HttpClientBuilder struct {
	config         *rest.Config
	tokenPath      string
	lazyFetchToken fetchTokenFunc
}

func NewHttpClientBuilder() *HttpClientBuilder {
	return &HttpClientBuilder{
		config: &rest.Config{
			TLSClientConfig: rest.TLSClientConfig{},
		},
		lazyFetchToken: voidToken,
	}
}

func voidToken(_ *HttpClientBuilder) (string, error) {
	return "", nil
}

func (b *HttpClientBuilder) WithTokenFetch(tokenPath string) *HttpClientBuilder {
	b.tokenPath = tokenPath
	b.lazyFetchToken = func(builder *HttpClientBuilder) (string, error) {
		if len(builder.config.Username) > 0 && len(builder.config.Password) > 0 {
			values := map[string]string{"username": builder.config.Username, "password": builder.config.Password}
			jsonValue, _ := json.Marshal(values)
			httpClient, err := b.getClient()

			result, err := httpClient.Post(builder.tokenPath, jsonValue)
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
	b.lazyFetchToken = voidToken
	return b
}

func (b *HttpClientBuilder) Build() (*HttpClient, error) {
	if b.lazyFetchToken != nil {
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

	return b.getClient()
}

func (b *HttpClientBuilder) getClient() (*HttpClient, error) {
	client, err := rest.HTTPClientFor(b.config)
	if err != nil {
		return nil, err
	}

	rootCAs, _ := x509.SystemCertPool()
	newTlsConfig := &tls.Config{}
	newTlsConfig.RootCAs = rootCAs

	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = newTlsConfig

	return &HttpClient{
		baseUrl: b.config.Host,
		client:  client,
	}, nil
}
