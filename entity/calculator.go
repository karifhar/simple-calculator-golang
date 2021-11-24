package entity

type Calculator struct {
	Number1  float32 `json:"number1"`
	Number2  float32 `json:"number2"`
	Operator string  `json:"operator"`
}

// membuat result ouput json secara primitive
type Status struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Result  float32 `json:"result"`
}
