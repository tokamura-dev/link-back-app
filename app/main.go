package main

import "link-back-app/handler"

func main() {
	router := handler.GetApiRouter()
	router.Run(":3001")
}
