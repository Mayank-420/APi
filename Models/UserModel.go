package Models

type Person struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
type Address struct {
	Id      uint   `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	Zipcode string `json:"zipcode"`
}
type PhonrNo struct {
	Id          uint   `json:"id"`
	Phonenumber string `json:"phonenumber"`
}

func (b *Person) TableName() string {
	return "person"
}
