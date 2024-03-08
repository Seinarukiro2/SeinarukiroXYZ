package main

import (
    "Seinarukiro_XYZ/database"
    "fmt"
)

func HandleLatestPost() {
    latestPost, err := database.GetLatestPost()
    if err != nil {
        fmt.Println("Error getting latest post:", err)
        return
    }
    fmt.Println("Latest post:", latestPost)
}

func main() {

    HandleLatestPost()
}
