package logic

import (
	"fmt"
	"testing"
)

func TestGetMeeting(t *testing.T) {
	meeting, err := GetMeeting(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(meeting)
}
