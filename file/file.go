package file

import "io/ioutil"

func Search_file_recursively(path string) (s []string) {
	file_info, _ := ioutil.ReadDir(path)
	for _, v := range file_info {
		if !v.IsDir() {
			s = append(s, path+"\\"+v.Name())
		}
	}
	for _, v := range file_info {
		if v.IsDir() {
			s = append(s, Search_file_recursively(path+"\\"+v.Name())...)
		}
	}
	return
}
