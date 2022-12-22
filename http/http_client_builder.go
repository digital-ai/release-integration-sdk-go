package http

import (
	"crypto/tls"
	"crypto/x509"
	"k8s.io/client-go/rest"
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

type HttpClientBuilder struct {
	config *rest.Config
}

func NewHttpClientBuilder() *HttpClientBuilder {
	return &HttpClientBuilder{config: &rest.Config{
		TLSClientConfig: rest.TLSClientConfig{},
	}}
}

func (b *HttpClientBuilder) WithHttpClientConfig(config *HttpClientConfig) *HttpClientBuilder {
	b.config.Host = config.Host
	b.config.Insecure = config.Insecure
	if config.CertificateAuthority != nil {
		b.config.TLSClientConfig.CAFile = config.CertificateAuthority.CAFile
		b.config.TLSClientConfig.CAData = config.CertificateAuthority.CAData
	}
	if config.BasicAuthentication != nil {
		b.config.Username = config.BasicAuthentication.Username
		b.config.Password = config.BasicAuthentication.Username
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
	return b
}

func (b *HttpClientBuilder) Build() error {

	httpClient, err := rest.HTTPClientFor(b.config)

	rootCAs, _ := x509.SystemCertPool()
	newTlsConfig := &tls.Config{}
	newTlsConfig.RootCAs = rootCAs

	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = newTlsConfig

	if err != nil {
		return err
	}
	SetClient(httpClient)
	return nil
}
