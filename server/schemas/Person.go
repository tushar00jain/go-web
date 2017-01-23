package schemas

type PhoneType int

const (
	MOBILE PhoneType = iota
	HOME
	WORK
)

type PhoneNumber struct {
	Number string `json:"Number"`
	Type   PhoneType `json:"Type"`
}

type Person struct {
	Id          int
	Name        string
	Email       string
	PhoneNumbers []PhoneNumber
}
