package stellar

import (
	"sync"

	"github.com/laxmicoinofficial/go/clients/orbit"
	"github.com/laxmicoinofficial/go/support/log"
)

// AccountConfigurator is responsible for configuring new Rover accounts that
// participate in ICO.
type AccountConfigurator struct {
	Orbit           orbit.ClientInterface `inject:""`
	NetworkPassphrase string
	IssuerPublicKey   string
	SignerSecretKey   string
	NeedsAuthorize    bool
	TokenAssetCode    string
	StartingBalance   string
	OnAccountCreated  func(destination string)
	OnAccountCredited func(destination string, assetCode string, amount string)

	signerPublicKey      string
	sequence             uint64
	sequenceMutex        sync.Mutex
	processingCount      int
	processingCountMutex sync.Mutex
	log                  *log.Entry
}
