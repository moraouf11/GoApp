package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobKey        string
	BusinessId    int
	Description   string
	NeededWorkers int
	BidTypeId     int
	NoOfDays      int
	CountryCode   int
	IsActive      bool
	BidType       BidType
}

type BidType struct {
	gorm.Model
	Description string
}

type Bid struct {
	gorm.Model
	JobId      int
	WorkerId   int
	isAccepted bool
	Job        Job
}
