package main

import (
    "Seinarukiro_XYZ/database"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strconv"
    "time"
)

func HandleLatestPost() {
    latestPost, err := database.GetLatestPost()
    if err != nil {
        fmt.Println("Error getting latest post:", err)
        return
    }
    fmt.Println("Latest post:", latestPost)

    err = createMarkdownFile(latestPost)
    if err != nil {
        fmt.Println("Error creating Markdown file:", err)
        return
    }
    fmt.Println("Markdown file created successfully.")
}

func createMarkdownFile(post database.Post) error {
    contentDir := "../frontend/content/blog"

    if _, err := os.Stat(contentDir); os.IsNotExist(err) {
        err := os.MkdirAll(contentDir, os.ModePerm)
        if err != nil {
            return err
        }
    }

    markdownContent := fmt.Sprintf(`---
title: "%s"
date: %s
---

%s
`, post.Title, post.Date.Format("2006-01-02"), post.Description)

    filename := strconv.FormatUint(uint64(post.ID), 10) + ".md"
    filePath := filepath.Join(contentDir, filename)

    err := ioutil.WriteFile(filePath, []byte(markdownContent), 0644)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    HandleLatestPost()
}
