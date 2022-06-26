package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func scanDirectory2(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	myErr := fmt.Errorf("permission denied")
	panic(myErr)

	for _, file := range files {

		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}

	}

}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	fmt.Println("2")
	err, ok := p.(error)
	if ok {
		fmt.Println("3")
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	/*files, err := ioutil.ReadDir("d:\\my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}*/

	/*err := scanDirectory("d:\\my_directory")
	if err != nil {
		log.Fatal(err)
	}*/

	defer reportPanic()
	panic("some other issue")
	scanDirectory2("C:\\Windows\\System32\\drivers\\etc")
	fmt.Println("Existing program")
}
