package model

// TeacherDetails represents teacher_details table
type TeacherDetails struct {
	TID        int64  `json:"id" gorm:"column:tid"`
	TFirstName string `json:"first_name" gorm:"column:tfirst_name"`
	TLastName  string `json:"last_name" gorm:"column:tlast_name"`
	Class      string `json:"father_name" gorm:"column:class"`
	Subject    string `json:"mother_name" gorm:"column:subject"`
}
