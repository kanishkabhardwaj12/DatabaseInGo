package main

import "sync"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type Post struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt string
}

type Comment struct {
	ID        int
	PostID    int
	UserID    int
	Content   string
	CreatedAt string
}

type Relationship struct {
	ID       int
	Follower int
	Followee int
}

type Storage struct {
	mu            sync.Mutex
	Users         map[int]User    // HashMap for Users
	Posts         map[int]Post    // HashMap for Posts
	Comments      map[int]Comment // HashMap for Comments
	PostTree      PostTree        // Binary Search Tree for Posts
	Relationships Graph           // Graph for User Relationships
}
