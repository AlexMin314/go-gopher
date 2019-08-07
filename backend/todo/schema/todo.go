package schema

import (
	"time"

	"github.com/AlexMin314/go-gopher/backend/todo/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title     string             `json:"title,omitempty" bson:"title"`
	Checked   constant.Checker   `json:"checked,omitempty" bson:"checked"`
	CreatedAt time.Time          `json:"-" bson:"created_at"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at,omitempty"`
}
