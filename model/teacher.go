package model

// TeacherDetails represents teacher_details table
type TeacherDetails struct {
	TID        int64  `json:"id" gorm:"column:tid"`
	TFirstName string `json:"first_name" gorm:"column:tfirst_name"`
	TLastName  string `json:"last_name" gorm:"column:tlast_name"`
	Class      string `json:"class" gorm:"column:class"`
	Subject    string `json:"subject" gorm:"column:subject"`
}


//SDR represents reponse for TeacherDetails.
type TeacherDetailsResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Errror  interface{}      `json:"error"`
	Data    []TeacherDetails `json:"Data"`
}

type TeacherDetailResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Errror  interface{}    `json:"error"`
	Data    TeacherDetails `json:"Data"`
}