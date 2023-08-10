package ig

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestParseHTML(t *testing.T) {
	const testfile = "testdata/golang.go.html"
	tzMinus8 := time.FixedZone("", -8*3600) // -08:00

	expectedProfile := &Profile{
		Context: "https://schema.org",
		Type:    "ProfilePage",
		Author: User{
			Type: "Organization",
			Identifier: &Identifier{
				Type:       "http://schema.org/PropertyValue",
				PropertyID: "Username",
				Value:      "golang.go",
			},
			Image:         "https://scontent.cdninstagram.com/v/t51.2885-19/79737449_589411391617714_4708578689623785472_n.jpg?stp=dst-jpg_s100x100&_nc_cat=100&ccb=1-7&_nc_sid=c4dd86&_nc_ohc=mMluYlxoLwwAX8Gbj4V&_nc_ht=scontent.cdninstagram.com&oh=00_AfAGSMDLl_qWISdgdHgXPzyGH5wGRzofOGmkwzmKTTwzzg&oe=64FCA401",
			Name:          "golang.go",
			AlternateName: "golang.go",
			SameAs:        "http://www.golang.org/",
			URL:           "https://www.instagram.com/golang.go",
		},
		MainEntityOfPage: Entity{
			Type: "ProfilePage",
			ID:   "https://www.instagram.com/golang.go/",
		},
		Identifier: Identifier{
			Type:       "http://schema.org/PropertyValue",
			PropertyID: "Username",
			Value:      "golang.go",
		},
		InteractionStatistic: []InteractionStatisticStr{
			{
				Type:                 "InteractionCounter",
				InteractionType:      "http://schema.org/WriteAction",
				UserInteractionCount: "127",
			},
			{
				Type:                 "InteractionCounter",
				InteractionType:      "http://schema.org/FollowAction",
				UserInteractionCount: "9329",
			},
		},
	}

	expectedPosts := []Post{
		{
			ArticleBody: "Function is group of statement to perform task. \n#function #go #golang #programming #task #code #coder #codding #programinglanguage #programmer #computersciene #informationtechnology #parameter #return #variable #comment",
			Author: User{
				Type:          "Organization",
				Image:         "https://scontent.cdninstagram.com/v/t51.2885-19/79737449_589411391617714_4708578689623785472_n.jpg?stp=dst-jpg_s100x100&_nc_cat=100&ccb=1-7&_nc_sid=c4dd86&_nc_ohc=mMluYlxoLwwAX8Gbj4V&_nc_ht=scontent.cdninstagram.com&oh=00_AfAGSMDLl_qWISdgdHgXPzyGH5wGRzofOGmkwzmKTTwzzg&oe=64FCA401",
				Name:          "golang.go",
				AlternateName: "golang.go",
				SameAs:        "http://www.golang.org/",
				URL:           "https://www.instagram.com/golang.go",
				InteractionStatistic: &InteractionStatistic{
					Type:                 "InteractionCounter",
					InteractionType:      "http://schema.org/FollowAction",
					UserInteractionCount: 9329,
				},
			},
			Comment: Comment{
				Type: "Comment",
				Text: "This is great! 🙌",
				Author: User{
					Type: "Organization",
					Identifier: &Identifier{
						Type:       "http://schema.org/PropertyValue",
						PropertyID: "Username",
						Value:      "pragmatic_reviews",
					},
					Image:         "https://scontent.cdninstagram.com/v/t51.2885-19/61800585_406425573315220_4029688829941121024_n.jpg?stp=dst-jpg_s100x100&_nc_cat=108&ccb=1-7&_nc_sid=c4dd86&_nc_ohc=mIemlITpHWkAX8hb5Zc&_nc_ht=scontent.cdninstagram.com&oh=00_AfDKEYPWAocpRyyQi64Ht-mKSDC6BFslXc-mgd3pavZHYQ&oe=64FCAC66",
					Name:          "Pragmatic Reviews",
					AlternateName: "pragmatic_reviews",
					SameAs:        "https://pragmaticreviews.com/nextjs",
					URL:           "https://www.instagram.com/pragmatic_reviews",
				},
				DateCreated: time.Date(2020, 2, 15, 10, 56, 4, 0, tzMinus8),
				InteractionStatistic: &InteractionStatistic{
					Type:                 "InteractionCounter",
					InteractionType:      "InteractionCounter",
					UserInteractionCount: 1,
				},
			},
			CommentCount: "1",
			Context:      "https://schema.org",
			DateCreated:  time.Date(2020, 2, 14, 11, 25, 13, 0, tzMinus8),
			DateModified: time.Date(2020, 2, 14, 11, 25, 15, 0, tzMinus8),
			Headline:     "Function is group of statement to perform task. \n#function #go #golang #programming #task #code #coder #codding #programinglanguage #programmer #computersciene #informationtechnology #parameter #return #variable #comment",
			Identifier: Identifier{
				Type:       "http://schema.org/PropertyValue",
				PropertyID: "Post Shortcode",
				Value:      "B8jxbQSAdgv",
			},
			Image: []Image{
				{
					Type:                 "https://schema.org/ImageObject",
					RepresentativeOfPage: "True",
					Height:               "706",
					Width:                "710",
					URL:                  "https://scontent.cdninstagram.com/v/t51.2885-15/83926938_842803719506589_7339813218322795275_n.jpg?stp=dst-jpg_s640x640&_nc_cat=111&ccb=1-7&_nc_sid=c4dd86&_nc_ohc=HDOBhQX4MiEAX9fdLlV&_nc_ht=scontent.cdninstagram.com&oh=00_AfBVaJ3CIK3D1td4BacKV3r5Wqi0gz90BhhGimrf0Vs8mA&oe=64DA30BA",
				},
				{
					Type:                 "https://schema.org/ImageObject",
					RepresentativeOfPage: "True",
					Height:               "667",
					Width:                "671",
					URL:                  "https://scontent.cdninstagram.com/v/t51.2885-15/84502248_103513134494194_4684242200608776001_n.jpg?stp=dst-jpg_s640x640&_nc_cat=107&ccb=1-7&_nc_sid=c4dd86&_nc_ohc=XPFTapmEH9oAX-aFFSo&_nc_ht=scontent.cdninstagram.com&oh=00_AfCKVwhUbfLUl0iZMwLJjgGqvq8VBwQ_0C9rG_C5tAzrsg&oe=64D9F0D3",
				},
			},
			InteractionStatistic: []InteractionStatistic{
				{
					Type:                 "InteractionCounter",
					InteractionType:      "http://schema.org/LikeAction",
					UserInteractionCount: 118,
				},
				{
					Type:                 "InteractionCounter",
					InteractionType:      "https://schema.org/CommentAction",
					UserInteractionCount: 1,
				},
			},
			MainEntityOfPage: Entity{
				Type: "ItemPage",
				ID:   "https://www.instagram.com/golang.go/",
			},
			Type:  "SocialMediaPosting",
			URL:   "https://www.instagram.com/p/B8jxbQSAdgv/",
			Video: []Video{},
		},
	}

	f, err := os.Open(testfile)
	if err != nil {
		t.Errorf("os.Open(%q) error = %v", testfile, err)
		return
	}
	defer f.Close()

	profile, posts, err := ParseHTML(f)
	if err != nil {
		t.Errorf(`ParseHTML() error = %v`, err)
		return
	}

	if !reflect.DeepEqual(profile, expectedProfile) {
		t.Errorf("ParseHTML()\nprofile:\n%v\n\nexpectedProfile:\n%v", profile, expectedProfile)
	}
	if !reflect.DeepEqual(posts, expectedPosts) {
		t.Errorf("ParseHTML()\nposts:\n%v\n\nexpectedPosts:\n%v", posts, expectedPosts)
	}
}
