package http

import (
	"crypto/tls"
	"crypto/x509"
	"k8s.io/client-go/rest"
	"log"
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

	s := rootCAs.Subjects()
	println(len(s))

	const rootPEM = `-----BEGIN CERTIFICATE-----
MIIBdjCCAR2gAwIBAgIBADAKBggqhkjOPQQDAjAjMSEwHwYDVQQDDBhrM3Mtc2Vy
dmVyLWNhQDE2NzE1NDQ3MTgwHhcNMjIxMjIwMTM1ODM4WhcNMzIxMjE3MTM1ODM4
WjAjMSEwHwYDVQQDDBhrM3Mtc2VydmVyLWNhQDE2NzE1NDQ3MTgwWTATBgcqhkjO
PQIBBggqhkjOPQMBBwNCAATcogzPvK2wunLn+UuKwllwiDYhnvct0ZKTJm7yEJju
3BSmCWWutuB/M215k76ncF3uNs2r+gZRx2U8HhlzFrneo0IwQDAOBgNVHQ8BAf8E
BAMCAqQwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU/Ro/dyLM21UqBH3NWQfq
PtsulPwwCgYIKoZIzj0EAwIDRwAwRAIga6Ub/9W9H0nH9wypqNQYTR0WVUHZUZCn
Z8chj2e6SuMCIGF91TKKL2sbgvP7bMr+lA4QdC+wHU9/zT6vIZIEHLQo
-----END CERTIFICATE-----`

	ok := rootCAs.AppendCertsFromPEM([]byte(rootPEM))
	// Append our cert to the system pool
	if !ok {
		log.Println("No certs appended, using system certs only")
	}

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
