package jtypes

type AdminMessage struct {
	HostelID  int64  `json:"hostelID"`
	StudentID int64  `json:"studentID"`
	Text      string `json:"text"`
}
