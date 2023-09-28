package gate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/DeniesKresna/gohelper/utinterface"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/DeniesKresna/bkn/config"
	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/bkn/service/helpers"
	"github.com/DeniesKresna/bkn/service/usecase"
)

type Gate struct {
	ListRoutes     []models.HTTPRoute
	Validator      *validator.Validate
	MessagerLogger config.IMessagerLogger
	UserUsecase    usecase.IUserUsecase
	AuthUsecase    usecase.IAuthUsecase
}

func InitGate(
	validator *validator.Validate,
	messagerLogger config.IMessagerLogger,
	userUsecase usecase.IUserUsecase,
	authUsecase usecase.IAuthUsecase,
) *Gate {
	gate := &Gate{
		Validator:      validator,
		AuthUsecase:    authUsecase,
		UserUsecase:    userUsecase,
		MessagerLogger: messagerLogger,
	}
	gate.InitRoutes()
	return gate
}

func (c *Gate) Get(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodGet,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Post(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))

	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPost,
		Path:    fPath,
		Handler: handl,
	}

	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Put(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPut,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Delete(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodDelete,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	var response models.ApiResponse
	response.Status = statusCode
	response.Message = message
	response.Success = false
	env := utstring.GetEnv(models.AppENV, "dev")

	if statusCode < 400 {
		response.Success = true
	} else {
		errorData, ok := data.(*serror.Serror)
		if ok {
			if env != "dev" {
				if statusCode > 499 {
					go c.MessagerLogger.SendLogToMessager(errorData)
				}
			}
			if errorData.GetErrorMessage() != "" {
				utlog.Error(errorData.GetComment())
				utlog.Error(errorData.GetErrorLine())
			}
		}
	}

	var (
		err error
	)
	if data != nil {
		if utinterface.IsSlice(data) {
			data = struct {
				Data interface{} `json:"data"`
			}{
				Data: data,
			}
		} else if _, ok := data.(*serror.Serror); ok {
			if env != "dev" {
				data = struct {
					Data *string `json:"data"`
				}{}
			}
		}
	} else {
		data = struct {
			Data *string `json:"data"`
		}{}
	}

	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Success = false
		response.Message = "Tidak bisa konversi data"
	}
	response.Data = data

	jsonDt, err := json.Marshal(response)
	if err != nil {
		utlog.Error("something went wrong")
	}
	// Set CORS headers for the main request.
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDt)
}

func (c *Gate) ValidateInputs(dataset interface{}) (error, bool, map[string]string) {
	errors := make(map[string]string)
	err := c.Validator.Struct(dataset)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			return err, false, errors
		}
		datasetPtr := reflect.ValueOf(dataset)
		datasetVal := reflect.Indirect(datasetPtr)
		datasetType := datasetVal.Type()

		for _, Valerr := range err.(validator.ValidationErrors) {
			field, _ := datasetType.FieldByName(Valerr.StructField())
			name := Valerr.StructField()
			errors[name] = field.Tag.Get("valerr")
		}

		return nil, false, errors
	}
	return nil, true, errors
}

func (c *Gate) GetVar(r *http.Request, key string, def ...string) (res string) {
	vars := mux.Vars(r)
	res, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
	}
	return
}

func (c *Gate) GetInt64Var(r *http.Request, key string, def ...int64) (res int64) {
	vars := mux.Vars(r)
	resStr, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		return
	}
	return utint.Convert64FromString(resStr, res)
}

func (c *Gate) GetIntVar(r *http.Request, key string, def ...int) (res int) {
	vars := mux.Vars(r)
	resStr, ok := vars[key]
	if !ok {
		if len(def) > 0 {
			return def[0]
		}
		return
	}
	return utint.ConvertFromString(resStr, res)
}

func (c *Gate) GetQuery(r *http.Request, key string, def ...string) (res string) {
	q := r.URL.Query().Get(key)
	if q == "" {
		if len(def) > 0 {
			res = def[0]
		}
		return
	}
	res = q
	return
}

func (c *Gate) GetInt64Query(r *http.Request, key string, def ...int64) (res int64) {
	q := r.URL.Query().Get(key)
	if q == "" {
		res = def[0]
		return
	}
	res = utint.Convert64FromString(q, res)
	return
}

func (c *Gate) GetIntQuery(r *http.Request, key string, def ...int) (res int) {
	q := r.URL.Query().Get(key)
	if q == "" {
		res = def[0]
		return
	}
	res = utint.ConvertFromString(q, res)
	return
}

func (c *Gate) GetRequestPaginationData(r *http.Request) models.PaginationData {
	queryObj := models.IndexRequest{}

	queryObj.Page = c.GetQuery(r, "page", "1")
	queryObj.PerPage = c.GetQuery(r, "per_page", "15")
	queryObj.Sort = c.GetQuery(r, "sort", "")
	queryObj.SortMode = c.GetQuery(r, "mode", "asc")
	queryObj.Search = c.GetQuery(r, "search", "")
	return helpers.GetPaginationDataFromService(queryObj)
}

func (c *Gate) DecodeRequestAndValidate(r *http.Request, functionName string, dest interface{}) (errx serror.SError) {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Unmarshal Request", functionName), "Data masukan permintaan tidak sesuai")
		return
	}

	err, valid, valErrors := c.ValidateInputs(dest)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, fmt.Sprintf("%s While Validate Inputs Error", functionName), "Data masukan permintaan tidak sesuai")
		return
	}

	if !valid {
		errx = serror.NewWithCommentMessageValidation(http.StatusBadRequest, fmt.Sprintf("%s While ValidateInputs Wrong", functionName), "Data masukan permintaan tidak sesuai", valErrors)
		return
	}
	return
}

func GetRequestStringData(r *http.Request) (res string, errx serror.SError) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, "While Read Body Request", "Kesalahan Server")
		return
	}

	// Close the request body after reading
	defer r.Body.Close()

	res = string(body)
	return
}
