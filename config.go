//
//  simplecert
//
//  Created by Philipp Mieden
//  Contact: dreadl0ck@protonmail.ch
//  Copyright © 2018 bestbytes. All rights reserved.
//

package simplecert

import (
	"errors"
	"os"
	"time"
)

var c *Config

// Default contains a default configuration
var Default = &Config{
	// 30 Days before expiration
	RenewBefore: 30 * 24,
	// Once a week
	CheckInterval: 7 * 24 * time.Hour,
	SSLEmail:      "",
	DirectoryURL:  "https://acme-v02.api.letsencrypt.org/directory",
	HTTPAddress:   ":80",
	TLSAddress:    ":443",
	CacheDirPerm:  0700,
	Domains:       []string{},
	CacheDir:      "",
	DNSProvider:   "",
}

// Config allows configuration of simplecert
type Config struct {

	// renew the certificate X hours before it expires
	// LetsEncrypt Certs are valid for 90 Days
	RenewBefore int

	// Interval for checking if cert is closer to expiration than RenewBefore
	CheckInterval time.Duration

	// SSLEmail for contact
	SSLEmail string

	// ACME Directory URL. Can be set to https://acme-staging-v02.api.letsencrypt.org/directory for testing
	DirectoryURL string

	// Endpoints for webroot challenge
	// CAUTION: challenge must be received on port 80 and 443
	// if you choose different ports here you must redirect the traffic
	HTTPAddress string
	TLSAddress  string

	// UNIX Permission for the CacheDir and all files inside
	CacheDirPerm os.FileMode

	// Domains for which to obtain the certificate
	Domains []string

	// Path of the CacheDir
	CacheDir string

	// DNSProvider name for DNS challenges (optional)
	// see: https://godoc.org/github.com/xenolf/lego/providers/dns
	DNSProvider string
}

// CheckConfig checks if config can be used to obtain a cert
func CheckConfig(c *Config) error {
	if len(c.Domains) == 0 {
		return errors.New("simplecert: no domains specified")
	}
	if c.SSLEmail == "" {
		return errors.New("simplecert: no SSLEmail in config")
	}
	if c.DirectoryURL == "" {
		return errors.New("simplecert: no directory url specified")
	}
	return nil
}
