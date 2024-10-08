// Code generated by Validator v0.2.5. DO NOT EDIT.

package coupon_meta

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

func (p *CouponMeta) IsValid() error {
	return nil
}
func (p *GetCouponMetaReq) IsValid() error {
	if p.BaseInfo != nil {
		if err := p.BaseInfo.IsValid(); err != nil {
			return fmt.Errorf("field BaseInfo not valid, %w", err)
		}
	}
	return nil
}
func (p *GetCouponMetaResp) IsValid() error {
	if p.BaseInfo != nil {
		if err := p.BaseInfo.IsValid(); err != nil {
			return fmt.Errorf("field BaseInfo not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *AddCouponMetaReq) IsValid() error {
	if p.Type.String() == "<UNSET>" {
		return fmt.Errorf("field Type defined_only rule failed")
	}
	if p.Stock <= int32(0) {
		return fmt.Errorf("field Stock gt rule failed, current value: %v", p.Stock)
	}
	return nil
}
func (p *AddCouponMetaResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *DeleteCouponMetaReq) IsValid() error {
	return nil
}
func (p *DeleteCouponMetaResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *UpdateCouponMetaReq) IsValid() error {
	_src := []CouponMetaType{}
	var _exist bool
	for _, src := range _src {
		if p.Type == src {
			_exist = true
			break
		}
	}
	if !_exist {
		return fmt.Errorf("field Type in rule failed, current value: %v", p.Type)
	}
	if p.Stock <= int32(0) {
		return fmt.Errorf("field Stock gt rule failed, current value: %v", p.Stock)
	}
	return nil
}
func (p *UpdateCouponMetaResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetCouponValidMetaInfoReq) IsValid() error {
	return nil
}
func (p *GetCouponValidMetaInfoResp) IsValid() error {
	if p.CouponMetaInfo != nil {
		if err := p.CouponMetaInfo.IsValid(); err != nil {
			return fmt.Errorf("field CouponMetaInfo not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *TryReduceCouponStockReq) IsValid() error {
	return nil
}
func (p *TryReduceCouponStockResp) IsValid() error {
	return nil
}
