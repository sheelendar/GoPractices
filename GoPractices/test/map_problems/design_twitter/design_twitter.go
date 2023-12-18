package main

import "sort"

func main() {

}

// follower[1]={1,2,3,4}
// post[1] ={....newpostID}
// post[2] ={....newpostID}
// post[3] ={....newpostID}
// post[4] ={....newpostID}

type Twitter struct {
	follower map[int][]int
	post     map[int][]int
	myPost   map[int][]int
}

func Constructor() Twitter {
	return Twitter{follower: make(map[int][]int), post: make(map[int][]int), myPost: make(map[int][]int)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.myPost[userId] = append(this.myPost[userId], tweetId)
	size := len(this.post[userId])
	if size > 10 {
		this.post[userId] = this.post[userId][0 : size-1]
	}
	this.post[userId] = append([]int{tweetId}, this.post[userId]...)
	//this.post[userId] = append(this.post[userId], tweetId)
	f := this.follower[userId]
	if f == nil {
		return
	}
	for i := 0; i < len(f); i++ {
		size := len(this.post[f[i]])
		if size > 10 {
			this.post[f[i]] = this.post[f[i]][1:size]
		}
		this.post[f[i]] = append([]int{tweetId}, this.post[f[i]]...)
	}

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return this.post[userId]
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.follower[followeeId] = append(this.follower[followeeId], followerId)

	size := len(this.myPost[followeeId])
	if size == 0 {
		return
	}
	post := this.post[followerId]
	for i := 0; i < size; i++ {
		post = append(post, this.myPost[followeeId][i])
	}
	sort.Slice(post, func(i, j int) bool {
		return post[i] < post[j]
	})
	if len(post) > 10 {
		post = post[0:10]
	}
	this.post[followerId] = post
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	f := this.follower[followeeId]
	i := 0
	var ff []int
	for ; i < len(f); i++ {
		if f[i] != followerId {
			ff = append(ff, f[i])
		}
	}
	this.follower[followeeId] = ff

	deleteP := this.post[followeeId]
	dpm := make(map[int]bool)
	for i := 0; i < len(deleteP); i++ {
		dpm[deleteP[i]] = true
	}

	post := this.post[followerId]
	var newPost []int
	for i := 0; i < len(post); i++ {
		if !dpm[post[i]] {
			newPost = append(newPost, post[i])
		}
	}

	this.post[followerId] = newPost
}
