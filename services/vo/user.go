package vo

import (
	"RizhaoLanshanLabourUnion/security/jwt/jwtmodel"
	"RizhaoLanshanLabourUnion/services/models"
)

type UserData struct {
	Common
	UserDataPayload
}

type UserDataPayload struct {
	User    *models.User         `json:"user"`
	Claims  *jwtmodel.UserClaims `json:"claims"`
	Profile *models.UserProfile  `json:"profile"`
}

type UserResetPassword struct {
	OldPassword     string `json:"old_password" form:"old_password" binding:"required,min=3,max=20" `
	NewPassword     string `json:"new_password" form:"new_password" binding:"required,min=3,max=20" `
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,min=3,max=20,eqfield=NewPassword" `
}

type UserUpdateInfo struct {
	Email string `json:"email" form:"email" binding:"omitempty,email" example:"hafrans@163.com"`
	Phone string `json:"phone" form:"phone" binding:"omitempty,numeric" example:"13800138000"`
}

type DepartmentUserUpdateInfo struct {
	DepartmentId string `json:"department_id"`
	UserName     string `json:"user_name"`
}

type LaborUserUpdateInfo struct {
	Applicant `binding:"dive"`
}

type EmployerUserUpdateInfo struct {
	Employer `binding:"dive"`
}

type UserRegisterLaborForm struct {
	//Username        string    `json:"username" form:"username" binding:"required,min=3,max=20"`
	Password        string    `json:"password" form:"password" binding:"required,min=3,max=20"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,eqfield=Password"`
	Phone           string    `json:"phone" binding:"required,numeric" example:"13800138000"`
	Email           string    `json:"email" binding:"required,email" example:"test@1234.com"`
	Applicant       Applicant `json:"applicant" binding:"required,dive"`
}

type UserRegisterEmployerForm struct {
	//Username        string   `json:"username" form:"username" binding:"required,min=3,max=20"`
	Password        string   `json:"password" form:"password" binding:"required,min=3,max=20"`
	ConfirmPassword string   `json:"confirm_password" binding:"required,eqfield=Password"`
	Phone           string   `json:"phone" binding:"required,numeric" example:"13800138000"`
	Email           string   `json:"email" binding:"required,email" example:"test@1234.com"`
	Employer        Employer `json:"employer" binding:"required,dive"`
}
