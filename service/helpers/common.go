package helpers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/DeniesKresna/gobridge/serror"
)

func StringIsEqual(string1 string, string2 string, isCaseSensitive bool) bool {
	if !isCaseSensitive {
		string1 = strings.ToLower(string1)
		string2 = strings.ToLower(string2)
	}
	return strings.TrimSpace(string1) == strings.TrimSpace(string2)
}

func CleanMoneyToInt(param string) (res int64, errx serror.SError) {
	tempString := ""
	for _, v := range param {
		if v >= 48 && v <= 57 {
			tempString += string(v)
		}
	}
	res, err := strconv.ParseInt(tempString, 10, 64)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusBadRequest, "Cannot Parse Money")
		return
	}
	return
}

func IsHTTPURL(link string) bool {
	// Define a regular expression to match HTTP URLs
	// This regex pattern is a simplified version, and you might need to adjust it for your use case
	pattern := `^http:\/\/|^https:\/\/`
	match, _ := regexp.MatchString(pattern, link)
	return match
}
