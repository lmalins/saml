package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func (app *application) start() error {
	keyPair, err := tls.LoadX509KeyPair(app.cfg.CertPath, app.cfg.KeyPath)
	if err != nil {
		return err
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return err
	}

	idpMetadataURL, err := url.Parse(app.cfg.IdpURL)
	if err != nil {
		return err
	}

	rootURL, err := url.Parse(app.cfg.RootURL)
	if err != nil {
		return err
	}

	app.infoLog.Printf("Starting server on '%s'", rootURL)
	samlSP, _ := samlsp.New(samlsp.Options{
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadataURL: idpMetadataURL,
		ForceAuthn:     true,
		Logger:         app.errorLog,
	})

	pattern := fmt.Sprintf("%s/static/", app.cfg.AppContext)
	prefix := fmt.Sprintf("%s/static", app.cfg.AppContext)
	fs := http.FileServer(http.Dir("./ui/static"))
	http.Handle(pattern, http.StripPrefix(prefix, fs))

	samlPattern := fmt.Sprintf("%s/saml/", app.cfg.AppContext)
	h := http.HandlerFunc(app.hello)
	http.Handle(app.cfg.AppContext, samlSP.RequireAccount(h))
	http.Handle(samlPattern, samlSP)

	testPattern := fmt.Sprintf("%s/testAuth/", app.cfg.AppContext)
	http.Handle(testPattern, http.HandlerFunc(app.testAuth))
	return http.ListenAndServe(app.cfg.AppPort, nil)
}
