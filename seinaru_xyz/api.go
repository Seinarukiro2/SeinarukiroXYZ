package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

type BlogPost struct {
    Title       string    `json:"title"`
    Date        time.Time `json:"date"`
    Description string    `json:"description"`
    Image       string    `json:"image"`
    Content     string    `json:"content"`
}

func createPost(post BlogPost) error {
    filename := post.Date.Format("2006-01-02") + "-" + post.Title + ".md"
    filepath := filepath.Join("content", "blog", filename)

    content := "---\n"
    content += "title: " + post.Title + "\n"
    content += "date: " + post.Date.Format("2006-01-02") + "\n"
    content += "description: " + post.Description + "\n"
    content += "image: " + post.Image + "\n"
    content += "---\n\n"
    content += post.Content

    err := ioutil.WriteFile(filepath, []byte(content), 0644)
    if err != nil {
        return err
    }

    return nil
}

func deletePost(filename string) error {
    filepath := filepath.Join("content", "blog", filename)
    err := os.Remove(filepath)
    if err != nil {
        return err
    }
    return nil
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
    var post BlogPost
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = createPost(post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("filename")
    err := deletePost(filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func handleGetPostFilenames(w http.ResponseWriter, r *http.Request) {
    files, err := ioutil.ReadDir("content/blog")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var filenames []string
    for _, file := range files {
        filenames = append(filenames, file.Name())
    }

    json.NewEncoder(w).Encode(filenames)
}

func main() {
    http.HandleFunc("/api/create-post", handleCreatePost)
    http.HandleFunc("/api/delete-post", handleDeletePost)
    http.HandleFunc("/api/get-post-filenames", handleGetPostFilenames)

    http.ListenAndServe(":8080", nil)
}
