package logic

import (
	"../../dbops"
	"../entity"
	"fmt"
	"log"
)

var (
	sql_of_get_one = "SELECT * FROM `meeting` WHERE `id` = ? ;"
	sql_of_remove_one = "DELETE FORM `meeting` WHERE `id` = ? ;"
)

func GetMeeting(id int) (entity.Meeting, error) {
	stmt, err := dbops.Connection.Prepare(sql_of_get_one)
	if err != nil {
		log.Printf("%s \n", err)
		return entity.Meeting{}, err
	}

	var meeting entity.Meeting
	//err = stmt.QueryRow(id).Scan(&meeting)
	//if err != nil {
	//	log.Printf("%s \n", err)
	//	return entity.Meeting{}, err
	//}

	row := stmt.QueryRow(id)

	fmt.Println(row.Scan(meeting))

	defer stmt.Close()
	return entity.Meeting{}, nil
}

func ListMeetings(page int, pageSize int, keywords string) ([]entity.Meeting, error) {
	return []entity.Meeting{}, nil
}

func RemoveMeeting(id int) error {
	stmt, err := dbops.Connection.Prepare(sql_of_remove_one)
	if err != nil {
		log.Printf("%s \n", err)
		return err
	}

	stmt.Exec(id)
	defer stmt.Close()
	return nil
}
