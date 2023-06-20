package contactsdto

type ContactResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name" validate:"required"`
	Gender string `json:"gender" form:"gender" validate:"required"`
	Phone  string `json:"phone" form:"phone" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required"`
}
