package main

// Define a simple Binary Search Tree (BST) to store posts for efficient sorting and searching.
type PostNode struct {
	post  Post
	left  *PostNode
	right *PostNode
}

type PostTree struct {
	root *PostNode
}

func (t *PostTree) Insert(post Post) {
	newNode := &PostNode{post: post}
	if t.root == nil {
		t.root = newNode
		return
	}
	t.root = insertPostNode(t.root, newNode)
}

func insertPostNode(node, newNode *PostNode) *PostNode {
	if node == nil {
		return newNode
	}
	if newNode.post.ID < node.post.ID {
		node.left = insertPostNode(node.left, newNode)
	} else {
		node.right = insertPostNode(node.right, newNode)
	}
	return node
}

func (t *PostTree) Search(postID int) *Post {
	return searchPostNode(t.root, postID)
}

func searchPostNode(node *PostNode, postID int) *Post {
	if node == nil {
		return nil
	}
	if postID == node.post.ID {
		return &node.post
	}
	if postID < node.post.ID {
		return searchPostNode(node.left, postID)
	}
	return searchPostNode(node.right, postID)
}
