package main

import (
	"context"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/atclient"
)

type SyncPDS struct{ client *atclient.APIClient }

func NewSyncPDS(pds, userAgent string) SyncPDS {
	client := atclient.NewAPIClient(pds)
	client.Headers.Set("User-Agent", userAgent)

	return SyncPDS{client: client}
}

func (s SyncPDS) GetRepos(ctx context.Context) []*atproto.SyncListRepos_Repo {
	cursor := ""
	var limit int64 = 50
	fr := true
	var repos []*atproto.SyncListRepos_Repo
	for len(cursor) > 0 || fr {
		fr = false

		ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()

		resp, err := atproto.SyncListRepos(ctx, s.client, cursor, limit)
		if err != nil {
			panic(err)
		}
		repos = append(repos, resp.Repos...)

		if resp.Cursor == nil {
			cursor = ""
		} else {
			cursor = *resp.Cursor
		}
	}
	return repos
}
