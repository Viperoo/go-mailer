package main

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	//	"io/ioutil"
	"os"
)

func addMailing(db gorm.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter path to template: ")
	filename, _ := reader.ReadString('\n')

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(f)

	fmt.Print("Enter subject mailing: ")
	subject, _ := reader.ReadString('\n')
	isEmpty(subject)

	fmt.Print("Enter from mailing (name <mail@localhost>): ")
	from, _ := reader.ReadString('\n')
	isEmpty(from)

	fmt.Printf(subject)
	fmt.Printf(from)

}
