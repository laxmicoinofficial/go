package orbit

import (
	hProblem "github.com/laxmicoinofficial/go/services/orbit/internal/render/problem"
	"github.com/laxmicoinofficial/go/support/render/problem"
)

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() {
	problem.Render(action.Ctx, action.W, hProblem.NotImplemented)
}
