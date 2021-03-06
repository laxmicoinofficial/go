package orbit

import (
	"encoding/json"

	"github.com/laxmicoinofficial/go/support/errors"
	"github.com/laxmicoinofficial/go/support/render/problem"
	"github.com/laxmicoinofficial/go/xdr"
)

func (herr Error) Error() string {
	return `Orbit error: "` + herr.Problem.Title + `". Check orbit.Error.Problem for more information.`
}

// ToProblem converts the Prolem to a problem.P
func (prob Problem) ToProblem() problem.P {
	extras := make(map[string]interface{})
	for k, v := range prob.Extras {
		extras[k] = v
	}

	return problem.P{
		Type:     prob.Type,
		Title:    prob.Title,
		Status:   prob.Status,
		Detail:   prob.Detail,
		Instance: prob.Instance,
		Extras:   extras,
	}
}

// Envelope extracts the transaction envelope that triggered this error from the
// extra fields.
func (herr *Error) Envelope() (*xdr.TransactionEnvelope, error) {
	raw, ok := herr.Problem.Extras["envelope_xdr"]
	if !ok {
		return nil, ErrEnvelopeNotPopulated
	}

	var b64 string
	var result xdr.TransactionEnvelope

	err := json.Unmarshal(raw, &b64)
	if err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	err = xdr.SafeUnmarshalBase64(b64, &result)
	if err != nil {
		return nil, errors.Wrap(err, "xdr decode failed")
	}

	return &result, nil
}

// ResultCodes extracts a result code summary from the error, if possible.
func (herr *Error) ResultCodes() (*TransactionResultCodes, error) {

	raw, ok := herr.Problem.Extras["result_codes"]
	if !ok {
		return nil, ErrResultCodesNotPopulated
	}

	var result TransactionResultCodes
	err := json.Unmarshal(raw, &result)
	if err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	return &result, nil
}
