package astibranch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	linkDataPayloadKeyCanonicalIdentifier = "$canonical_identifier"
	linkDataPayloadKeyTitle               = "$og_title"
)

type Link struct {
	URL string `json:"url"`
}

type LinkPayload struct {
	BranchKey string    `json:"branch_key"`
	Data      *LinkData `json:"data"`
}

type LinkData struct {
	d map[string]interface{}
	m *sync.Mutex
}

func NewLinkData() *LinkData {
	return &LinkData{
		d: make(map[string]interface{}),
		m: &sync.Mutex{},
	}
}

func (d *LinkData) MarshalJSON() ([]byte, error) {
	d.m.Lock()
	defer d.m.Unlock()
	return json.Marshal(d.d)
}

func (d *LinkData) set(k string, v interface{}) {
	d.m.Lock()
	defer d.m.Unlock()
	d.d[k] = v
}

func (d *LinkData) SetCanonicalIdentifier(v string) *LinkData {
	d.set(linkDataPayloadKeyCanonicalIdentifier, v)
	return d
}

func (d *LinkData) SetTitle(v string) *LinkData {
	d.set(linkDataPayloadKeyTitle, v)
	return d
}

func (d *LinkData) SetCustom(k string, v interface{}) *LinkData {
	d.set(k, v)
	return d
}

// CreateLink creates a link
func (c *Client) CreateLink(d *LinkData) (l Link, err error) {
	// Send
	if err = c.send(http.MethodPost, "/v1/url", LinkPayload{
		BranchKey: c.c.Key,
		Data:      d,
	}, &l); err != nil {
		err = fmt.Errorf("astibranch: sending failed: %w", err)
		return
	}
	return
}
