package repository

import (
	"database/sql"
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
	//GetCount(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	GetAllWithQueryProcessor(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error
	GetAllRows(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) (*sql.Rows, error)
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
func GroupBy(entity interface{}, statement string, group string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Model(entity).Select(statement).Group(group)
		return db, nil
	}
}
func Count(out interface{}, count *int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Model(out).Count(&count)
		return db, nil
	}
}
func OrderBy(condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Order(condition), nil
	}
}
func Model() QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Model(out), nil
	}
}
func Select(statement string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Select(statement), nil
	}
}
func Joins(condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Joins(condition), nil
	}
}
func Group(group string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Group(group)
		return db, nil
	}
}
func Join(entity interface{}, statement, condition string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Model(entity).Select(statement).Joins(condition), nil
	}
}
func Having(condition string, val int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		return db.Debug().Having(condition, val), nil
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

//func (repository *GormRepository) ExecuteGroupBy(row
func (repository *GormRepository) GetAllWithQueryProcessor(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
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

func (repository *GormRepository) GetAllRows(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) (*sql.Rows, error) {
	db := uow.DB
	var rows *sql.Rows
	var err error
	for _, queryProcessor := range queryProcessors {
		db, err = queryProcessor(db, out)
		if err != nil {
			return nil, err
		}
	}
	//db, err = queryProcessors(db, out)
	rows, err = db.Rows()
	if err != nil {
		return nil, err
	}

	return rows, nil
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
