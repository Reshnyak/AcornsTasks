package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"id,omitempty"`
	Title       string             `json:"title" bson:"name"`
	Description string             `json:"description" bson:"description"`

	SubTasks []Task `json:"sub_tasks" bson:"sub_tasks,omitempty"`

	Status    string             `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"update_at" bson:"updated_at"`
	Deadline  primitive.DateTime `json:"deadline" bson:"deadline"`
}

type SubTask struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Deadline    primitive.DateTime `json:"deadline" bson:"deadline"`
}
