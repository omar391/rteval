package routing

import (
	"context"
	"strings"

	"github.com/omar391/rteval/pkg/event"
)

type Action func(ctx context.Context, e event.Event[any]) error

type Rule struct {
	TypePrefix string
	Action     Action
}

type Evaluator struct {
	rules []Rule
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (ev *Evaluator) AddRule(typePrefix string, action Action) {
	ev.rules = append(ev.rules, Rule{TypePrefix: typePrefix, Action: action})
}

func (ev *Evaluator) Process(ctx context.Context, e event.Event[any]) error {
	for _, rule := range ev.rules {
		if strings.HasPrefix(e.Type(), rule.TypePrefix) {
			if err := rule.Action(ctx, e); err != nil {
				return err
			}
		}
	}
	return nil
}
