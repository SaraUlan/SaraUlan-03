package repository

import (
	"context"
	"database/sql"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"github.com/google/uuid"
)

type ContactRepositoryImpl struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &ContactRepositoryImpl{db: db}
}

func (r *ContactRepositoryImpl) CreateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
    id := ctx.Value("ID")

    query := "INSERT INTO contacts (id, first_name, last_name, middle_name, phone_number) VALUES (?, ?, ?, ?, ?)"
    result, err := r.db.ExecContext(ctx, query, id, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber)
    if err != nil {
        return nil, err
    }

    contactID, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

    contact.ID = int(contactID)
    return contact, nil
}


func (r *ContactRepositoryImpl) GetContactByID(ctx context.Context, contactID int) (*domain.Contact, error) {
	ctx = context.WithValue(ctx, "ID", uuid.New().String())
	query := "SELECT id, first_name, last_name, middle_name, phone_number FROM contacts WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, contactID)

	var contact domain.Contact
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *ContactRepositoryImpl) CreateGroup(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	 ctx = context.WithValue(ctx, "ID", uuid.New().String())
	query := "INSERT INTO groups (name) VALUES (?)"
	result, err := r.db.ExecContext(ctx, query, group.Name)
	if err != nil {
		return nil, err
	}

	groupID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	group.ID = int(groupID)
	return group, nil
}

func (r *ContactRepositoryImpl) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	ctx = context.Background()
	query := "INSERT INTO contacts_groups (contact_id, group_id) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx,query, contactID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepositoryImpl) DeleteContact(ctx context.Context, contactID int) error {
    query := "DELETE FROM contacts WHERE id = ?"
    _, err := r.db.Exec(query, contactID)
    if err != nil {
        return err
    }

    return nil
}


func (r *ContactRepositoryImpl) GetGroupByID(ctx context.Context, groupID int) (*domain.Group, error) {
    ctx = context.WithValue(ctx, "ID", uuid.New().String())
    query := "SELECT id, name FROM groups WHERE id = ?"
    row := r.db.QueryRowContext(ctx, query, groupID)

    var group domain.Group
    err := row.Scan(&group.ID, &group.Name)
    if err != nil {
        return nil, err
    }

    return &group, nil
}


func (r *ContactRepositoryImpl) UpdateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
    ctx = context.WithValue(ctx, "ID", uuid.New().String())
    query := "UPDATE contacts SET first_name = ?, last_name = ?, middle_name = ?, phone_number = ? WHERE id = ?"
    _, err := r.db.ExecContext(ctx, query, contact.FirstName, contact.LastName, contact.MiddleName, contact.PhoneNumber, contact.ID)
    if err != nil {
        return nil, err
    }

    return contact, nil
}

func (r *ContactRepositoryImpl) CreateAndAddContactToGroup(ctx context.Context, contact *domain.Contact, groupID int) (*domain.Contact, error) {
	return contact, nil
}
