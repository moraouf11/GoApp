package main

import (
	controller "src/controller"
	sqlserver "src/database"
)

// func handleRequests() {
// 	myrouter := mux.NewRouter().StrictSlash(false)
// 	myrouter.HandleFunc("/articles", allArticles).Methods("GET")
// 	myrouter.HandleFunc("/article/{id}", getArticle).Methods("GET")
// 	myrouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
// 	myrouter.HandleFunc("/article", createArticle).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8080", myrouter))
// }

func main() {
	sqlserver.Init()
	controller.HandleRequests()
}
