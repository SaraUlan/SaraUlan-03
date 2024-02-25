// internal/repository/contact_repository_impl.go

package repository

import (
	"database/sql"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
)

type ContactRepositoryImpl struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &ContactRepositoryImpl{db: db}
}

func (r *ContactRepositoryImpl) CreateContact(contact *domain.Contact) (*domain.Contact, error) {
	// Реализация вставки контакта в базу данных
	query := "INSERT INTO contacts (first_name, last_name, middle_name, phone_number) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// Получаем ID вставленной записи
	contactID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	contact.ID = int(contactID)
	return contact, nil
}

func (r *ContactRepositoryImpl) GetContactByID(contactID int) (*domain.Contact, error) {
	query := "SELECT id, first_name, last_name, middle_name, phone_number FROM contacts WHERE id = ?"
	row := r.db.QueryRow(query, contactID)

	var contact domain.Contact
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *ContactRepositoryImpl) CreateGroup(group *domain.Group) (*domain.Group, error) {
	// Реализация вставки группы в базу данных
	query := "INSERT INTO groups (name) VALUES (?)"
	result, err := r.db.Exec(query, group.Name)
	if err != nil {
		return nil, err
	}

	// Получаем ID вставленной записи
	groupID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	group.ID = int(groupID)
	return group, nil
}

func (r *ContactRepositoryImpl) AddContactToGroup(contactID, groupID int) error {
	// Реализация добавления контакта в группу в базе данных
	query := "INSERT INTO contacts_groups (contact_id, group_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, contactID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepositoryImpl) DeleteContact(contactID int) error {
	// Реализация удаления контакта из базы данных
	query := "DELETE FROM contacts WHERE id = ?"
	_, err := r.db.Exec(query, contactID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepositoryImpl) GetGroupByID(groupID int) (*domain.Group, error) {
	// Реализация чтения группы из базы данных
	query := "SELECT id, name FROM groups WHERE id = ?"
	row := r.db.QueryRow(query, groupID)

	var group domain.Group
	err := row.Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *ContactRepositoryImpl) UpdateContact(contact *domain.Contact) (*domain.Contact, error) {
	// Реализация обновления контакта в базе данных
	query := "UPDATE contacts SET first_name = ?, last_name = ?, middle_name = ?, phone_number = ? WHERE id = ?"
	_, err := r.db.Exec(query, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber, contact.ID)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
