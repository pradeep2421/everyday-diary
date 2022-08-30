package main

import "backend/Routers"
func main(){
	r := Routers.SetupRouter()
	r.Run()
}