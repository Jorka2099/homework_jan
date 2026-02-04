package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Jorka2099/homework_jan/internal/service"
)

// GetGTML обрабатывает начальную страницу, возвращает форму html из файла index.html
func GetHTML(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Failed to open index.html"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

// HandleUpload - хендлер для эндпонта /upload парсит загружаемый файл из формы html
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Запрос на /upload получен")

	err := r.ParseMultipartForm(10 << 20) // Максимум 10 МБ
	if err != nil {
		http.Error(w, "Failed to parse html form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file"+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	f, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	ConvertedText := service.TextOrMorse(string(f))

	fileName := time.Now().UTC().Format("2006-01-02") + " converted text" + filepath.Ext(header.Filename)
	err = os.WriteFile(fileName, []byte(ConvertedText), 0o644)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ConvertedText))
}
