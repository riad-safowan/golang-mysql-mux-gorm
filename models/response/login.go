package response

type LoginResponse struct {
	ID           uint    `json:"user_id"`
	FirstName    *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string `json:"last_name" validate:"required"`
	Email        *string `json:"email" validate:"email,required"`
	ImageUrl     *string `json:"image_url"`
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}
