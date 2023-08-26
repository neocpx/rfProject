package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	listernAddr string
}

type FileData struct {
	Name string `json:"file_name"`
	Info string `json:"file_information"`
	Size int64  `json:"file_size"`
}

func newServer(listenAddr string) *Server {
	return &Server{
		listernAddr: listenAddr,
	}
}

func (s *Server) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "UI error", http.StatusInternalServerError)
			return
		}
		template.Execute(w, nil)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "failed to upload", http.StatusBadRequest)
			return
		}
		defer file.Close()
		path := "uploads/" + header.Filename
		f, err := os.Create(path)
		if err != nil {
			http.Error(w, "failed to save (1)", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "failed to save (2)", http.StatusInternalServerError)
			return
		}
		cmd := exec.Command("file", path)
		stdout, err := cmd.Output()
		if err != nil {
			http.Error(w, "Failed to determine file type", http.StatusInternalServerError)
			return
		}
		os.Remove(path)
		res := FileData{
			Name: header.Filename,
			Info: string(stdout),
			Size: header.Size,
		}
		JsonRes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Failed to generate JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(JsonRes)
	})

	fmt.Println("server running!!!")
	http.ListenAndServe(s.listernAddr, r)
}
