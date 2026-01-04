package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/atclient"
)

var (
	pds       = "localhost"
	userAgent = "pds-to-ldap/0.0 (https://github.com/Nouveau-Printemps/pds-to-ldap)"
	timeout   = 30
)

func init() {
	if v := os.Getenv("PDS_LDAP__TARGET_PDS"); len(v) > 0 {
		pds = v
	}
	flag.StringVar(&pds, "pds", pds, "link to the PDS")
	if v := os.Getenv("PDS_LDAP__USER_AGENT"); len(v) > 0 {
		userAgent = v
	}
	flag.StringVar(&userAgent, "user-agent", userAgent, "user-agent to use")

	flag.IntVar(&timeout, "timeout", timeout, "time before timeout (in seconds)")
}

func main() {
	flag.Parse()
	client := atclient.NewAPIClient(pds)
	client.Headers.Set("User-Agent", userAgent)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor := ""
	var limit int64 = 50
	fr := true
	for len(cursor) > 0 || fr {
		fr = false

		ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()

		resp, err := atproto.SyncListRepos(ctx, client, cursor, limit)
		if err != nil {
			panic(err)
		}
		for _, r := range resp.Repos {
			b, err := json.Marshal(r)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		}

		if resp.Cursor == nil {
			cursor = ""
		} else {
			cursor = *resp.Cursor
		}
	}
}
