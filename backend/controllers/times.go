package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parkrun/backend/models"
	"strconv"
)

func NewTimesService(trs *models.TimeRecordService) *TimeRecords {
	return &TimeRecords{
		trs: trs,
	}
}

type TimeRecords struct {
	trs *models.TimeRecordService
}

type TimeRecord struct {
	Time      string
	BarcodeID string
}

func (t *TimeRecords) ViewTimes(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r.URL.Query()["date"])
	hasDateParam := r.URL.Query()["date"] != nil
	var times []models.TimeRecord
	var err error
	if hasDateParam {
		fmt.Println("bydate")
		times, err = t.trs.TimeRecodsByDate(r.URL.Query()["date"][0])
	} else {
		fmt.Println("all")
		times, err = t.trs.AllTimeRecords()
	}

	if err != nil {
		panic(err)
	}
	message := ""

	for _, v := range times {
		fmt.Printf("%T\n", v.BarcodeID)
		message = message + strconv.FormatInt(int64(v.BarcodeID), 10) + "  " + strconv.FormatInt(int64(v.Time), 10) + "\n"
	}
	fmt.Fprint(w, message)
}

func (t *TimeRecords) AddTimes(w http.ResponseWriter, r *http.Request) {

	var TimeRecords []TimeRecord
	err := json.NewDecoder(r.Body).Decode(&TimeRecords)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(TimeRecords)

	for _, tr := range TimeRecords {
		bcid, _ := strconv.ParseUint(tr.BarcodeID, 10, 32)
		tr, _ := strconv.ParseUint(tr.Time, 10, 32)
		timeRecord := models.TimeRecord{
			BarcodeID: uint(bcid),
			Time:      uint(tr),
		}
		err = t.trs.Create(&timeRecord)

	}
	if err == nil {
		fmt.Fprint(w, "all times uploaded succsessfully")
	} else {
		fmt.Fprint(w, "unsuccsessful please try again")
	}
}
