# PDS to LDAP

Simple tool syncing users from an [ATProto PDS](https://atproto.com/guides/glossary#pds-personal-data-server) to a LDAP
server.

## Why?

At [Nouveau Printemps](https://nouveauprintemps.org/), our infrastructure works closely with ATProto.
We want our users to use only one account, but not every service are working with ATProto.
Syncing users from an ATProto PDS to a classical LDAP server enables us to achieve this goal.

## Install

You can build it by yourself with
```bash
# reduces binary size and targets modern CPU (x86_64v3 arch)
GOAMD64=v3 go build -ldflags "-s" .
```
If you have `just` installed, you can simply run
```
just build
```
to get the same result.

If you have a working Go toolchain (i.e. `GOBIN` is in your path), you can install it automatically with
```
go install github.com/Nouveau-Printemps/pds-to-ldap@latest
```

## How?

When you run the command, it will fetch every user from the PDS and it will check if the corresponding `did:plc` has an
entry in the LDAP server.
If there is no account with this `did:plc`, it creates a new one with a random password.

If the user decides to change their handle, they will not loose their account on the LDAP server because the `did:plc`
didn't change.
