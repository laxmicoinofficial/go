package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/rover/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("rover_address", govalidator.CustomTypeValidator(isStellarAddress))
}

func isStellarAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}