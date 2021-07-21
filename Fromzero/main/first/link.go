package main

import "package/Link"

func main() {
	Head := Link.NewEmpty(2, nil)
	Link.AssignmentEmpty(Head, 1, 2)
	Link.PrintAllList(Head)
}
