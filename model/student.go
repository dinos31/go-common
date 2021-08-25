package model

// StudentDetails represents student_details table
type StudentDetails struct {
	ID         int64  `json:"id" gorm:"column:id"`
	FirstName  string `json:"first_name" gorm:"column:first_name"`
	LastName   string `json:"last_name" gorm:"column:last_name"`
	FatherName string `json:"father_name" gorm:"column:father_name"`
	MotherName string `json:"mother_name" gorm:"column:mother_name"`
}
