package mappers

import (
	"admin/jtypes"
	"admin/types"
)

func MapJStudentToPgStudent(d *jtypes.StudentDetails) *types.StudentDetails {
	return &types.StudentDetails{
		Name:        d.Name,
		Gender:      d.Gender,
		Dob:         d.Dob,
		Address:     d.Address,
		Pincode:     d.Pincode,
		District:    d.District,
		Contact:     int64(d.Contact),
		Email:       d.Email,
		Course_name: d.Course_name,
		Hostel_name: d.Hostel_name,
		Hostel_room: d.Hostel_room,
		Hostel_fee:  d.Hostel_fee,
	}
}
