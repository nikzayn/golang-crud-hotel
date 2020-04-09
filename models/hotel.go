//Package models contains the models required by the application
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nikzayn/golang-crud-hotel/config"
)

//Hotel holding info about an hotel
type Hotel struct {
	gorm.Model
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	GoogleLocation string `json:"google_location,omitempty"`
	Address        string `json:"address,omitempty"`
	Contact        string `json:"contact,omitempty"`
}

//GetHotels returns the list of hotels available in the platform
func GetHotels(a *config.AppContext) ([]Hotel, error) {
	result := []Hotel{}
	err := a.Db.Find(&result).Error
	return result, err
}

//Create will create a hotel
func (h *Hotel) Create(a *config.AppContext) error {
	return a.Db.Create(h).Error
}

//Update will update a hotel
func (h *Hotel) Update(a *config.AppContext) error {
	return a.Db.Model(h).Updates(map[string]interface{}{
		"name":           h.Name,
		"Description":    h.Description,
		"GoogleLocation": h.GoogleLocation,
		"Address":        h.Address,
		"Contact":        h.Contact,
	}).Error
}

//Delete will delete a hotel
func (h *Hotel) Delete(a *config.AppContext) error {
	return a.Db.Delete(h).Error
}

func init() {
	config.AddModels(&Hotel{})
}
