package main

import (
	"fmt"

	"github.com/mmaxemm/internship_tasks/download"
)

// ErrInvalidURL - можно вывести сообщение об ошибке и запросить у пользователя ввести действительный URL.
// InvalidProtocolError -  если функция вернула эту ошибку, необходимо указать какой именно неверный протокол был указан в переданном пользователем URL и запросить ввести действительный URL.
// ErrConnectionFailed - если функция вернула эту ошибку, можно вывести сообщение об ошибке и предложить пользователю повторить попытку или проверить свои настройки подключения.
// ErrFileNotFound - если функция вернула эту ошибку, можно вывести сообщение об ошибке и запросить у пользователя проверить правильность URL или наличие файла на сервере.
// ErrDownloadFailed - если функция вернула эту ошибку, можно вывести сообщение об ошибке и запросить у пользователя проверить доступность файла для скачивания.
// В случае, если в функции произошла какая-то другая ошибка, которая не была описана выше, функция может вернуть стандартную ошибку Go ErrUnexpectedEOF или создать свою собственную ошибку с пояснением причины ошибки.

func main() {
	var fileName = "tmp.html"
	var fileURL = "https://duckduckgo.com"
	err := download.DownloadFile(fileURL, fileName)
	if err == nil {
		return
	}
	switch e := err.(type) {
	case download.ErrInvalidURL:
		fmt.Println(err.Error())
	case download.InvalidProtocolError:
		fmt.Println(err.Error())
		fmt.Println("Please, enter URL with a valid protocol(http/https)")
	case download.ErrConnectionFailed:
		fmt.Println(err.Error())
		fmt.Println("Please, try again or check network connection")
	case download.ErrFileNotFound:
		fmt.Println(err.Error())
		fmt.Println("Please, check the URL or the existence of this file on the server")
	case download.ErrDownloadFailed:
		fmt.Println(err.Error())
		fmt.Println("Please, check if the file is available")
	default:
		fmt.Printf("unexpected error: %v", e)
	}


}
