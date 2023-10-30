package tools

import (
	"fmt"
	"io"
	"log"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func ReadFile(path string) []byte {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return nil
	}

	fileinfo, _ := file.Stat()
	ln := fileinfo.Size()

	defer file.Close()

	// baca file
	var text = make([]byte, ln)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return nil
	}

	fmt.Println("==> file berhasil dibaca")
	return text
}

func CreateFile(path string) {
	myfile, e := os.Create(path)
	if e != nil {
		log.Fatal(e)
	}
	log.Println(myfile)
	myfile.Close()
}

func WriteFile(path string, str []string) {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	co := len(str)
	for i := 0; i < co; i++ {
		_, err = file.WriteString(str[i])
	}
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}
