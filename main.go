package main

import "blog-backend/initialize"

func init() {
	if err := initialize.InitProject(); err != nil {
		panic(err)
	}
}

func main() {

}
