package main

type BlogPost struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}