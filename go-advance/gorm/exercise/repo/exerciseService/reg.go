package exerciseservice

import (
	m "app/model"
	r "app/repository"
	"fmt"

	"github.com/jinzhu/gorm"
)

type RegionService struct {
	//uow *r.UnitOfWork
	Repo r.Repository
	DB   *gorm.DB
}

func NewRegionsService(r r.Repository, DB *gorm.DB) *RegionService {
	return &RegionService{Repo: r, DB: DB}
}

func (rs *RegionService) GetRegionsWithNoCountries(entity interface{}, statement, condition string) {
	// select region_name from regions left join countries on regions.region_id=countries.REGION_ID where country_name is null
	uow := r.NewUnitOfWork(rs.DB, true)
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model(entity)
	joins := r.Joins(condition)
	var result []m.Region
	where := r.Filter("country_name is null")
	rows, err := rs.Repo.GetAllRows(uow, entityModel, selectStatement, joins, where)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	for rows.Next() {
		r := m.Region{}
		rows.Scan(&r.RegionName)
		result = append(result, r)
	}
	for _, val := range result {
		fmt.Println("Region name with no countries -", val.RegionName)
	}
}

func (rs *RegionService) GetCountriesWithNoStates(entity interface{}, statement, condition, group string) {
	//select country_name from countries left join locations on countries.COUNTRY_ID=locations.COUNTRY_ID group by countries.country_id having count(state_province)=0
	var result []m.Country
	//rs.DB.Debug().Model(entity).Joins("left join locations on countries.COUNTRY_ID=locations.COUNTRY_ID").Group("countries.country_id").Having("count(state_province)=0").Find(&result)
	uow := r.NewUnitOfWork(rs.DB, true)
	//join := r.Join(entity, statement, condition)
	selectStatement := r.Select(statement)
	entityModel := r.Model(entity)
	joins := r.Joins(condition)
	//g := r.GroupBy(entity, statement, group)
	g := r.Group(group)
	havingClause := r.Having("count(state_province)= ?", 0)
	err := rs.Repo.GetAllWithQueryProcessor(uow, &result, entityModel, selectStatement, joins, g, havingClause)
	if err != nil {
		fmt.Println("Error in join and group by ", err)
		return
	}
	/*fmt.Println("rows ", rows)
	for rows.Next() {
		r := m.Country{}
		rows.Scan(&r.CountryName)
		//fmt.Println(r.CountryName)
	}*/
	//fmt.Println(len(result))
	for _, val := range result {
		fmt.Println("Country name with no states -", val.CountryName)
		fmt.Println("=======")
	}
}

func (rs *RegionService) GetRegionCountryState(entity interface{}, statement string, condition ...string) {
	// SELECT regions.region_name,countries.country_name,locations.STATE_PROVINCE FROM `regions` left join countries ON regions.region_id = countries.REGION_ID left join locations on countries.country_id = locations.country_id
	uow := r.NewUnitOfWork(rs.DB, true)
	//var reg []m.Region
	selectStatement := r.Select(statement)
	entityModel := r.Model(entity)
	join := r.Joins(condition[0])
	joinTwo := r.Joins(condition[1])
	//join := r.Join(entity, statement, condition[0])
	//joinTwo := r.Join(entity, statement, condition[1])
	rows, err := rs.Repo.GetAllRows(uow, entityModel, selectStatement, join, joinTwo)
	if err != nil {
		fmt.Println("Error in join ", err)
		return
	}
	//result := make(map[*m.Employee]*m.Department)
	for rows.Next() {
		reg := m.Region{}
		country := m.Country{}
		loc := m.Location{}
		rows.Scan(&reg.RegionName, &country.CountryName, &loc.StateProvince)
		fmt.Println("Region name -", reg.RegionName)
		fmt.Println("Country Name -", country.CountryName)
		fmt.Println("State name -", loc.StateProvince)
		fmt.Println("================================")
	}
}

func (rs *RegionService) CreateNewEntry(entity interface{}) {
	uow := r.NewUnitOfWork(rs.DB, false)
	err := rs.Repo.Add(uow, entity)
	if err != nil {
		fmt.Println("Error in create ")
	} else {
		uow.Commit()
	}
}

func (rs *RegionService) GetEnitiesWithCity(city string) {
	// select * from locations where city = "mumbai";
	uow := r.NewUnitOfWork(rs.DB, true)
	qp := r.Filter("city = ?", city)
	var entities []m.Location
	err := rs.Repo.GetAllWithQueryProcessor(uow, &entities, qp)
	if err != nil {
		fmt.Println("Error in getting details with city name ")
	} else {
		fmt.Println(entities)
	}
}

func (rs *RegionService) FooTable() {
	uow := r.NewUnitOfWork(rs.DB, false)
	type Foo struct {
		Name string
		Age  int
	}
	rs.DB.AutoMigrate(&Foo{})
	entry := Foo{Name: "Chinmay", Age: 21}
	er := rs.Repo.Add(uow, &entry)
	if er != nil {
		uow.Complete()
		fmt.Println("Error adding in foo table")
	}
	uow.Commit()
}
