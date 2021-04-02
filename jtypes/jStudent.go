package jtypes

type StudentDetails struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Pincode     int    `json:"pincode"`
	District    string `json:"district"`
	Contact     int    `json:"contact"`
	Email       string `json:"email"`
	Course_name string `json:"course_name"`
	Hostel_name string `json:"hostel_name"`
	Hostel_room string `json:"hostel_room"`
	Hostel_fee  int    `json:"hostel_fee"`
}
