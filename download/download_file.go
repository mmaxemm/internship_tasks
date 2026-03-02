package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
Напишите программу которая загружает файл по заданному URL и возвращает имя сохраненного файла, а также ошибку, если загрузка не удалась.

Пример входных параметров:
./loader -url "https://example.com/myfile.txt" -output "myfile.txt"
*/



// Ошибка ErrInvalidURL - возникает, если заданный URL недействительный или некорректный.
type ErrInvalidURL struct {
	URL string
}
func (e ErrInvalidURL) Error() string {
	return fmt.Sprintf("Invalid URL: %s", e.URL)

}

// Ошибка InvalidProtocolError - возникает, если если в URL указан протокол не http и https.
type InvalidProtocolError struct {
	Protocol string
}
func (e InvalidProtocolError) Error() string {
	return fmt.Sprintf("Invalid protocol: %s", e.Protocol)
}

// Ошибка ErrConnectionFailed - возникает, если не удалось установить соединение с сервером или сервер не ответил в течение определенного времени.
type ErrConnectionFailed struct {}

func (e ErrConnectionFailed) Error() string {
	return "connection failed"
}

// Ошибка ErrFileNotFound - возникает, если сервер вернул ошибку 404.
type ErrFileNotFound struct {
	FileName string
}
func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("404 no such file: %s", e.FileName)
}

// Ошибка ErrDownloadFailed - возникает, если загрузка файла не удалась по какой-то другой причине (например, сервер вернул ошибку 500).
type ErrDownloadFailed struct {
	text string
}
func (e ErrDownloadFailed) Error() string {
	return e.text
}

// Скачивание файла должно проходить используя функцию downloadFile(fileURL, fileName string) error. Программа должна реагировать соответсвующим образом на вернувшейся результат.
// В случае успеха, программа не должна ничего выводить.
func DownloadFile(fileURL, fileName string) error {
	parsedUrl, err := url.Parse(fileURL)
	if err != nil {
		return ErrInvalidURL{fileURL}
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return InvalidProtocolError{parsedUrl.Scheme}
	}

	if fileName == "" {
		path := parsedUrl.Path
		segments := strings.Split(path, "/")
		fileName = segments[len(segments)-1]
	}

	file, err := os.Create(fileName)
	if err != nil {
		return ErrDownloadFailed{"failed to create file: " + err.Error()}
	}
	defer file.Close()

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fileURL)
	if err != nil {
		return ErrConnectionFailed{}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrFileNotFound{fileName}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ErrDownloadFailed{fmt.Sprintf("server returned status %d", resp.StatusCode)}
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return ErrDownloadFailed{"failed to write file: " + err.Error()}
	}

	return nil
}
