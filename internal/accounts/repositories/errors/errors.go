package errors

type DuplicateEntity struct{}

func NewDuplicateEntity() *DuplicateEntity {
	return &DuplicateEntity{}
}

func (e *DuplicateEntity) Error() string {
	return "user account already exists"
}
