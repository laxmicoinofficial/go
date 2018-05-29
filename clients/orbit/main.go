// Package orbit provides client access to a orbit server, allowing an
// application to post transactions and lookup ledger information.
//
// Create an instance of `Client` to customize the server used, or alternatively
// use `DefaultTestNetClient` or `DefaultPublicNetClient` to access the SDF run
// orbit servers.
package orbit

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	"github.com/laxmicoinofficial/go/build"
	"github.com/laxmicoinofficial/go/support/errors"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{
	URL:  "https://orbit-testnet.rover.network",
	HTTP: http.DefaultClient,
}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{
	URL:  "https://orbit.rover.network",
	HTTP: http.DefaultClient,
}

// At is a paging parameter that can be used to override the URL loaded in a
// remote method call to orbit.
type At string

// Cursor represents `cursor` param in queries
type Cursor string

// Limit represents `limit` param in queries
type Limit uint

// Order represents `order` param in queries
type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

var (
	// ErrResultCodesNotPopulated is the error returned from a call to
	// ResultCodes() against a `Problem` value that doesn't have the
	// "result_codes" extra field populated when it is expected to be.
	ErrResultCodesNotPopulated = errors.New("result_codes not populated")

	// ErrEnvelopeNotPopulated is the error returned from a call to
	// Envelope() against a `Problem` value that doesn't have the
	// "envelope_xdr" extra field populated when it is expected to be.
	ErrEnvelopeNotPopulated = errors.New("envelope_xdr not populated")
)

// Client struct contains data required to connect to Orbit instance
type Client struct {
	// URL of Orbit server to connect
	URL string

	// HTTP client to make requests with
	HTTP HTTP

	fixURLOnce sync.Once
}

type ClientInterface interface {
	Root() (Root, error)
	HomeDomainForAccount(aid string) (string, error)
	LoadAccount(accountID string) (Account, error)
	LoadAccountOffers(accountID string, params ...interface{}) (offers OffersPage, err error)
	LoadMemo(p *Payment) error
	LoadOrderBook(selling Asset, buying Asset, params ...interface{}) (orderBook OrderBookSummary, err error)
	StreamLedgers(ctx context.Context, cursor *Cursor, handler LedgerHandler) error
	StreamPayments(ctx context.Context, accountID string, cursor *Cursor, handler PaymentHandler) error
	StreamTransactions(ctx context.Context, accountID string, cursor *Cursor, handler TransactionHandler) error
	SubmitTransaction(txeBase64 string) (TransactionSuccess, error)
}

// Error struct contains the problem returned by Orbit
type Error struct {
	Response *http.Response
	Problem  Problem
}

// HTTP represents the HTTP client that a orbit client uses to communicate
type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// LedgerHandler is a function that is called when a new ledger is received
type LedgerHandler func(Ledger)

// PaymentHandler is a function that is called when a new payment is received
type PaymentHandler func(Payment)

// TransactionHandler is a function that is called when a new transaction is received
type TransactionHandler func(Transaction)

// ensure that the orbit client can be used as a SequenceProvider
var _ build.SequenceProvider = &Client{}

// ensure that the orbit client implements ClientInterface
var _ ClientInterface = &Client{}
