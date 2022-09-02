package main

import "github.com/jonnyorman/firead"

type FirestoreDocument struct {
	Prop1 string
	Prop2 int
}

func main() {
	firead.RunTypedIntId[FirestoreDocument]()
}
