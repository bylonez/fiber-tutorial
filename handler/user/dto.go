package user

type (
	User struct {
		Name string `validate:"required,min=3,max=20"` // Required field, min 3 char long max 20
	}
)
