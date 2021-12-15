package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Repository represents generic interface for interacting with DB
type Repository interface {
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error
	GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}) error
	GetFirst(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	GetAllUsers(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
}

// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{DB: db.New(), committed: false, readOnly: true}
	}
	return &UnitOfWork{DB: db.New().Begin(), committed: false, readOnly: false}
}

// Complete marks end of unit of work
func (uow *UnitOfWork) Complete() {
	if !uow.committed && !uow.readOnly {
		uow.DB.Rollback()
	}
}

// Commit the transaction
func (uow *UnitOfWork) Commit() {
	if !uow.readOnly {
		uow.DB.Commit()
	}
	uow.committed = true
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
	return &GormRepository{}
}
func (repository *GormRepository) GetFirst(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err := db.First(out).Error; err != nil {
		return err
	}
	return nil
}
func PreloadAssociations(preloadAssociations []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {

		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				fmt.Println("Association --> ", association)
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}
func FilterAndPreloadAssociations(condition string, preloadAssociations []string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				fmt.Println("Association --> ", association)
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}
func (repository *GormRepository) GetAllUsers(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Find(out).Error; err != nil {
		return err
	}
	return nil
}

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.First(out, "id = ?", id).Error
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Find(out).Error
}

// GetAllForTenant returns all objects of specified tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Where("id = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Create(entity).Error
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	fmt.Println("Here in update")
	return uow.DB.Model(entity).Debug().Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Delete(entity).Error
}
