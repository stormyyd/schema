package schema

import (
	"net/url"
	"strings"
)

type UrlValue struct {
	Key   string
	Value string
}

type UrlValues []UrlValue

func (v UrlValues) Values() map[string][]string {
	m := map[string][]string{}
	for _, p := range v {
		m[p.Key] = append(m[p.Key], p.Value)
	}
	return m
}

func (v UrlValues) Encode() string {
	if len(v) == 0 {
		return ""
	}
	var buf strings.Builder
	for _, p := range v {
		keyEscaped := url.QueryEscape(p.Key)
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(p.Value))
	}
	return buf.String()
}
