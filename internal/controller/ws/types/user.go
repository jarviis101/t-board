package types

type (
	CreateUser struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginUserResponse struct {
		Token string `json:"token"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)
