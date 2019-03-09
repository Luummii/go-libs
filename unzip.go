package libs

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/*
Unzip will decompress a zip archive, moving all files and folders
within the zip file (parameter 1) to an output directory (parameter 2).
*/
func Unzip(in string, out string) (filenames []string, err error) {
	r, err := zip.OpenReader(in)
	if err != nil {
		Err.Println(err)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			Err.Println(err)
			return filenames, err
		}

		// Store filename/path for returning and using later on
		fpath := filepath.Join(out, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(out)+string(os.PathSeparator)) {
			Err.Println("zipslip error")
			return filenames, err
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				Critical.Println(err)
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				Critical.Println(err)
				return filenames, err
			}
			_, err = io.Copy(outFile, rc)
			if err != nil {
				Critical.Println(err)
				return filenames, err
			}
			outFile.Close()
		}
		rc.Close()
	}
	return
}
