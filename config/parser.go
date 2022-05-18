package config

import "github.com/vlamai/papir_discord/message"

func NewParser(t message.TaskTracker) message.Parser {
	exchangeableLinks := message.ExchangeableLinks{
		Links: []message.Link{
			{
				Schema: message.SchemaHttp,
				Host:   message.LocalDomain,
			},
			{
				Schema: message.SchemaHttps,
				Host:   message.ProductionDomain,
			},
			{
				Schema: message.SchemaHttps,
				Host:   message.StagingDomain,
			},
		},
	}
	jiraTracker := message.TaskLink{
		Host:    message.JiraDomain,
		Tracker: t,
	}
	m := message.New(exchangeableLinks, jiraTracker)
	return m
}
