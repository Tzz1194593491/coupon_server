// Code generated by thriftgo (0.3.15). DO NOT EDIT.

package business_code

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type BusinessCode int64

const (
	BusinessCode_SUCCESS          BusinessCode = 200
	BusinessCode_SYSTEM_ERROR     BusinessCode = 500
	BusinessCode_PARAM_MISSING    BusinessCode = 501
	BusinessCode_REPEAT_ERROR     BusinessCode = 502
	BusinessCode_SYSTEM_BUSY      BusinessCode = 503
	BusinessCode_ADD_FAIL         BusinessCode = 504
	BusinessCode_DELETE_FAIL      BusinessCode = 505
	BusinessCode_GET_FAIL         BusinessCode = 506
	BusinessCode_UPDATE_FAIL      BusinessCode = 507
	BusinessCode_PARAM_PARSE_FAIL BusinessCode = 508
)

func (p BusinessCode) String() string {
	switch p {
	case BusinessCode_SUCCESS:
		return "SUCCESS"
	case BusinessCode_SYSTEM_ERROR:
		return "SYSTEM_ERROR"
	case BusinessCode_PARAM_MISSING:
		return "PARAM_MISSING"
	case BusinessCode_REPEAT_ERROR:
		return "REPEAT_ERROR"
	case BusinessCode_SYSTEM_BUSY:
		return "SYSTEM_BUSY"
	case BusinessCode_ADD_FAIL:
		return "ADD_FAIL"
	case BusinessCode_DELETE_FAIL:
		return "DELETE_FAIL"
	case BusinessCode_GET_FAIL:
		return "GET_FAIL"
	case BusinessCode_UPDATE_FAIL:
		return "UPDATE_FAIL"
	case BusinessCode_PARAM_PARSE_FAIL:
		return "PARAM_PARSE_FAIL"
	}
	return "<UNSET>"
}

func BusinessCodeFromString(s string) (BusinessCode, error) {
	switch s {
	case "SUCCESS":
		return BusinessCode_SUCCESS, nil
	case "SYSTEM_ERROR":
		return BusinessCode_SYSTEM_ERROR, nil
	case "PARAM_MISSING":
		return BusinessCode_PARAM_MISSING, nil
	case "REPEAT_ERROR":
		return BusinessCode_REPEAT_ERROR, nil
	case "SYSTEM_BUSY":
		return BusinessCode_SYSTEM_BUSY, nil
	case "ADD_FAIL":
		return BusinessCode_ADD_FAIL, nil
	case "DELETE_FAIL":
		return BusinessCode_DELETE_FAIL, nil
	case "GET_FAIL":
		return BusinessCode_GET_FAIL, nil
	case "UPDATE_FAIL":
		return BusinessCode_UPDATE_FAIL, nil
	case "PARAM_PARSE_FAIL":
		return BusinessCode_PARAM_PARSE_FAIL, nil
	}
	return BusinessCode(0), fmt.Errorf("not a valid BusinessCode string")
}

func BusinessCodePtr(v BusinessCode) *BusinessCode { return &v }
func (p *BusinessCode) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = BusinessCode(result.Int64)
	return
}

func (p *BusinessCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}
