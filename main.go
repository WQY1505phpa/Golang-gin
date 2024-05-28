package main

import "wqy/router"

func main(){
	r := router.Router()
	r.Run("127.0.0.1:8080")
}

