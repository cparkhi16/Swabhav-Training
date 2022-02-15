package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// Repository represents generic interface for interacting with DB
type Repository interface {
	//ConnectToDB(uow *UnitOfWork, username string, password string, host string, port string, dbName string) error
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string, primaryKeyName string) error
	GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
	//GetByTable(uow *UnitOfWork, out interface{}, queryProcessor []QueryProcessor) error
	GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}) error
	GetCount(uow *UnitOfWork, entity interface{}, count interface{}) error
}

// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// QueryProcessor allows to modify the query before it is executed
type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

func Wherefunc() QueryProcessor {
	fmt.Println("yes")
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where("name=?", "sam")
		return db, nil
	}

}

// Filter will filter the results
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}

// Order will order the results
func Order(value interface{}, reorder bool) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Order(value, reorder)
		return db, nil
	}
}

// Count
func Count() QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Count(out)
		return db, nil
	}
}

//Select
func Select(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Select(value)
		return db, nil
	}
}

//Table
func Table(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Table(value)
		return db, nil
	}
}

//Join
func Joins(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Joins(value)
		return db, nil
	}
}

//Group
func Group(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Group(value)
		return db, nil
	}
}

func Having(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Having(value)
		return db, nil
	}
}

// PreloadAssociations specified associations to be preloaded
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

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
	return &GormRepository{}
}

/*
func (repository *GormRepository) GetByTable(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor, tableName string) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			fmt.Println(db)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Table(tableName).Scan(out).Error; err != nil {
		return err
	}
	return nil
}
*/
// GetFirst gets first record matching the given criteria
func (repository *GormRepository) GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			fmt.Println(db)
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

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string, primaryKeyName string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Debug().First(out, primaryKeyName+" = ?", id).Error
}

func (repository *GormRepository) GetCount(uow *UnitOfWork, entity interface{}, count interface{}) error {
	db := uow.DB
	return db.Debug().Model(entity).Count(count).Error
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
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
	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

// GetAllForTenant returns all objects of specified tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Where("tenantID = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Create(entity).Error
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Model(entity).Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Debug().Delete(entity).Error
}
