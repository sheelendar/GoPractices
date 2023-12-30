package main

import "fmt"

/*
APPLICATION REQUIREMENTS
1. Users should be able to comment to posts
2. Should be able to reply to comments
3. At most one level of nesting is allowed. If a user replies to a nested comment, the comment should appear below the
previous comment.
4. Users should be able to edit and delete comments. The nesting should adjust accordingly.

*/

func main() {
	user1 := getUserConstractor("Sheelendar")
	comment1 := GetCommentConstractor("first comment")
	comment2 := GetCommentConstractor("second comment")
	comment22 := GetCommentConstractor("second nexted comment")
	comment11 := GetCommentConstractor("first nexted comment")

	comment12 := GetCommentConstractor("first-2 nexted comment")
	comment13 := GetCommentConstractor("first-3 nexted comment")

	post1 := getPostConstractor()

	user1.addCommentOnPost(post1, comment1)
	user1.addCommentOnPost(post1, comment2)
	user1.replyOnComment(post1, comment1.GetId(), comment11)
	user1.replyOnComment(post1, comment1.GetId(), comment12)
	user1.replyOnComment(post1, comment1.GetId(), comment13)
	user1.replyOnComment(post1, comment2.GetId(), comment22)

	for _, comment := range post1.comments {
		fmt.Println(comment.description)
		for _, chieldC := range comment.comments {
			fmt.Println(chieldC.description)
		}
	}
	fmt.Println("**************************************************************************")
	s11 := "1st 1nd comment"

	//  Edit first first nexted comment
	user1.editComment(post1, comment1.GetId(), comment11.GetId(), s11)

	for _, comment := range post1.comments {
		fmt.Println(comment.description)
		for _, chieldC := range comment.comments {
			fmt.Println(chieldC.description)
		}
	}
	fmt.Println("**************************************************************************")

	// Deleted first second nexted comment
	user1.deleteComment(post1, comment1.GetId(), comment12.GetId())
	for _, comment := range post1.comments {
		fmt.Println(comment.description)
		for _, chieldC := range comment.comments {
			fmt.Println(chieldC.description)
		}
	}
	fmt.Println("**************************************************************************")

	// delete first comment completly
	user1.deleteComment(post1, comment1.GetId(), comment1.GetId())

	for _, comment := range post1.comments {
		fmt.Println(comment.description)
		for _, chieldC := range comment.comments {
			fmt.Println(chieldC.description)
		}
	}
	fmt.Println("**************************************************************************")

}
