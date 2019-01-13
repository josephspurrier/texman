package texman_test

import (
	"fmt"
	"log"

	"github.com/josephspurrier/texman"
)

func Example() {
	s := texman.NewFile("testdata/base.txt")
	err := s.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("---Before---")
	fmt.Println(s.String())
	fmt.Println("---After---")

	err = s.Overwrite(2, 3, "ack")
	if err != nil {
		log.Fatalln(err)
	}

	err = s.Insert(3, 5, "-Pink")
	if err != nil {
		log.Fatalln(err)
	}

	err = s.DeleteLine(4)
	if err != nil {
		log.Fatalln(err)
	}

	// Output:
	// ---Before---
	// 日本語
	// Blue
	//  Red
	//   Green
	// Yellow
	// ---After---
	// 日本語
	// Black
	//  Red-Pink
	// Yellow
	fmt.Println(s.String())
}
