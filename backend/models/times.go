package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/lib/pq"
)

type TimeRecord struct {
	gorm.Model
	BarcodeID uint
	Time      uint
}

type TimeRecordService struct {
	db *gorm.DB
}

func NewTimesService(conectionInfo string) (*TimeRecordService, error) {
	db, err := gorm.Open("postgres", conectionInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the Database")
	db.LogMode(true)
	return &TimeRecordService{db: db}, nil
}

func (us *TimeRecordService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&TimeRecord{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *TimeRecordService) Close() error {
	return us.db.Close()
}

func (us *TimeRecordService) Create(timeRecord *TimeRecord) error {

	return us.db.Create(timeRecord).Error
}

func (us *TimeRecordService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&TimeRecord{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

func (trs *TimeRecordService) AllTimeRecords() ([]TimeRecord, error) {
	var timesRecords []TimeRecord

	db := trs.db.Find(&timesRecords)

	if db.Error != nil {
		return nil, db.Error
	}

	return timesRecords, nil
}

func (trs *TimeRecordService) TimeRecodsByDate(dateStr string) ([]TimeRecord, error) {
	var timesRecords []TimeRecord
	startDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}
	endDate := startDate.Add(time.Hour * 24)
	startDateStr := startDate.Format("2006-01-02 00:00:00")
	endDateStr := endDate.Format("2006-01-02 00:00:00")

	db := trs.db.Where("created_at BETWEEN ? AND ?", startDateStr, endDateStr).Order("time ASC").Find(&timesRecords)

	if db.Error != nil {
		return nil, db.Error
	}
	return timesRecords, nil
}
