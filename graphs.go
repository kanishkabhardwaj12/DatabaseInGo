package main

//Define the graph data structure using an adjacency list to represent user relationships (followers/followees).

type Graph struct {
	adjList map[int][]int
}

func (g *Graph) AddEdge(followerID, followeeID int) {
	if g.adjList == nil {
		g.adjList = make(map[int][]int)
	}
	g.adjList[followerID] = append(g.adjList[followerID], followeeID)
}

func (g *Graph) GetFollowers(userID int) []int {
	followers := []int{}
	for follower, followees := range g.adjList {
		for _, followee := range followees {
			if followee == userID {
				followers = append(followers, follower)
			}
		}
	}
	return followers
}

func (g *Graph) GetFollowees(userID int) []int {
	return g.adjList[userID]
}
