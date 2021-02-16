package logical

import (
	"context"
	"fmt"

	"github.com/cube2222/octosql/octosql"
	"github.com/cube2222/octosql/physical"
)

type Trigger interface {
	Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment, keyTimeIndex int) physical.Trigger
}

type CountingTrigger struct {
	Count uint
}

func NewCountingTrigger(count uint) *CountingTrigger {
	return &CountingTrigger{Count: count}
}

func (w *CountingTrigger) Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment, keyTimeIndex int) physical.Trigger {
	return physical.Trigger{
		TriggerType: physical.TriggerTypeCounting,
		CountingTrigger: &physical.CountingTrigger{
			TriggerAfter: w.Count,
		},
	}
}

type DelayTrigger struct {
	Delay Expression
}

func NewDelayTrigger(delay Expression) *DelayTrigger {
	return &DelayTrigger{Delay: delay}
}

func (w *DelayTrigger) Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment, keyTimeIndex int) physical.Trigger {
	panic("implement me")
}

type WatermarkTrigger struct {
}

func NewWatermarkTrigger() *WatermarkTrigger {
	return &WatermarkTrigger{}
}

func (w *WatermarkTrigger) Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment, keyTimeIndex int) physical.Trigger {
	if keyTimeIndex == -1 {
		panic(fmt.Errorf("can't use watermark trigger when not grouping by time field"))
	}
	return physical.Trigger{
		TriggerType: physical.TriggerTypeWatermark,
		WatermarkTrigger: &physical.WatermarkTrigger{
			TimeFieldIndex: keyTimeIndex,
		},
	}
}

type GroupBy struct {
	source   Node
	key      []Expression
	keyNames []string

	expressions    []Expression
	aggregates     []string
	aggregateNames []string

	triggers []Trigger
}

func NewGroupBy(source Node, key []Expression, keyNames []string, expressions []Expression, aggregates []string, aggregateNames []string, triggers []Trigger) *GroupBy {
	return &GroupBy{source: source, key: key, keyNames: keyNames, expressions: expressions, aggregates: aggregates, aggregateNames: aggregateNames, triggers: triggers}
}

func (node *GroupBy) Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment) physical.Node {
	source := node.source.Typecheck(ctx, env, logicalEnv)

	keyTimeIndex := -1

	key := make([]physical.Expression, len(node.key))
	for i := range node.key {
		key[i] = node.key[i].Typecheck(ctx, env.WithRecordSchema(source.Schema), logicalEnv)
		if source.Schema.TimeField != -1 &&
			key[i].ExpressionType == physical.ExpressionTypeVariable &&
			physical.VariableNameMatchesField(key[i].Variable.Name, source.Schema.Fields[source.Schema.TimeField].Name) {

			keyTimeIndex = i
		}
	}

	expressions := make([]physical.Expression, len(node.expressions))
	for i := range node.expressions {
		expressions[i] = node.expressions[i].Typecheck(ctx, env.WithRecordSchema(source.Schema), logicalEnv)
	}

	aggregates := make([]physical.Aggregate, len(node.aggregates))
aggregateLoop:
	for i, aggname := range node.aggregates {
		descriptors := env.Aggregates[aggname]
		for _, descriptor := range descriptors {
			if expressions[i].Type.Is(descriptor.ArgumentType) == octosql.TypeRelationIs {
				aggregates[i] = physical.Aggregate{
					Name:                node.aggregates[i],
					AggregateDescriptor: descriptor,
				}
				continue aggregateLoop
			}
		}
		for _, descriptor := range descriptors {
			if expressions[i].Type.Is(descriptor.ArgumentType) == octosql.TypeRelationMaybe {
				aggregates[i] = physical.Aggregate{
					Name:                node.aggregates[i],
					AggregateDescriptor: descriptor,
				}
				expressions[i] = physical.Expression{
					ExpressionType: physical.ExpressionTypeTypeAssertion,
					Type:           *octosql.TypeIntersection(descriptor.ArgumentType, expressions[i].Type),
					TypeAssertion: &physical.TypeAssertion{
						Expression: expressions[i],
						TargetType: descriptor.ArgumentType,
					},
				}
				continue aggregateLoop
			}
		}
		panic(fmt.Sprintf("unknown aggregate: %s(%s)", aggname, expressions[i].Type))
	}

	triggers := make([]physical.Trigger, len(node.triggers))
	for i := range node.triggers {
		triggers[i] = node.triggers[i].Typecheck(ctx, env, logicalEnv, keyTimeIndex)
	}
	var trigger physical.Trigger
	if len(triggers) == 0 {
		trigger = physical.Trigger{
			TriggerType:        physical.TriggerTypeEndOfStream,
			EndOfStreamTrigger: &physical.EndOfStreamTrigger{},
		}
	} else if len(triggers) == 1 {
		trigger = triggers[0]
	} else {
		trigger = physical.Trigger{
			TriggerType: physical.TriggerTypeMulti,
			MultiTrigger: &physical.MultiTrigger{
				Triggers: triggers,
			},
		}
	}

	schemaFields := make([]physical.SchemaField, len(key)+len(aggregates))
	for i := range key {
		schemaFields[i] = physical.SchemaField{
			Name: node.keyNames[i],
			Type: key[i].Type,
		}
	}
	for i := range aggregates {
		schemaFields[len(key)+i] = physical.SchemaField{
			Name: node.aggregateNames[i],
			Type: aggregates[i].AggregateDescriptor.OutputType,
		}
	}

	// TODO: Calculate time field if grouping by time field.

	return physical.Node{
		Schema:   physical.NewSchema(schemaFields, -1),
		NodeType: physical.NodeTypeGroupBy,
		GroupBy: &physical.GroupBy{
			Source:               source,
			Aggregates:           aggregates,
			AggregateExpressions: expressions,
			Key:                  key,
			Trigger:              trigger,
		},
	}
}
