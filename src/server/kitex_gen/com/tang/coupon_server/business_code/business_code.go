// Code generated by thriftgo (0.3.15). DO NOT EDIT.

package business_code

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type BusinessCode int64

const (
	BusinessCode_SUCCESS       BusinessCode = 200
	BusinessCode_SYSTEM_ERROR  BusinessCode = 500
	BusinessCode_PARAM_MISSING BusinessCode = 501
	BusinessCode_REPEAT_ERROR  BusinessCode = 502
	BusinessCode_SYSTEM_BUSY   BusinessCode = 503
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