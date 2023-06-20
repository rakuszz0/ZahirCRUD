package repositories

import (
	"CRUD/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	FindContacts() ([]models.Contact, error)
	GetContact(ID int) (models.Contact, error)
	CreateContact(contact models.Contact) (models.Contact, error)
	UpdateContact(contact models.Contact) (models.Contact, error)
	DeleteContact(contact models.Contact, ID int) (models.Contact, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryContact(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Find(&contacts).Error

	return contacts, err
}

func (r *repository) GetContact(ID int) (models.Contact, error) {
	var contact models.Contact
	err := r.db.First(&contact, ID).Error

	return contact, err
}

func (r *repository) CreateContact(contact models.Contact) (models.Contact, error) {
	err := r.db.Create(&contact).Error

	return contact, err
}

func (r *repository) UpdateContact(contact models.Contact) (models.Contact, error) {
	err := r.db.Save(&contact).Error

	return contact, err
}

func (r *repository) DeleteContact(contact models.Contact, ID int) (models.Contact, error) {
	err := r.db.Delete(&contact, ID).Scan(&contact).Error

	return contact, err
}
