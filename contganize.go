package main

import "fmt"

const maxContents = 100

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

func main() {
	// Initial State:
	// - No content has been added yet
	// - contentCount and nextID start from 0 and 1 respectively
	//
	// Final State:
	// - Application continues to run and update data via menu interaction

	var contents [maxContents]Content
	var contentCount = 0
	var nextID = 1

	for {
		fmt.Println("\n=== Welcome to Contganize! ===")
		fmt.Println("1. Add Content")
		fmt.Println("2. Show All Contents")
		fmt.Println("3. Search Content Idea (Sequential)")
		fmt.Println("4. Search Content Idea (Binary by Title)")
		fmt.Println("5. Sort")
		fmt.Println("6. Edit Content")
		fmt.Println("7. Delete Content")
		fmt.Println("8. Show Top Content by Period")
		fmt.Println("9. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			contents, contentCount, nextID = addContent(contents, contentCount, nextID)
		case 2:
			showAllContents(contents, contentCount)
		case 3:
			sequentialSearch(contents, contentCount)
		case 4:
			binarySearch(contents, contentCount)
		case 5:
			contents = sortContent(contents, contentCount)
		case 6:
			contents = editContent(contents, contentCount)
		case 7:
			contents, contentCount = deleteContent(contents, contentCount)
		case 8:
			showTopContentByPeriod(contents, contentCount)
		case 9:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addContent(contents [maxContents]Content, contentCount int, nextID int) ([maxContents]Content, int, int) {
	// Initial State:
	// - contents: list of existing content
	// - contentCount: total contents stored
	// - nextID: next available unique ID
	//
	// Final State:
	// - A new content is added to the list
	// - contentCount and nextID are both incremented

	if contentCount >= maxContents {
		fmt.Println("Sorry, your content list is full.")
		return contents, contentCount, nextID
	}

	var title, platform, category string
	var postDate, postHour string
	var likes, comments, shares int

	fmt.Print("Title: ")
	fmt.Scan(&title)
	fmt.Print("Platform: ")
	fmt.Scan(&platform)
	fmt.Print("Category: ")
	fmt.Scan(&category)
	fmt.Print("Post Date & Time (YYYY-MM-DD HH:MM): ")
	fmt.Scan(&postDate, &postHour)
	fmt.Print("Likes: ")
	fmt.Scan(&likes)
	fmt.Print("Comments: ")
	fmt.Scan(&comments)
	fmt.Print("Shares: ")
	fmt.Scan(&shares)

	contents[contentCount] = Content{
		ID:       nextID,
		Title:    title,
		Platform: platform,
		Category: category,
		PostTime: postDate + " " + postHour,
		Likes:    likes,
		Comments: comments,
		Shares:   shares,
	}

	fmt.Println("Content added successfully!")
	return contents, contentCount + 1, nextID + 1
}

func showAllContents(contents [maxContents]Content, contentCount int) {
	// Initial State:
	// - contents: array containing added content
	// - contentCount: number of stored contents
	//
	// Final State:
	// - All contents are printed out
	// - No data is changed

	if contentCount == 0 {
		fmt.Println("No content available.")
		return
	}
	for i := 0; i < contentCount; i++ {
		c := contents[i]
		eng := c.Likes + c.Comments + c.Shares
		fmt.Printf("ID:%d | %s | %s | %s | %s | %d | %d | %d | Engagement: %d\n",
			c.ID, c.Title, c.Platform, c.Category, c.PostTime, c.Likes, c.Comments, c.Shares, eng)
	}
}

func sequentialSearch(contents [maxContents]Content, contentCount int) {
	// Initial State:
	// - contents: all content entries
	// - contentCount: total number of contents
	// - keyword input is required from user
	//
	// Final State:
	// - Matching content(s) printed if found
	// - No data is modified

	var kw string
	fmt.Print("Enter keyword (title/category): ")
	fmt.Scan(&kw)

	found := false
	for i := 0; i < contentCount; i++ {
		if contents[i].Title == kw || contents[i].Category == kw {
			c := contents[i]
			fmt.Printf("ID:%d | %s | %s | %s | %s | %d | %d | %d\n",
				c.ID, c.Title, c.Platform, c.Category, c.PostTime, c.Likes, c.Comments, c.Shares)
			found = true
		}
	}
	if !found {
		fmt.Println("No matching content found.")
	}
}

func binarySearch(contents [maxContents]Content, contentCount int) {
	// Initial State:
	// - contents: all content entries
	// - contentCount: total number of contents
	// - keyword input from user (Title)
	// - array will be sorted first using Bubble Sort
	//
	// Final State:
	// - Content with matching title is displayed
	// - No data is permanently modified

	for i := 0; i < contentCount-1; i++ {
		for j := 0; j < contentCount-i-1; j++ {
			if contents[j].Title > contents[j+1].Title {
				contents[j], contents[j+1] = contents[j+1], contents[j]
			}
		}
	}

	var kw string
	fmt.Print("Enter title to search: ")
	fmt.Scan(&kw)

	low, high := 0, contentCount-1
	for low <= high {
		mid := (low + high) / 2
		if contents[mid].Title == kw {
			c := contents[mid]
			fmt.Printf("Found: ID:%d | %s | %s | %s | %s | %d | %d | %d\n",
				c.ID, c.Title, c.Platform, c.Category, c.PostTime, c.Likes, c.Comments, c.Shares)
			return
		} else if contents[mid].Title < kw {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Content not found.")
}

func sortContent(contents [maxContents]Content, contentCount int) [maxContents]Content {
	// Initial State:
	// - contents: list of all stored content
	// - contentCount: total number of content
	// - user selects sorting method (1 or 2)
	//
	// Final State:
	// - contents is sorted either by PostTime or Engagement
	// - contents array order is permanently changed

	fmt.Println("Sort by:")
	fmt.Println("1. Post Time (Selection Sort)")
	fmt.Println("2. Engagement (Insertion Sort)")
	fmt.Print("Choose: ")
	var opt int
	fmt.Scan(&opt)

	switch opt {
	case 1:
		for i := 0; i < contentCount-1; i++ {
			minIdx := i
			for j := i + 1; j < contentCount; j++ {
				if contents[j].PostTime < contents[minIdx].PostTime {
					minIdx = j
				}
			}
			contents[i], contents[minIdx] = contents[minIdx], contents[i]
		}
		fmt.Println("Sorted by Post Time.")
	case 2:
		for i := 1; i < contentCount; i++ {
			key := contents[i]
			j := i - 1
			keyEng := key.Likes + key.Comments + key.Shares
			for j >= 0 && (contents[j].Likes+contents[j].Comments+contents[j].Shares) < keyEng {
				contents[j+1] = contents[j]
				j--
			}
			contents[j+1] = key
		}
		fmt.Println("Sorted by Engagement.")
	default:
		fmt.Println("Invalid option.")
	}
	return contents
}

func editContent(contents [maxContents]Content, contentCount int) [maxContents]Content {
	// Initial State:
	// - contents: list of all content
	// - contentCount: number of stored contents
	// - user provides ID of content to edit
	//
	// Final State:
	// - Selected content field is updated
	// - No new content is added or deleted

	var id int
	fmt.Print("Enter content ID to edit: ")
	fmt.Scan(&id)

	for i := 0; i < contentCount; i++ {
		if contents[i].ID == id {
			fmt.Println("1.Title 2.Platform 3.Category 4.PostTime 5.Likes 6.Comments 7.Shares 8.Cancel")
			fmt.Print("Choose field to edit: ")
			var e int
			fmt.Scan(&e)
			switch e {
			case 1:
				fmt.Print("New Title: ")
				fmt.Scan(&contents[i].Title)
			case 2:
				fmt.Print("New Platform: ")
				fmt.Scan(&contents[i].Platform)
			case 3:
				fmt.Print("New Category: ")
				fmt.Scan(&contents[i].Category)
			case 4:
				var d, h string
				fmt.Print("New Post Date & Time (YYYY-MM-DD HH:MM): ")
				fmt.Scan(&d, &h)
				contents[i].PostTime = d + " " + h
			case 5:
				fmt.Print("New Likes: ")
				fmt.Scan(&contents[i].Likes)
			case 6:
				fmt.Print("New Comments: ")
				fmt.Scan(&contents[i].Comments)
			case 7:
				fmt.Print("New Shares: ")
				fmt.Scan(&contents[i].Shares)
			case 8:
				fmt.Println("Edit cancelled.")
				return contents
			default:
				fmt.Println("Invalid choice.")
			}
			fmt.Println("Content updated.")
			break
		}
	}
	return contents
}

func deleteContent(contents [maxContents]Content, contentCount int) ([maxContents]Content, int) {
	// Initial State:
	// - contents: full list of contents
	// - contentCount: number of active contents
	// - user provides ID of content to delete
	//
	// Final State:
	// - Matching content is removed
	// - contentCount is decreased by 1

	var id int
	fmt.Print("Enter content ID to delete: ")
	fmt.Scan(&id)

	for i := 0; i < contentCount; i++ {
		if contents[i].ID == id {
			for j := i; j < contentCount-1; j++ {
				contents[j] = contents[j+1]
			}
			fmt.Println("Content deleted.")
			return contents, contentCount - 1
		}
	}
	fmt.Println("Content ID not found.")
	return contents, contentCount
}

func showTopContentByPeriod(contents [maxContents]Content, contentCount int) {
	// Initial State:
	// - contents: all added content entries
	// - contentCount: number of contents
	// - user inputs a period in format YYYY-MM
	//
	// Final State:
	// - Content with highest engagement for that period is displayed
	// - No data is changed or deleted

	if contentCount == 0 {
		fmt.Println("No content to analyze.")
		return
	}

	var period string
	fmt.Print("Enter period (YYYY-MM): ")
	fmt.Scan(&period)

	topIdx := -1
	topEng := -1

	for i := 0; i < contentCount; i++ {
		match := len(contents[i].PostTime) >= len(period) && contents[i].PostTime[:len(period)] == period
		if match {
			eng := contents[i].Likes + contents[i].Comments + contents[i].Shares
			if eng > topEng {
				topEng = eng
				topIdx = i
			}
		}
	}

	if topIdx >= 0 {
		c := contents[topIdx]
		fmt.Println("Top Content in", period)
		fmt.Printf("ID:%d | %s | %s | %s | %s | %d | %d | %d | Eng:%d\n",
			c.ID, c.Title, c.Platform, c.Category, c.PostTime,
			c.Likes, c.Comments, c.Shares, topEng)
	} else {
		fmt.Println("No content found for this period.")
	}
}
