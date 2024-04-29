package controller

import (
	"fmt"
	ioutil "io"
	"log"
	"net/http"
	sqlserver "src/database"
	models "src/models"
	utils "src/utils"

	//repos "src/repositories"

	"github.com/gorilla/mux"
)

func getJobsFromDb() []models.Job {
	var Jobs []models.Job
	sqlserver.Database.Find(&Jobs) //.Where()
	return Jobs
}

func allJobs(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, getJobsFromDb())
}

func getJobWithBids(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["jobKey"]

	//repos.GetByID()
	var bids models.Bid
	//db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)

	//result := sqlserver.Database.Where("Id = ?", id).First(&job)
	result := sqlserver.Database.Joins("Job", sqlserver.Database.Where(&models.Job{JobKey: id})).Find(&bids)
	if result.Error == nil {
		utils.JsonResponse(w, bids)
	}
}

func createJob(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newJob models.Job

	utils.JsonDeserialize(reqBody, &newJob)

	result := sqlserver.Database.Create(&newJob)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Job has been created",
	})
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result := sqlserver.Database.Model(&models.Job{}).Where("Id = ?", id).Update("IsActive", false)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Job has been deleted",
	})
}

func HandleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/jobs", allJobs).Methods("GET")
	myrouter.HandleFunc("/jobs/{id}", getJobWithBids).Methods("GET")
	myrouter.HandleFunc("/jobs/{id}", deleteJob).Methods("DELETE")
	myrouter.HandleFunc("/jobs", createJob).Methods("POST")
	log.Fatal(http.ListenAndServe(":8085", myrouter))
}
