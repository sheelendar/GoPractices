package main

var uniqueCommentID = 1

//-----------------------------Comment-----------------------------

type Comment struct {
	id          int
	userID      int
	postID      int
	description string
	parantID    int
	comments    []*Comment
}

func getCommentUniqueID() int {
	id := uniqueCommentID
	uniqueCommentID++
	return id
}

func (comment *Comment) setNextedCommentDescription(commentID int, des string) {
	for _, nextComment := range comment.comments {
		if nextComment.GetId() == commentID {
			nextComment.description = des
			break
		}
	}
}
func (comment *Comment) deleteNextedComment(commentID int) {
	var newComment []*Comment
	for _, nextComment := range comment.comments {
		if nextComment.GetId() != commentID {
			newComment = append(newComment, nextComment)
		}
	}
	comment.comments = newComment
}

func GetCommentConstractor(des string) *Comment {
	return &Comment{id: getCommentUniqueID(), description: des}
}

func (c *Comment) addComment(comment *Comment) {
	c.comments = append(c.comments, comment)
}

func (comment *Comment) GetId() int {
	return comment.id
}

func (comment *Comment) GetUserID() int {
	return comment.userID
}

func (comment *Comment) GetPostID() int {
	return comment.postID
}

func (comment *Comment) GetDescription() string {
	return comment.description
}

func (comment *Comment) GetComments() []*Comment {
	return comment.comments
}

func (comment *Comment) SetId(id int) *Comment {
	comment.id = id
	return comment
}

func (comment *Comment) SetUserID(userID int) *Comment {
	comment.userID = userID
	return comment
}

func (comment *Comment) SetPostID(postID int) *Comment {
	comment.postID = postID
	return comment
}

func (comment *Comment) SetDescription(description string) *Comment {
	comment.description = description
	return comment
}

func (comment *Comment) SetComments(comments []*Comment) *Comment {
	comment.comments = comments
	return comment
}

func (comment *Comment) SetPatantID(id int) {
	comment.parantID = id

}

func (comment *Comment) GetPatantID() int {
	return comment.parantID

}
