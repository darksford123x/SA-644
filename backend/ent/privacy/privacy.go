// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/darksford123x/repairinvoice-record/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The DeviceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeviceQueryRuleFunc func(context.Context, *ent.DeviceQuery) error

// EvalQuery return f(ctx, q).
func (f DeviceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeviceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeviceQuery", q)
}

// The DeviceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeviceMutationRuleFunc func(context.Context, *ent.DeviceMutation) error

// EvalMutation calls f(ctx, m).
func (f DeviceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeviceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeviceMutation", m)
}

// The RepairInvoiceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RepairInvoiceQueryRuleFunc func(context.Context, *ent.RepairInvoiceQuery) error

// EvalQuery return f(ctx, q).
func (f RepairInvoiceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RepairInvoiceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RepairInvoiceQuery", q)
}

// The RepairInvoiceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RepairInvoiceMutationRuleFunc func(context.Context, *ent.RepairInvoiceMutation) error

// EvalMutation calls f(ctx, m).
func (f RepairInvoiceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RepairInvoiceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RepairInvoiceMutation", m)
}

// The StatusRQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StatusRQueryRuleFunc func(context.Context, *ent.StatusRQuery) error

// EvalQuery return f(ctx, q).
func (f StatusRQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StatusRQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StatusRQuery", q)
}

// The StatusRMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StatusRMutationRuleFunc func(context.Context, *ent.StatusRMutation) error

// EvalMutation calls f(ctx, m).
func (f StatusRMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StatusRMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StatusRMutation", m)
}

// The SymptomQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SymptomQueryRuleFunc func(context.Context, *ent.SymptomQuery) error

// EvalQuery return f(ctx, q).
func (f SymptomQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SymptomQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SymptomQuery", q)
}

// The SymptomMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SymptomMutationRuleFunc func(context.Context, *ent.SymptomMutation) error

// EvalMutation calls f(ctx, m).
func (f SymptomMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SymptomMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SymptomMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
