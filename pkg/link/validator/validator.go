package validator

type Validator interface {
	Validate(link string) bool
}
