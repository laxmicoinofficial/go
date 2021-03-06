---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the rover ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/laxmicoinofficial/go/keypair) (used for creating stellar-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/laxmicoinofficial/go).

## Client Packages

The Go SDK contains packages for interacting with the various rover services:

- [`orbit`](https://godoc.org/github.com/laxmicoinofficial/go/clients/orbit) provides client access to a orbit server, allowing you to load account information, stream payments, post transactions and more.
- [`stellartoml`](https://godoc.org/github.com/laxmicoinofficial/go/clients/stellartoml) provides the ability to resolve Rover.toml files from the internet.  You can read about [Rover.toml concepts here](../../guides/concepts/stellar-toml.md).
- [`federation`](https://godoc.org/github.com/laxmicoinofficial/go/clients/federation) makes it easy to resolve a rover addresses (e.g. `scott*rover.network`) into a rover account ID suitable for use within a transaction.

