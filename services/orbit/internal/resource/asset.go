package resource

import (
	"context"

	"github.com/laxmicoinofficial/go/xdr"
)

func (this *Asset) Populate(ctx context.Context, asset xdr.Asset) error {
	return asset.Extract(&this.Type, &this.Code, &this.Issuer)
}
