package services

import (
	"admin/ipc"
	"admin/proto/api"
)

type AdminService struct{}

func NewAdminService() (*AdminService, error) {
	return &AdminService{}, nil
}

func (a *AdminService) MonitorAdminMessages(req *api.StudentAdminMessageRequest, srv api.AdminMessageService_MonitorAdminMessagesServer) error {
	reader := make(chan string)
	go ipc.Consume(reader)

	for {
		select {
		case val := <-reader:
			err := srv.Send(
				&api.ChatMessage{
					Text: val,
				},
			)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
