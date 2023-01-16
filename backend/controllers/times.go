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
	Time      uint
	BarcodeID uint
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
		timeRecord := models.TimeRecord{
			BarcodeID: tr.BarcodeID,
			Time:      tr.Time,
		}
		err = t.trs.Create(&timeRecord)
		if err != nil {
			fmt.Println("Time Not added: ", strconv.FormatInt(int64(tr.BarcodeID), 10))
		}
		fmt.Fprint(w, "all times uploaded succsessfully")

	}

}
