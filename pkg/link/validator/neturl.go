package validator

import "net/url"

type NetUrlValidator struct{}

var _ Validator = (*NetUrlValidator)(nil)

func NewNetUrlValidator() *NetUrlValidator {
	return &NetUrlValidator{}
}

func (v *NetUrlValidator) Validate(link string) bool {
	if link == "" {
		return false
	}

	parsedLink, err := url.Parse(link)

	if err != nil {
		return false
	}

	if parsedLink.Scheme != "http" && parsedLink.Scheme != "https" {
		return false
	}

	if parsedLink.Host == "" {
		return false
	}

	return true
}
