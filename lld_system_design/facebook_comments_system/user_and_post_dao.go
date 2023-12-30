package main

import "fmt"

var uniquePostID = 1

var uniqueUserID = 1

func getPostUniqueID() int {
	id := uniquePostID
	uniquePostID++
	return id
}

func getUserUniqueID() int {
	id := uniqueUserID
	uniqueUserID++
	return id
}

type User struct {
	userID int
	name   string
}

func getUserConstractor(name string) *User {
	return &User{userID: getUserUniqueID(), name: name}
}

func (user *User) addCommentOnPost(post *Post, comment *Comment) {
	comment.SetUserID(user.userID)
	comment.SetPostID(post.id)
	comment.SetPatantID(comment.id)
	post.comments[comment.id] = comment
}

func (user *User) replyOnComment(post *Post, parantID int, comment *Comment) {
	comment.SetUserID(user.userID)
	comment.SetPostID(post.id)
	comment.SetPatantID(comment.id)
	post.addNextedComment(parantID, comment)
}

func (user *User) editComment(post *Post, parantID, commentID int, des string) bool {
	comment := post.comments[parantID]
	if comment.GetId() == commentID {
		if comment.GetUserID() != user.userID {
			fmt.Println("Unauthorized to edit comment!")
			return false
		}
	}
	post.editComment(parantID, commentID, des)
	return true
}

func (user *User) deleteComment(post *Post, parantID, commentID int) bool {
	comment := post.comments[parantID]
	if comment.GetId() == commentID {
		if comment.GetUserID() != user.userID {
			fmt.Println("Unauthorized to detete comment!")
			return false
		}
	}
	post.deleteComment(parantID, commentID)
	return true
}

// -----------------------------Post-----------------------------
type Post struct {
	id       int
	comments map[int]*Comment
}

func getPostConstractor() *Post {
	return &Post{id: getPostUniqueID(), comments: make(map[int]*Comment)}
}
func (post *Post) addComment(comment *Comment) {
	post.comments[comment.id] = comment
}

func (post *Post) editComment(parentID, commentID int, description string) {
	c := post.comments[parentID]
	if c.GetId() == commentID {
		c.description = description
	} else {
		c.setNextedCommentDescription(commentID, description)
	}
}

func (post *Post) deleteComment(parentID, commentID int) {
	c := post.comments[parentID]
	if c.GetId() == commentID {
		delete(post.comments, commentID)
	} else {
		c.deleteNextedComment(commentID)
	}
}

func (post *Post) addNextedComment(commentID int, comment *Comment) {
	post.comments[commentID].comments = append(post.comments[commentID].comments, comment)
}
