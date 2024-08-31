// Code generated by Validator v0.2.5. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *BasePageInfo) IsValid() error {
	if p.PageNum < int32(1) {
		return fmt.Errorf("field PageNum ge rule failed, current value: %v", p.PageNum)
	}
	if p.PageSize < int32(1) {
		return fmt.Errorf("field PageSize ge rule failed, current value: %v", p.PageSize)
	}
	return nil
}
