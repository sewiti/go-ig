package ig

import (
	"context"
	"fmt"
)

func ExampleGet() {
	const username = "instagram"
	profile, posts, err := Get(context.Background(), username)
	if err != nil {
		panic(err)
	}
	fmt.Printf("extracted %d %s's posts", len(posts), profile.Identifier.Value)
	// Output:
	// extracted 9 instagram's posts
}
