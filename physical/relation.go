package physical

import (
	"context"
	"log"

	"github.com/cube2222/octosql/execution"
)

type Relation string

const (
	Equal    Relation = "equal"
	NotEqual Relation = "not_equal"
	MoreThan Relation = "more_than"
	LessThan Relation = "less_than"
	Like     Relation = "like"
	In       Relation = "in"
)

func NewRelation(relation string) Relation {
	return Relation(relation)
}

func (rel Relation) Materialize(ctx context.Context) execution.Relation {
	switch rel {
	case Equal:
		return execution.NewEqual()
	case NotEqual:
		return execution.NewNotEqual()
	case MoreThan:
		return execution.NewMoreThan()
	case LessThan:
		return execution.NewLessThan()
	case Like:
		return execution.NewLike()
	case In:
		return execution.NewIn()
	default:
		log.Fatalf("Invalid relation: %+v", rel)
		return nil
	}
}
