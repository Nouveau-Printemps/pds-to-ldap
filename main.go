package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sPDS := NewSyncPDS(pds, userAgent)
	repos := sPDS.GetRepos(ctx)
	b, _ := json.Marshal(repos)
	println(string(b))
}
