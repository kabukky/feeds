
// +build gofuzz

package feeds

import (
	"bytes"
	"time"
)

func Fuzz(input []byte) int {
	s := string(input)
	now, err := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	if err != nil {
		panic(err)
	}
	tz := time.FixedZone("EST", -5*60*60)
	now = now.In(tz)

	feed := &Feed{
		Title: s,
		Link: &Link{Href: s},
		Description: s,
		Author: &Author{Name: s, Email: s},
		Created: now,
		Copyright: s,
	}
	
	feed.Items = []*Item{
		{
			Title:       s,
			Link:        &Link{Href: s},
			Description: s,
			Author:      &Author{Name: s, Email: s},
			Created:     now,
		},
		{
			Title:       s,
			Link:        &Link{Href: s},
			Description: s,
			Created:     now,
		}}

	_, err = feed.ToAtom()
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := feed.WriteAtom(&buf); err != nil {
		panic(err)
	}

	_, err = feed.ToRss()
	if err != nil {
		panic(err)
	}
	buf.Reset()
	if err := feed.WriteRss(&buf); err != nil {
		panic(err)
	}
	return 1
}
