package main

import "fmt"

func main() {
	// Initialize storage
	storage := Storage{
		Users:         make(map[int]User),
		Posts:         make(map[int]Post),
		Comments:      make(map[int]Comment),
		PostTree:      PostTree{},
		Relationships: Graph{},
	}

	// Load existing data from JSON files (if they exist)
	if err := storage.LoadFromFile(); err != nil {
		fmt.Println("Error loading data:", err)
	}

	// Create users
	user1ID, err := storage.CreateUser("john_doe", "john@example.com")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	user2ID, err := storage.CreateUser("jane_smith", "jane@example.com")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// Create a post
	postID := storage.CreatePost(user1ID, "Hello, this is my first post!")

	// Create a comment
	storage.CreateComment(user2ID, postID, "Nice post!")

	// Follow a user
	storage.FollowUser(user2ID, user1ID)

	// Get posts by user
	posts := storage.GetPostsByUser(user1ID)
	fmt.Printf("User %d's posts: %+v\n", user1ID, posts)

	// Get comments on a post
	comments := storage.GetCommentsByPost(postID)
	fmt.Printf("Comments on post %d: %+v\n", postID, comments)

	// Get followers
	followers := storage.GetFollowers(user1ID)
	fmt.Printf("User %d's followers: %+v\n", user1ID, followers)

	// Save data to JSON files for future use
	if err := storage.SaveToFile(); err != nil {
		fmt.Println("Error saving data:", err)
	}
}
