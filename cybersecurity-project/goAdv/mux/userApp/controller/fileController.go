package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"userPassport/model"
	"userPassport/service"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type FileController struct {
	fileService *service.FileService
	userService *service.UserService
}

func NewFileController(fileService *service.FileService, userService *service.UserService) *FileController {
	return &FileController{
		fileService: fileService,
		userService: userService,
	}
}

func (fc *FileController) RegisterFileRoutes(router *mux.Router) {
	fc.fileService.Logger.Info().Msg("Registered File Routes")
	subRoute := router.PathPrefix("/files").Subrouter()
	subRoute.HandleFunc("/{userId}", fc.getAllAccessibleFilesMetadata).Methods("GET")
	subRoute.HandleFunc("/{userId}/write/{id}", fc.writeToFile).Methods("POST")
	subRoute.HandleFunc("/{userId}/read/{id}", fc.readFromFile).Methods("GET")
}

type accessibleFileList struct {
	ReadList  []model.File `json:"ReadList"`
	WriteList []model.File `json:"WriteList"`
}

func (cc *FileController) getAllAccessibleFilesMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.fileService.GetCoursesCount()))

	cc.fileService.Logger.Info().Msg("Get All Files")
	params := mux.Vars(r)
	userId, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID invalid!")
		return
	}
	var user model.User
	cc.userService.GetUserById(&user, userId)
	readAccessibleFiles := cc.fileService.GetBLPAndBIBAAccessibleFiles(user, "r")
	writeAccessibleFiles := cc.fileService.GetBLPAndBIBAAccessibleFiles(user, "w")
	//var files []model.File
	//cc.fileService.GetAllFilesMetadata(&files, limit, offset)
	accessibleFiles := accessibleFileList{ReadList: readAccessibleFiles, WriteList: writeAccessibleFiles}
	json.NewEncoder(w).Encode(accessibleFiles)
}

type filedata struct {
	Data string `json:"Data"`
}

func (fc *FileController) writeToFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	userId, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID invalid!")
		return
	}
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var file model.File
	fc.fileService.GetFileMetadataById(&file, id, []string{})
	var user model.User
	fc.userService.GetUserById(&user, userId)
	valid := fc.fileService.CheckIfFileIsAccessibleToUser(user, file, "w")
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Permission denied")
		return
	}

	//fmt.Println("data from api-", filedata1.Data)
	var filedata1 filedata
	json.NewDecoder(r.Body).Decode(&filedata1)

	err = fc.fileService.WriteToFile(file.FileName, filedata1.Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to write data")
		return
	}
	json.NewEncoder(w).Encode("Done with write operation")
}

func (fc *FileController) readFromFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Course-Count", strconv.Itoa(cc.courseService.GetCoursesCount()))
	params := mux.Vars(r)
	userId, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID invalid!")
		return
	}
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var file model.File
	fc.fileService.GetFileMetadataById(&file, id, []string{})
	var user model.User
	fc.userService.GetUserById(&user, userId)
	valid := fc.fileService.CheckIfFileIsAccessibleToUser(user, file, "r")
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Permission denied")
		return
	}
	fmt.Println(file.FileName)
	data, err := fc.fileService.ReadFromFile(file.FileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unable to read data")
		return
	}
	json.NewEncoder(w).Encode(data)
}
