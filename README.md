# 📊 CONTENT MANAGER APPLICATION FOR CONTENT CREATOR 

A simple command-line application built in Go for managing and analyzing content ideas for social media. This application has many features that really help content creators to manage their content creation.

---

## 🛠 Key Features

1.) Add new content with title, platform, category, timestamp, and engagement metrics (likes, comments, shares)
2.) Display all of stored contents
3.) Search content either by category or title:
    - Sequential search by title or category
    - Binary search by title (sorted automatically using bubble sort)
4.) Sort your contents either by post time or engagement:
    - By Post Time (Selection Sort)
    - By Engagement (Insertion Sort)
5.) Edit individual fields of content entries
6.) Delete content by ID
7.) View top-performing content in a specific period (`YYYY-MM`)

---

## 🎯 Purpose

This program helps users manage and analyze social media content ideas. Users can add, view, edit, search, sort, and delete content easily using a simple menu.
The goal is to:
  -Organize content with details like title, platform, and post time
  -Track engagement through likes, comments, and shares
  -Find top content by period
  -Practice using arrays, structs, and basic algorithms in Go
It’s also a learning project to improve programming skills in Go.

---

## 🧱 Data Structure

All content is stored in a fixed-size array of `Content` structs:

type Content struct {
	ID       int
	Title    string
	Platform string
	Category string
	PostTime string
	Likes    int
	Comments int
	Shares   int
}

---

## 🧑‍💻 User Interaction
  -Command-line interface (CLI)
  -Menu-driven navigation (1–9 options)

---

## 📌 Limitations
  -Max content: 100 entries
  -Data is stored temporarily in memory (we're currently working on this so it can be for long      time use)
  -Inputs are taken using fmt.Scan() (no validation or advanced input handling)

---

## ▶️ How to Run 
  1. Install Go language if you haven't install it yet: https://golang.org/doc/install
  2. Put it in notepad++ or Visual Studio Code, save the code in a file, e.g. 'main.go'
  3. Open the terminal and navigate to the directory where the file is located
  4. Run the program at the terminal:
     'go run main.go'
  5. The application will print the main menu, you can choose which feature you want to use by        typing the number of the feature
  6. Enjoy your application!


