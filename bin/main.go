package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/sewiti/go-ig"
	"golang.org/x/net/html"
)

func main() {
	var username string
	var pretty, raw bool
	flag.StringVar(&username, "username", "", "Instagram username.")
	flag.BoolVar(&pretty, "pretty", false, "Pretty JSON indentation.")
	flag.BoolVar(&raw, "raw", false, "Print raw data.")
	flag.Parse()

	if username == "" {
		fmt.Fprintln(os.Stderr, "username is required")
		os.Exit(1)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if raw {
		if err := doRaw(ctx, username); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		return
	}

	if err := doParsed(ctx, username, pretty); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}
}

func doParsed(ctx context.Context, username string, pretty bool) error {
	profile, posts, err := ig.Get(ctx, username)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	if pretty {
		enc.SetIndent("", "    ")
	}

	return enc.Encode(struct {
		Profile *ig.Profile `json:"profile"`
		Posts   []ig.Post   `json:"posts"`
	}{profile, posts})
}

func doRaw(ctx context.Context, username string) error {
	url := "https://www.instagram.com/" + url.PathEscape(username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("get %s: %s", url, res.Status)
	}

	root, err := html.Parse(res.Body)
	if err != nil {
		return err
	}

	for _, v := range ig.ExtractData(root) {
		fmt.Println(v)
	}
	return nil
}
