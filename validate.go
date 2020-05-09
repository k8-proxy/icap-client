package icapclient

import (
	"errors"
	"net/url"
)

func validMethod(method string) (bool, error) {
	if _, registered := registeredMethods[method]; !registered {
		return false, errors.New(ErrMethodNotRegistered)
	}

	return true, nil
}

func validURL(url *url.URL) (bool, error) {

	if url.Scheme != SchemeICAP {
		return false, errors.New(ErrInvalidScheme)
	}

	if url.Host == "" {
		return false, errors.New(ErrInvalidHost)
	}

	return true, nil
}