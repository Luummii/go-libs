package libs

import "io/ioutil"

/*
IsEmptyDir check is empty dir
*/
func IsEmptyDir(dir string) (bool, error) {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		Err.Println(err)
		return false, err
	}
	return len(entries) == 0, nil
}
