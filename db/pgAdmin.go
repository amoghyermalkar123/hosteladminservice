package db

import (
	"admin/types"
	"log"
)

func CreateStudentDetails(dets *types.StudentDetails) error {
	db, err := getDbSession()

	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Insert(dets)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateStudentMessage(data *types.AdminMessage) error {
	db, err := getDbSession()

	if err != nil {
		log.Println(err)
		return err
	}

	err = db.Insert(data)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
