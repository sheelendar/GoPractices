package main

import (
	"fmt"
	"time"
)

const (
	Read = iota + 1
	Write
)

type User struct {
	UserID    int
	UserName  string
	UserEmail string
	ListDocs  map[string]*Docs
}

func (u *User) CreateDoc(docName string) Docs {
	doc := Docs{Name: docName}
	doc.CreatedTime = time.Now()
	doc.UserId = u.UserID
	doc.Doc = "Hi"
	doc.SharedUsers = make(map[int]map[int]bool)

	if u.ListDocs == nil {
		u.ListDocs = make(map[string]*Docs)
	}

	u.ListDocs[docName] = &doc
	return doc
}

func (u *User) GivenPermission(docName string, givenUserID int, permision int) string {
	if u.ListDocs == nil {
		return "No doc found for this user"
	}

	doc, ok := u.ListDocs[docName]
	if !ok {
		return "No doc found with name " + docName

	}

	doc.SharedUsers[givenUserID] = map[int]bool{permision: true}
	u.ListDocs[docName] = doc
	return "permission given"
}

func (u User) RemovePermission(docName string, givenUserID int, permision int) string {
	if u.ListDocs == nil {
		return "No doc found for this user"
	}

	doc, ok := u.ListDocs[docName]
	if !ok {
		return "No doc found with name " + docName
	}
	perMap := doc.SharedUsers[givenUserID]
	if perMap[permision] {
		delete(perMap, permision)
	}
	doc.SharedUsers[givenUserID] = perMap
	return "permission deleted"
}

type Docs struct {
	Name        string
	Doc         string
	CreatedTime time.Time
	SharedUsers map[int]map[int]bool
	UserId      int
}

func (d *Docs) WriteInDoc(msg string) {
	d.Doc = msg
}

func (d *Docs) AppendInDoc(msg string) {
	d.Doc = fmt.Sprintf("%s %s", d.Doc, msg)
}
