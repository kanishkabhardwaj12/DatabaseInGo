package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// SaveToFile saves the in-memory data to JSON files
func (s *Storage) SaveToFile() error {
	// Save users to users.json
	if err := saveToFile("users.json", s.Users); err != nil {
		return err
	}

	// Save posts to posts.json
	if err := saveToFile("posts.json", s.Posts); err != nil {
		return err
	}

	// Save comments to comments.json
	if err := saveToFile("comments.json", s.Comments); err != nil {
		return err
	}

	// Save relationships (graph) to relationships.json
	if err := saveToFile("relationships.json", s.Relationships); err != nil {
		return err
	}

	fmt.Println("Data saved successfully!")
	return nil
}

// Helper function to save any object to a JSON file
func saveToFile(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadFromFile loads data from JSON files into memory
func (s *Storage) LoadFromFile() error {
	// Load users from users.json
	if err := loadFromFile("users.json", &s.Users); err != nil {
		return err
	}

	// Load posts from posts.json
	if err := loadFromFile("posts.json", &s.Posts); err != nil {
		return err
	}

	// Load comments from comments.json
	if err := loadFromFile("comments.json", &s.Comments); err != nil {
		return err
	}

	// Load relationships (graph) from relationships.json
	if err := loadFromFile("relationships.json", &s.Relationships); err != nil {
		return err
	}

	fmt.Println("Data loaded successfully!")
	return nil
}

// Helper function to load any JSON file into a provided object
func loadFromFile(filename string, data interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if len(file) == 0 {
		fmt.Printf("File %s is empty, starting with empty data.\n", filename)
		return nil
	}
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s does not exist, starting with empty data.\n", filename)
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateUser(username, email string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.Users {
		if user.Username == username {
			return 0, fmt.Errorf("username %s is already taken", username)
		}
		if user.Email == email {
			return 0, fmt.Errorf("email %s is already in use", email)
		}
	}

	newUserID := len(s.Users) + 1
	s.Users[newUserID] = User{ID: newUserID, Username: username, Email: email}
	return newUserID, nil
}

func (s *Storage) CreatePost(userID int, content string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	newPostID := len(s.Posts) + 1
	post := Post{ID: newPostID, UserID: userID, Content: content}
	s.Posts[newPostID] = post
	s.PostTree.Insert(post)
	return newPostID
}

func (s *Storage) CreateComment(postID, userID int, content string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Ensure post exists
	if _, ok := s.Posts[postID]; !ok {
		return 0, fmt.Errorf("post with ID %d does not exist", postID)
	}

	// Ensure user exists
	if _, ok := s.Users[userID]; !ok {
		return 0, fmt.Errorf("user with ID %d does not exist", userID)
	}

	newCommentID := len(s.Comments) + 1
	comment := Comment{ID: newCommentID, PostID: postID, UserID: userID, Content: content}
	s.Comments[newCommentID] = comment
	return newCommentID, nil
}

// FollowUser adds a follower-followee relationship using the Graph
func (s *Storage) FollowUser(followerID, followeeID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Relationships.AddEdge(followerID, followeeID)
}

// GetPostsByUser returns all posts created by a specific user
func (s *Storage) GetPostsByUser(userID int) []Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	posts := []Post{}
	for _, post := range s.Posts {
		if post.UserID == userID {
			posts = append(posts, post)
		}
	}
	return posts
}

// GetCommentsByPost returns all comments associated with a specific post
func (s *Storage) GetCommentsByPost(postID int) []Comment {
	s.mu.Lock()
	defer s.mu.Unlock()

	comments := []Comment{}
	for _, comment := range s.Comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments
}

// GetFollowers returns a list of users who are following a specific user
func (s *Storage) GetFollowers(userID int) []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Relationships.GetFollowers(userID)
}
