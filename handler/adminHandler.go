package handler

import (
	"admin/db"
	"admin/ipc"
	"admin/jtypes"
	"admin/mappers"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/segmentio/encoding/json"
)

type Handler struct{}

func (h *Handler) RegisterNewStudent(c echo.Context) error {
	student := &jtypes.StudentDetails{}

	if err := c.Bind(student); err != nil {
		c.Error(err)
		return err
	}

	pgStudent := mappers.MapJStudentToPgStudent(student)

	err := db.CreateStudentDetails(pgStudent)

	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, "Student was successfully registered")
}

func (h *Handler) CreateStudentMessage(c echo.Context) error {
	message := &jtypes.AdminMessage{}

	if err := c.Bind(message); err != nil {
		c.Error(err)
		return err
	}

	pgMsg := mappers.MapJMessageToPgMessage(message)

	err := db.CreateStudentMessage(pgMsg)

	if err != nil {
		return err
	}

	msg, err := json.Marshal(message)

	if err != nil {
		log.Println("err in marshal", err)
		go ipc.Produce([]byte("message failed to retrieve"))
	} else {
		go ipc.Produce(msg)
	}

	return c.JSON(http.StatusOK, "Message sent successfully")
}
