package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/DeniesKresna/bkn/models"
)

func ExpertConditionQueryBuilder(filter models.ExpertSearch) string {
	q := " where `datas`->>'$.deleted_at' = 'null' and `datas`->>'$.active' = 2"

	filterValue := reflect.ValueOf(filter)
	filterType := filterValue.Type()
	for i := 0; i < filterType.NumField(); i++ {
		fieldValue := filterValue.Field(i).Interface()
		if fieldValue == nil {
			continue
		}
		fieldName := filterType.Field(i).Tag.Get("db")
		switch fieldName {
		case "services":
			services := fieldValue.([]string)
			if len(services) > 0 {
				for _, service := range services {
					q += fmt.Sprintf(" and '%s' member of (`datas`->>'$.available_services')", service)
				}
			}
		case "sectors":
			secs := fieldValue.([]string)
			if len(secs) > 0 {
				for _, sec := range secs {
					q += fmt.Sprintf(" and '%s' member of (`datas`->>'$.sectors')", sec)
				}
			}
		case "name":
			val := strings.TrimSpace(fieldValue.(string))
			if val != "" {
				q += fmt.Sprintf(" and LOWER(`datas`->>'$.name') like LOWER('%%%s%%')", val)
			}
		case "profession":
			val := strings.TrimSpace(fieldValue.(string))
			if val != "" {
				q += fmt.Sprintf(" and LOWER(`datas`->>'$.profession') like LOWER('%%%s%%')", val)
			}
		case "domicile":
			val := strings.TrimSpace(fieldValue.(string))
			if val != "" {
				q += fmt.Sprintf(" and `datas`->>'$.domicile' = '%s'", val)
			}
		case "education":
			childFilterValue := reflect.ValueOf(fieldValue)
			childFilterType := childFilterValue.Type()
			for i := 0; i < childFilterType.NumField(); i++ {
				childFieldValue := childFilterValue.Field(i).Interface()
				if childFieldValue == nil {
					continue
				}
				childFieldName := childFilterType.Field(i).Tag.Get("db")
				switch childFieldName {
				case "degree":
					degrees := childFieldValue.([]string)
					if len(degrees) > 0 {
						for _, degree := range degrees {
							q += fmt.Sprintf(" and '%s' member of (`datas`->>'$education.degree')", degree)
						}
					}
				case "school":
					val := strings.TrimSpace(childFieldValue.(string))
					if val != "" {
						q += fmt.Sprintf(" and LOWER(`datas`->>'$.education.school') like LOWER('%%%s%%')", val)
					}
				}
			}
		case "company_name":
			val := strings.TrimSpace(fieldValue.(string))
			if val != "" {
				q += fmt.Sprintf(" and LOWER(`datas`->>'$.experiences') like LOWER('%%%s%%')", val)
			}
		case "star":
			val := fieldValue.(int)
			q += fmt.Sprintf(" and `datas`->>'$.star' >= %d", val)
		case "recruit_expert":
			childFilterValue := reflect.ValueOf(fieldValue)
			childFilterType := childFilterValue.Type()
			for i := 0; i < childFilterType.NumField(); i++ {
				childFieldValue := childFilterValue.Field(i).Interface()
				if childFieldValue == nil {
					continue
				}
				childFieldName := childFilterType.Field(i).Tag.Get("db")
				switch childFieldName {
				case "capabilities":
					caps := childFieldValue.([]string)
					if len(caps) > 0 {
						for _, capp := range caps {
							q += fmt.Sprintf(" and '%s' member of (`datas`->>'$.service.recruit_expert.capabilities')", capp)
						}
					}
				}
			}
		}
	}
	return q
}
