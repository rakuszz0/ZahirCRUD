package contactsdto

type CreateContactRequest struct {
	Name   string `json:"name" form:"name" gorm:"type: varchar(225)" validate:"required"`
	Gender string `json:"gender"  form:"gender" gorm:"type: varchar(225)"`
	Phone  string `json:"phone" form:"phone" gorm:"type: varchar(225)" validate:"required"`
	Email  string `json:"email" form:"email" gorm:"type: varchar(225)" validate:"required"`
}

type UpdateContactRequest struct {
	Name   string `json:"name" form:"name" gorm:"type: varchar(225)" validate:"required"`
	Gender string `json:"gender" gorm:"type: varchar(225)"`
	Phone  string `json:"phone" form:"phone" gorm:"type: varchar(225)" validate:"required"`
	Email  string `json:"email" form:"email" gorm:"type: varchar(225)" validate:"required"`
}
