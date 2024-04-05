package controller

import (
	"fmt"
	ioutil "io"
	"log"
	"net/http"
	sqlserver "src/database"
	models "src/models"
	utils "src/utils"

	"github.com/gorilla/mux"
)

func getArticlesFromDb() []models.Article {
	var Articles []models.Article
	sqlserver.Database.Find(&Articles)
	return Articles
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, getArticlesFromDb())
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var article models.Article

	result := sqlserver.Database.Where("Id = ?", id).First(&article)
	if result.Error == nil {
		utils.JsonResponse(w, article)
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle models.Article

	utils.JsonDeserialize(reqBody, &newArticle)

	result := sqlserver.Database.Create(&newArticle)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Article has been created",
	})
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var deletedArticle models.Article
	result := sqlserver.Database.Where("Id = ?", id).Delete(deletedArticle)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Article has been deleted",
	})
}

func HandleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/articles", allArticles).Methods("GET")
	myrouter.HandleFunc("/article/{id}", getArticle).Methods("GET")
	myrouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myrouter.HandleFunc("/article", createArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myrouter))
}
