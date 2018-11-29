package db

import (
	"testing"
)

func countEmailInSlice(s []string, email string) (count int) {
	for _, e := range s {
		if e == email {
			count++
		}
	}
	return count
}

func TestInsertANewEmail(t *testing.T) {
	newEmail := "pudding850806+100@gmail.com"
	err := InsertAEmail(newEmail)
	if err != nil {
		t.Error(err)
	}
	emails, err := GetEmails()
	if err != nil {
		t.Error(err)
	}
	if countEmailInSlice(emails, newEmail) == 0 {
		t.Error("new email should be in db")
	}
	RemoveAEmail(newEmail)
}

func TestInsertADuplicateEmail(t *testing.T) {
	newEmail := "pudding850806+101@gmail.com"
	err := InsertAEmail(newEmail)
	if err != nil {
		t.Error(err)
	}
	err = InsertAEmail(newEmail)
	if err != nil {
		t.Error(err)
	}
	emails, err := GetEmails()
	if err != nil {
		t.Error(err)
	}
	if countEmailInSlice(emails, newEmail) != 1 {
		t.Errorf("there should be only a %s in db", newEmail)
	}
	RemoveAEmail(newEmail)
}

func TestRemoveAEmail(t *testing.T) {
	newEmail := "pudding850806+102@gmail.com"
	err := InsertAEmail(newEmail)
	if err != nil {
		t.Error(err)
	}
	err = RemoveAEmail(newEmail)
	if err != nil {
		t.Error(err)
	}
	emails, err := GetEmails()
	if err != nil {
		t.Error(err)
	}
	if countEmailInSlice(emails, newEmail) != 0 {
		t.Errorf("there should be NO %s in db", newEmail)
	}
}
