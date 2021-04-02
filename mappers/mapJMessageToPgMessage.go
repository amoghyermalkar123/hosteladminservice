package mappers

import (
	"admin/jtypes"
	"admin/types"
)

func MapJMessageToPgMessage(req *jtypes.AdminMessage) *types.AdminMessage {
	return &types.AdminMessage{
		HostelID:  req.HostelID,
		StudentID: req.StudentID,
		Text:      req.Text,
	}
}
