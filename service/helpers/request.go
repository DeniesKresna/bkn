package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/serror"
)

func GetPaginationDataFromService(ir models.IndexRequest) models.PaginationData {
	var perPage, page, offset int
	pageReq, err := strconv.Atoi(ir.Page)
	if err != nil {
		page = 1
	} else {
		page = pageReq
		if page == 0 {
			page = 1
		}
	}
	perPageReq, err := strconv.Atoi(ir.PerPage)
	if err != nil {
		perPage = 15
	} else {
		perPage = perPageReq
		if perPage == 0 {
			perPage = 15
		}
	}
	sort := ir.Sort
	offset = 0

	offset = (page - 1) * perPage
	data := models.PaginationData{
		Page:   page,
		Limit:  perPage,
		Offset: offset,
		Sort:   sort,
	}

	return data
}

func PerformHTTPRequest(method, url string, headers map[string]string, bodies ...string) (res []byte, dump models.HttpDump, errx serror.SError) {
	var body = ""
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		if len(bodies) > 0 {
			body = bodies[0]
		}
		if !IsValidJSON(body) {
			err := errors.New("the body payload was invalid json format")
			errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Validate Body")
			return
		}
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Create Http New Request")
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Do HttpDump Request")
		return
	}
	dump.ReqDump = string(reqDump)

	response, err := client.Do(req)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Create Do Http Request")
		return
	}
	defer response.Body.Close()

	resDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Do HttpDump Response")
		return
	}
	dump.ResDump = string(resDump)

	res, err = ioutil.ReadAll(response.Body)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, "While Read Http Response")
		return
	}

	return
}

func IsValidJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}
