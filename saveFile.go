package libs

import (
	"io"
	"mime/multipart"
	"os"
)

/*
SaveFile save upload files
*/
func SaveFile(file *multipart.FileHeader, dst string) (err error) {
	src, err := file.Open()
	if err != nil {
		Err.Println(err)
		return
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		Critical.Println(err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		Critical.Println(err)
		return
	}
	return
}
