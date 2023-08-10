package ig

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Profile struct {
	Context              string                 `json:"@context"`
	Type                 string                 `json:"@type"`
	Description          string                 `json:"description"`
	Author               User                   `json:"author"`
	MainEntityOfPage     Entity                 `json:"mainEntityOfPage"`
	Identifier           Identifier             `json:"identifier"`
	InteractionStatistic []InteractionStatistic `json:"interationStatistic"`
}

type User struct {
	Type                 string                `json:"@type"`
	Identifier           *Identifier           `json:"identifier"`
	Image                string                `json:"image"`
	Name                 string                `json:"name"`
	AlternateName        string                `json:"alternateName"`
	SameAs               string                `json:"sameAs"`
	URL                  string                `json:"url"`
	InteractionStatistic *InteractionStatistic `json:"interactionStatistic"`
}

type Identifier struct {
	Type       string `json:"@type"`
	PropertyID string `json:"propertyID"`
	Value      string `json:"value"`
}

type Entity struct {
	Type string `json:"@type"`
	ID   string `json:"@id"`
}

type InteractionStatistic struct {
	Type                 string `json:"@type"`
	InteractionType      string `json:"interactionType"`
	UserInteractionCount int    `json:"userInteractionCount"`
}

type Post struct {
	ArticleBody          string                 `json:"articleBody"`
	Author               User                   `json:"author"`
	Comment              Comment                `json:"comment"`
	CommentCount         string                 `json:"commentCount"`
	ContentLocation      Location               `json:"contentLocation"`
	Context              string                 `json:"@context"`
	DateCreated          time.Time              `json:"dateCreated"`
	DateModified         time.Time              `json:"dateModified"`
	Headline             string                 `json:"headline"`
	Identifier           Identifier             `json:"identifier"`
	Image                []Image                `json:"image"`
	InteractionStatistic []InteractionStatistic `json:"interationStatistic"`
	MainEntityOfPage     Entity                 `json:"mainEntityOfPage"`
	Type                 string                 `json:"@type"`
	URL                  string                 `json:"url"`
	Video                []Video                `json:"video"`
}

type Comment struct {
	Type                 string                `json:"@type"`
	Text                 string                `json:"text"`
	Author               User                  `json:"author"`
	DateCreated          time.Time             `json:"dateCreated"`
	InteractionStatistic *InteractionStatistic `json:"interationStatistic"`
}

type Location struct {
	Type             string `json:"@type"`
	Name             string `json:"name"`
	MainEntityOfPage Entity `json:"mainEntityOfPage"`
}

type Image struct {
	Type                 string `json:"@type"`
	Caption              string `json:"caption"`
	RepresentativeOfPage string `json:"representativeOfPage"`
	Height               string `json:"height"`
	Width                string `json:"width"`
	URL                  string `json:"url"`
}

type Video struct {
	Type                 string                 `json:"@type"`
	UploadDate           time.Time              `json:"uploadDate"`
	Description          string                 `json:"description"`
	Name                 string                 `json:"name"`
	Caption              string                 `json:"caption"`
	Height               string                 `json:"height"`
	Width                string                 `json:"width"`
	ContentURL           string                 `json:"contentUrl"`
	ThumbnailURL         string                 `json:"thumbnailUrl"`
	Genre                []json.RawMessage      `json:"genre"`    // TODO
	Keywords             []json.RawMessage      `json:"keywords"` // TODO
	InteractionStatistic []InteractionStatistic `json:"interationStatistic"`
	Creator              User                   `json:"creator"`
	Comment              []Comment              `json:"comment"`
	CommentCount         string                 `json:"commentCount"`
	IsFamilyFriendly     bool                   `json:"isFamilyFriendly"`
	Mentions             []User                 `json:"mentions"`
	InLanguage           string                 `json:"inLanguage"`
	Duration             string                 `json:"duration"`
	EmbeddedTextCaption  string                 `json:"embeddedTextCaption"`
	Transcript           string                 `json:"transcript"`
}

func Get(ctx context.Context, username string) (*Profile, []Post, error) {
	url := "https://www.instagram.com/" + url.PathEscape(username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, nil, fmt.Errorf("get %s: %s", url, res.Status)
	}

	return ParseHTML(res.Body)
}

func ParseHTML(r io.Reader) (*Profile, []Post, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, nil, err
	}
	data := ExtractData(root)

	profile, i, err := parseProfile(data)
	if err != nil {
		return nil, nil, err
	}
	if i >= 0 {
		data = append(data[:i], data[i+1:]...)
	}

	posts, _, err := parsePosts(data)
	if err != nil {
		return nil, nil, err
	}

	return profile, posts, nil
}

func ExtractData(root *html.Node) []string {
	var data []string
	walkHTML(root, func(n *html.Node) (continue_ bool) {
		if n.Parent == nil || n.Parent.DataAtom != atom.Script || n.Type != html.TextNode {
			return true
		}
		if getAttr(n.Parent, "type") != "application/ld+json" {
			return true
		}
		data = append(data, n.Data)
		return true
	})
	return data
}

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func walkHTML(root *html.Node, fn func(*html.Node) (continue_ bool)) {
	queue := []*html.Node{root}
	pop := func() *html.Node {
		if len(queue) == 0 {
			return nil
		}
		v := queue[0]
		queue = queue[1:]
		return v
	}

	for n := pop(); n != nil; n = pop() {
		if !fn(n) {
			return
		}
		if n.FirstChild != nil {
			queue = append(queue, n.FirstChild)
		}
		if n.NextSibling != nil {
			queue = append(queue, n.NextSibling)
		}
	}
}

func parseProfile(data []string) (p *Profile, dataIndex int, err error) {
	var first error
	for i, v := range data {
		var p Profile
		if err := json.Unmarshal([]byte(v), &p); err != nil {
			if first == nil {
				first = err
			}
			continue
		}
		return &p, i, nil
	}
	return nil, -1, first
}

func parsePosts(data []string) (posts []Post, dataIndex int, err error) {
	var first error
	for i, v := range data {
		var posts []Post
		if err := json.Unmarshal([]byte(v), &posts); err != nil {
			if first == nil {
				first = err
			}
			continue
		}
		return posts, i, nil
	}
	return nil, -1, first
}
