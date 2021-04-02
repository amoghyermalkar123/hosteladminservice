package types

type AdminMessage struct {
	tableName struct{}
	ID        int64  `json:"id"`
	HostelID  int64  `json:"hostelID"`
	StudentID int64  `json:"studentID"`
	Text      string `json:"text"`
}
