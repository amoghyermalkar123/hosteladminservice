package types

type StudentDetails struct {
	tableName   struct{} `pg:"student_details,alias:student_details"`
	ID          int64    `pg:"id"`
	Name        string   `pg:"name"`
	Gender      string   `pg:"gender"`
	Dob         string   `pg:"dob"`
	Address     string   `pg:"address"`
	Pincode     int      `pg:"pincode"`
	District    string   `pg:"district"`
	Contact     int64    `pg:"contact"`
	Email       string   `pg:"email"`
	Course_name string   `pg:"course_name"`
	Hostel_name string   `pg:"hostel_name"`
	Hostel_room string   `pg:"hostel_room"`
	Hostel_fee  int      `pg:"hostel_fee"`
}
