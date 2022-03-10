package entity

type User struct {
	ID   uint64  `json:"id"`
	Name string  `json:"name"`
	Age  int     `json:"age"`
	Sex  string  `json:"sex"`
	BMI  float64 `json:"bmi"`
}
