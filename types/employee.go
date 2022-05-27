package types

type Employee struct {
	ID       string  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Manager  string  `json:"manager"`
	Age      int     `json:"age"`
	Salary   float64 `json:"salary"`
}
