package service

import (
	"os"
	"userPassport/aes"
	"userPassport/model"
	"userPassport/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

type FileService struct {
	Repo   repository.Repository
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func NewFileService(Repo repository.Repository, db *gorm.DB, logger *zerolog.Logger) *FileService {
	return &FileService{
		Repo:   Repo,
		DB:     db,
		Logger: logger,
	}
}

func (s *FileService) GetAllFilesMetadata(out *[]model.File, limit int, offset int) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	var queryp []repository.QueryProcessor
	var count int
	if limit != 0 {
		queryp = append(queryp, repository.Paginate(limit, offset, &count))
	}
	s.Logger.Info().Int("count", count)
	err := s.Repo.GetAll(unit, out, queryp)
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Msg("Get all files metadata")
	unit.Commit()
	return nil
}

func (s *FileService) GetFileMetadataById(out *model.File, id uuid.UUID, preloadAssociations []string) error {
	unit := repository.NewUnitOfWork(s.DB, true)
	err := s.Repo.Get(unit, out, id, preloadAssociations, "id")
	if err != nil {
		unit.Complete()
		return err
	}
	s.Logger.Info().Interface("file ID", id).Msg("Get file by file ID")
	unit.Commit()
	return nil
}

func (s *FileService) WriteToFile(filename string, data string) error {
	//fmt.Println("data from writeToFile", data)
	file, err := os.OpenFile("./data/"+filename+".txt", os.O_WRONLY, 0777)
	fileContent, _ := s.ReadFromFile(filename)
	//fmt.Println(fileContent, "---")
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer file.Close()
	if _, err := file.Write(aes.Encrypt([]byte(fileContent+data), "hello")); err != nil {
		//fmt.Println(err)
		return err
	}
	s.Logger.Info().Interface("file Name", filename).Msg("Write to file")
	return nil
}

func (s *FileService) ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile("./data/" + filename + ".txt")
	//fmt.Println(data)
	if err != nil {
		//fmt.Println(err)
		return "", err
	}
	if string(data) == "" {
		//fmt.Println(err)
		return "", nil
	}
	//fmt.Println(string(aes.Decrypt(data, "hello")))
	s.Logger.Info().Interface("file Name", filename).Msg("Read from file")
	return string(aes.Decrypt(data, "hello")), nil
}

func (s *FileService) GetBLPAndBIBAAccessibleFiles(user model.User, operation string) []model.File {
	accessibleFiles := []model.File{}
	var files []model.File
	s.GetAllFilesMetadata(&files, 0, 0)
	for _, file := range files {
		blpCheck := !((file.LevelBell > user.LevelBell && operation == "r") || (file.LevelBell < user.LevelBell && operation == "w"))
		bibaCheck := !((file.LevelBIBA > user.LevelBIBA && operation == "w") || (file.LevelBIBA < user.LevelBIBA && operation == "r"))
		if blpCheck && bibaCheck {
			accessibleFiles = append(accessibleFiles, file)
		}
	}
	s.Logger.Info().Interface("user ID", user.ID).Msg("Get BLP And BIBA Accessible Files for user with id")
	return accessibleFiles
}

func (s *FileService) CheckIfFileIsAccessibleToUser(user model.User, file model.File, operation string) bool {
	blpCheck := !((file.LevelBell > user.LevelBell && operation == "r") || (file.LevelBell < user.LevelBell && operation == "w"))
	bibaCheck := !((file.LevelBIBA > user.LevelBIBA && operation == "w") || (file.LevelBIBA < user.LevelBIBA && operation == "r"))
	s.Logger.Info().Interface("user ID", user.ID).Interface("operation", operation).Msg("Check if File is Accessible to user with id")
	return blpCheck && bibaCheck
}
