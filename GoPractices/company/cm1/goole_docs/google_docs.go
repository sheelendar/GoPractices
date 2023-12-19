package main

import "fmt"

/*
Document Storage system ( Similar to google docs ) Ability to create users and documents A document is
a string of characters for simplicity. The user who creates document is the sole owner of
the document. Permission management. Ability for the owner to assign read / write access
 to other users. Only users with correct permissions should be able to access. Hot and cold tier storage.
  Cold tier storage will be available even if the application is switched off. Do no use external database for stor


User
Docs
Permission
*/

var UserMap map[int]User

func main() {
	UserMap = make(map[int]User)
	startDocs()

}

func startDocs() {
	docName1 := "doc_1"

	createUser(23432, "sheelendar")
	createUser(12321, "Niki")
	user1 := UserMap[23432]
	user1.CreateDoc(docName1)

	printDocs(*user1.ListDocs[docName1])
	user1.GivenPermission(docName1, 12321, Read)
	printDocs(*user1.ListDocs[docName1])

	user1Doc1 := user1.ListDocs[docName1]
	user1Doc1.AppendInDoc("my name sheelendar")

	printDocs(*user1.ListDocs[docName1])
}

func printDocs(doc Docs) {
	fmt.Print("Msg:  ", doc.Doc, ",  ")
	fmt.Print("doc_name:   ", doc.Name, "  ", doc.CreatedTime, "   ", doc.SharedUsers)
	fmt.Println()
}
func createUser(userId int, userName string) {
	user := User{UserID: userId, UserName: userName}
	UserMap[userId] = user

}
