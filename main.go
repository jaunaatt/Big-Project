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

var contents [maxContents]Content
var contentCount = 0
var nextID = 1

func main() {
	for {
		fmt.Println("\n=== Content Manager ===")
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
			addContent()
		case 2:
			showAllContents()
		case 3:
			sequentialSearch()
		case 4:
			binarySearch()
		case 5:
			sortContent()
		case 6:
			editContent()
		case 7:
			deleteContent()
		case 8:
			showTopContentByPeriod()
		case 9:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addContent() {
	if contentCount >= maxContents {
		fmt.Println("Sorry, your content list is full.")
		return
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

	contentCount++
	nextID++

	fmt.Println("Content added successfully!")
}

func showAllContents() {
	if contentCount == 0 {
		fmt.Println("No content available.")
		return
	}

	fmt.Println("\n--- All Content Ideas ---")
	for i := 0; i < contentCount; i++ {
		c := contents[i]
		eng := c.Likes + c.Comments + c.Shares
		fmt.Printf("ID:%d |Title:%s |Platform:%s |Category:%s |PostTime:%s |Likes:%d |Comments:%d |Shares:%d |Engagement:%d\n",
			c.ID, c.Title, c.Platform, c.Category, c.PostTime,
			c.Likes, c.Comments, c.Shares, eng)
	}
}

func sequentialSearch() {
	if contentCount == 0 {
		fmt.Println("No content to search.")
		return
	}

	var kw string
	fmt.Print("Enter keyword (title/category): ")
	fmt.Scan(&kw)

	found := false
	for i := 0; i < contentCount; i++ {
		if contents[i].Title == kw || contents[i].Category == kw {
			c := contents[i]
			fmt.Printf("ID:%d |%s |%s |%s |%s |%d|%d|%d\n",
				c.ID, c.Title, c.Platform, c.Category, c.PostTime,
				c.Likes, c.Comments, c.Shares)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching content found.")
	}
}

func binarySearch() {
	if contentCount == 0 {
		fmt.Println("No content available.")
		return
	}

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
			fmt.Printf("Found: ID:%d |%s |%s |%s |%s |%d|%d|%d\n",
				c.ID, c.Title, c.Platform, c.Category, c.PostTime,
				c.Likes, c.Comments, c.Shares)
			return
		} else if contents[mid].Title < kw {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Content not found.")
}

func sortContent() {
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
}

func editContent() {
	if contentCount == 0 {
		fmt.Println("No content to edit.")
		return
	}

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
				fmt.Print("New Post Date & Time (YYYY-MM-DD HH:MM): ")
				var d, h string
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
				return
			default:
				fmt.Println("Invalid choice.")
			}
			fmt.Println("Content updated.")
			return
		}
	}

	fmt.Println("Content ID not found.")
}

func deleteContent() {
	if contentCount == 0 {
		fmt.Println("No content to delete.")
		return
	}

	var id int
	fmt.Print("Enter content ID to delete: ")
	fmt.Scan(&id)

	for i := 0; i < contentCount; i++ {
		if contents[i].ID == id {

			for j := i; j < contentCount-1; j++ {
				contents[j] = contents[j+1]
			}
			contentCount--
			fmt.Println("Content deleted.")
			return
		}
	}

	fmt.Println("Content ID not found.")
}

func showTopContentByPeriod() {
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

		match := true
		if len(contents[i].PostTime) < len(period) {
			match = false
		} else {
			for j := 0; j < len(period); j++ {
				if contents[i].PostTime[j] != period[j] {
					match = false
					break
				}
			}
		}
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
		fmt.Printf("ID:%d |%s |%s |%s |%s |%d|%d|%d |Eng:%d\n",
			c.ID, c.Title, c.Platform, c.Category, c.PostTime,
			c.Likes, c.Comments, c.Shares, topEng)
	} else {
		fmt.Println("No content found for this period.")
	}
}
