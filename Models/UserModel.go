package Models

type Person struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName   string `json:"lastname"`
	DOB   string `json:"phone"`
	Address string `json:"address"`
	Subject string  `json:"subject"`
	Marks string   `json:"marks"`

}
func (b *Person) TableName() string {
	return "person"
}