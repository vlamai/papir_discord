package message

import (
	"errors"
	"log"
	"mvdan.cc/xurls/v2"
	"net/url"
)

const (
	JiraDomain       = "iv-com.atlassian.net"
	LocalDomain      = "localhost:54727"
	StagingDomain    = "testpapir.skotchapp.com"
	ProductionDomain = "home.skotchapp.com"
	SchemaHttp       = "http"
	SchemaHttps      = "https"
)

var errHostNotMath = errors.New("host isn't match")

type Parser struct {
	urls              []string
	exchangeableLinks ExchangeableLinks
	jiraTracker       TaskLink
}

func New(links ExchangeableLinks, tracker TaskLink) Parser {
	return Parser{
		exchangeableLinks: links,
		jiraTracker:       tracker,
	}
}

func (m *Parser) Parse(s string) []string {
	rxStrict := xurls.Strict()
	urls := rxStrict.FindAllString(s, -1)
	m.urls = urls

	return m.responseForUrls()
}

func (m *Parser) responseForUrls() []string {
	var result result
	for _, u := range m.urls {
		cu, err := url.Parse(u)
		if err != nil {
			log.Printf("parse url | %s | %v\n", u, err)
			continue
		}
		taskName, err := m.jiraTracker.Parse(cu)
		result.Add(err, taskName)

		replacedLnks, err := m.exchangeableLinks.ReplaceHost(cu)
		result.Add(err, replacedLnks...)
	}
	return result
}
