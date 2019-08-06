package schema

import (
	"time"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
)

type Todo struct {
	// ID      primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title     string           `json:"title,omitempty" bson:"title"`
	Checked   constant.Checker `json:"checked,omitempty" bson:"checked"`
	CreatedAt time.Time        `bson:"created_at"`
	UpdatedAt time.Time        `bson:"updated_at"`
}
