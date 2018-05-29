package main

import (
	"net/http"

	"github.com/laxmicoinofficial/go/clients/orbit"
	"github.com/laxmicoinofficial/go/services/dakibot/internal"
	"github.com/laxmicoinofficial/go/strkey"
)

func initDakibot(dakibotSecret string, networkPassphrase string, horizonURL string, startingBalance string) *internal.Bot {
	if dakibotSecret == "" || networkPassphrase == "" || horizonURL == "" || startingBalance == "" {
		return nil
	}

	// ensure its a seed if its not blank
	strkey.MustDecode(strkey.VersionByteSeed, dakibotSecret)

	return &internal.Bot{
		Secret: dakibotSecret,
		Orbit: &orbit.Client{
			URL:  horizonURL,
			HTTP: http.DefaultClient,
		},
		Network:         networkPassphrase,
		StartingBalance: startingBalance,
	}
}
