package main

import (
	orm "new_village/src/db"
	router2 "new_village/src/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router2.InitRouter()
	router.Run("8080")
}
