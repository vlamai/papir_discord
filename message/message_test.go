package message

import (
	"sort"
	"testing"
)

type JiraMock struct{}

func (j JiraMock) GetTaskName(string) (string, error) {
	return "MOCK task name", nil
}

func Test_message_Responses(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  []string
	}{
		{
			name: "empty",
			fields: fields{
				message: "",
			},
			wantR: nil,
		},
		{
			name: "local url",
			fields: fields{
				message: "some text http://localhost:54727/Application/Catalog/InvestProjectManager end of text",
			},
			wantR: []string{
				"https://home.skotchapp.com/Application/Catalog/InvestProjectManager",
				"https://testpapir.skotchapp.com/Application/Catalog/InvestProjectManager",
			},
		},
		{
			name: "production url",
			fields: fields{
				message: "some text https://home.skotchapp.com/Application/Catalog/InvestProjectManager end of text",
			},
			wantR: []string{
				"http://localhost:54727/Application/Catalog/InvestProjectManager",
				"https://testpapir.skotchapp.com/Application/Catalog/InvestProjectManager",
			},
		},
		{
			name: "staging url",
			fields: fields{
				message: "some text https://testpapir.skotchapp.com/Application/Catalog/InvestProjectManager end of text",
			},
			wantR: []string{
				"http://localhost:54727/Application/Catalog/InvestProjectManager",
				"https://home.skotchapp.com/Application/Catalog/InvestProjectManager",
			},
		},
		{
			name: "jira url",
			fields: fields{
				message: "some text https://iv-com.atlassian.net/browse/APP-2286 end of text",
			},
			wantR: []string{
				"MOCK task name",
			},
		},
		{
			name: "random url",
			fields: fields{
				message: "some text https://google.com end of text",
			},
			wantR: nil,
		},
	}

	exchangeableLinks := ExchangeableLinks{
		Links: []Link{
			{
				Schema: SchemaHttp,
				Host:   LocalDomain,
			},
			{
				Schema: SchemaHttps,
				Host:   ProductionDomain,
			},
			{
				Schema: SchemaHttps,
				Host:   StagingDomain,
			},
		},
	}
	jiraTracker := TaskLink{
		Host:    JiraDomain,
		Tracker: JiraMock{},
	}
	m := New(exchangeableLinks, jiraTracker)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := m.Parse(tt.fields.message); !SortCompare(gotR, tt.wantR) {
				t.Errorf("Responses() = \n%v\nwant \n%v\n", gotR, tt.wantR)
			}
		})
	}
}

func SortCompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
