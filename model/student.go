package model

// StudentDetails represents student_details table
type StudentDetails struct {
	ID         int64  `json:"id" gorm:"column:id"`
	FirstName  string `json:"first_name" gorm:"column:first_name"`
	LastName   string `json:"last_name" gorm:"column:last_name"`
	FatherName string `json:"father_name" gorm:"column:father_name"`
	MotherName string `json:"mother_name" gorm:"column:mother_name"`
	TID        int64  `json:"tid" gorm:"column:tid"`
}

// func (this StudentDetails) Marshal() (outstr []byte, err error) {
// 	body := bytes.NewBuffer(nil)
// 	err = json.MarshalPayload(body, &this)
// 	if err == nil {
// 		outstr = body.Bytes()
// 	}
// 	return
// }
//

// func (this *StudentDetails) UnMarshal(data []byte) (err error) {
// 	buffer := bytes.NewBufferString(string(data))
// 	err = json.UnmarshalPayload(buffer, this)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

//SDR represents reponse for StudentDetails.
// type StudentDetailsResponse struct {
// 	Success bool             `json:"success"`
// 	Message string           `json:"message"`
// 	Errror  interface{}      `json:"error"`
// 	Data    []StudentDetails `json:"Data"`
// }
