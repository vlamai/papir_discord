package message

import (
	"fmt"
	"net/url"
)

type Link struct {
	Schema string
	Host   string
}

type ExchangeableLinks struct {
	Links []Link
}

func (el ExchangeableLinks) ReplaceHost(cUrl *url.URL) ([]string, error) {
	foundIndex, err := el.hostConations(cUrl)
	if err != nil {
		return nil, err
	}

	var result []string
	for i, link := range el.Links {
		if i == foundIndex {
			continue
		}
		l := fmt.Sprintf("%s://%s%s", link.Schema, link.Host, cUrl.RequestURI())
		result = append(result, l)
	}
	return result, nil
}

func (el ExchangeableLinks) hostConations(cUrl *url.URL) (int, error) {
	for i, link := range el.Links {
		if link.Host == cUrl.Host {
			return i, nil
		}
	}
	return 0, errHostNotMath
}
