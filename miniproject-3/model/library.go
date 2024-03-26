package model

import "gorm.io/gorm"

type Library struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

var ListBook []Library

func (lb *Library) Create(db *gorm.DB) error {
	err := db.Model(Library{}).Create(&lb).Error
	if err != nil {
		return err
	}

	return nil
}

func (lb *Library) GetById(db *gorm.DB, id uint) (Library, error) {
	res := Library{}

	err := db.Model(Library{}).Where("id = ?", id).Take(&res).Error
	if err != nil {
		return Library{}, err
	}

	return res, nil
}

func (lb *Library) GetAll(db *gorm.DB) ([]Library, error) {
	res := []Library{}

	err := db.Model(Library{}).Find(&res).Error
	if err != nil {
		return []Library{}, err
	}

	return res, nil
}

func (lb *Library) UpdateOneByID(db *gorm.DB, id uint) error {
	//err := db.Model(Library{}).Where("id = ?", lb.Model.ID).Updates(&lb).Error
	err := db.Model(Library{}).
		Select("isbn", "penulis", "tahun", "judul", "gambar", "stok").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"isbn":    lb.ISBN,
			"penulis": lb.Penulis,
			"tahun":   lb.Tahun,
			"judul":   lb.Judul,
			"gambar":  lb.Gambar,
			"stok":    lb.Stok,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (lb *Library) DeleteByID(db *gorm.DB, id uint) error {
	err := db.Model(Library{}).Where("id = ?", id).Delete(&lb).Error
	if err != nil {
		return err
	}
	return nil
}
