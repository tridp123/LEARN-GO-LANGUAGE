package models

import (
	"entities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MobileModel struct {
	Db *mgo.Database
}

func (mobileModel MobileModel) FindAll() ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Find(id string) (entities.Mobile, error) {
	var mobile entities.Mobile
	err := mobileModel.Db.C("mobile").FindId(bson.ObjectIdHex(id)).One(&mobile)
	return mobile, err
}

func (mobileModel MobileModel) Search1(keyword string) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"name": bson.RegEx{
			Pattern: "^" + keyword,
			Options: "i",
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Search2(keyword string) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"name": bson.RegEx{
			Pattern: keyword + "$",
			Options: "i",
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Search3(keyword string) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"name": bson.RegEx{
			Pattern: keyword,
			Options: "i",
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Sort1() ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Sort("price").All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Sort2() ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Sort("-price").All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Limit1(n int) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Limit(n).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Limit2(n int) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Sort("-price").Limit(n).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Create(mobile *entities.Mobile) error {
	return mobileModel.Db.C("mobile").Insert(&mobile)
}
func (mobileModel MobileModel) Update(mobile *entities.Mobile) error {
	return mobileModel.Db.C("mobile").Update(mobile.Id, &mobile)
}
func (mobileModel MobileModel) Delete(id string) error {
	return mobileModel.Db.C("mobile").RemoveId(bson.ObjectIdHex(id))
}
