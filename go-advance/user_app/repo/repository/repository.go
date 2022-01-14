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
	ClearAssociation(uow *UnitOfWork, entity interface{}, tableName string) error
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}) error
	DeleteAssociation(uow *UnitOfWork, entity interface{}, tableName string, associationToBeDeleted interface{}) error
	HardDelete(uow *UnitOfWork, entity interface{}) error
	GetFirst(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	GetAllWithQueryProcessor(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
	GetCount(uow *UnitOfWork, model interface{}, out interface{}, condition string, value interface{}) error
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
	if err := db.Debug().First(out).Error; err != nil {
		return err
	}
	return nil
}
func (repository *GormRepository) DeleteAssociation(uow *UnitOfWork, entity interface{}, tableName string, associationToBeDeleted interface{}) error {
	db := uow.DB
	return db.Debug().Model(entity).Association(tableName).Delete(associationToBeDeleted).Error
}
func PreloadAssociations(preloadAssociations []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {

		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}
func Limit(val int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Limit(val), nil
	}
}
func Offset(val int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Offset(val), nil
	}
}

func (repository *GormRepository) GetAllWithQueryProcessor(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
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
func (repository *GormRepository) GetCount(uow *UnitOfWork, model interface{}, out interface{}, condition string, value interface{}) error {
	db := uow.DB
	return db.Model(model).Debug().Where(condition, value).Count(out).Error
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Find(out).Error
}

func (repository *GormRepository) ClearAssociation(uow *UnitOfWork, entity interface{}, tableName string) error {
	db := uow.DB
	return db.Model(entity).Debug().Association(tableName).Clear().Error
}

// GetAllForTenant returns all objects of specified tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Debug().Where("id = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	//fmt.Printf("Here in add func", entity)
	return uow.DB.Debug().Create(entity).Error
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	fmt.Printf("Here in update %T", entity)
	return uow.DB.Model(entity).Debug().Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Delete(entity).Error
}

func (repository *GormRepository) HardDelete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Unscoped().Delete(entity).Error
}
