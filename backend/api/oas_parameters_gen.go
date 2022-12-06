// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// PostcodesBySettlementGetParams is parameters of GET /postcodes_by_settlement operation.
type PostcodesBySettlementGetParams struct {
	// Array of filtered type of post offices.
	Type []PostOfficeType
}

func unpackPostcodesBySettlementGetParams(packed map[string]any) (params PostcodesBySettlementGetParams) {
	if v, ok := packed["type"]; ok {
		params.Type = v.([]PostOfficeType)
	}
	return params
}

func decodePostcodesBySettlementGetParams(args [0]string, r *http.Request) (params PostcodesBySettlementGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: type.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotTypeVal PostOfficeType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotTypeVal = PostOfficeType(c)
						return nil
					}(); err != nil {
						return err
					}
					params.Type = append(params.Type, paramsDotTypeVal)
					return nil
				})
			}); err != nil {
				return params, errors.Wrap(err, "query: type: parse")
			}
			if err := func() error {
				if err := (validate.Array{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
				}).ValidateLength(len(params.Type)); err != nil {
					return errors.Wrap(err, "array")
				}
				var failures []validate.FieldError
				for i, elem := range params.Type {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: type: invalid")
			}
		}
	}
	return params, nil
}

// SendingFilterGetParams is parameters of GET /sending_filter operation.
type SendingFilterGetParams struct {
	// Current page.
	Page Page
	// The number of items displayed on the page.
	ElemsOnPage ElementsOnPage
	// UUID or a part of it.
	OrderID OptString
	// Array of filtered type of sending.
	Type []SendingType
	// Array of filtered status of sending.
	Status []SendingStatus
	// Sending registration date range.
	Date OptSendingFilterGetDate
	// Sender's or/and receiver's settlements or a part of them.
	Settlements OptSendingFilterGetSettlements
	// Sending length range or a part of range.
	Length OptSendingFilterGetLength
	// Sending width range or a part of range.
	Width OptSendingFilterGetWidth
	// Sending height range or a part of range.
	Height OptSendingFilterGetHeight
	// Sending weight range or a part of range.
	Weight OptSendingFilterGetWeight
	// Sending sort.
	Sort OptSendingSort
}

func unpackSendingFilterGetParams(packed map[string]any) (params SendingFilterGetParams) {
	params.Page = packed["page"].(Page)
	params.ElemsOnPage = packed["elems_on_page"].(ElementsOnPage)
	if v, ok := packed["order_id"]; ok {
		params.OrderID = v.(OptString)
	}
	if v, ok := packed["type"]; ok {
		params.Type = v.([]SendingType)
	}
	if v, ok := packed["status"]; ok {
		params.Status = v.([]SendingStatus)
	}
	if v, ok := packed["date"]; ok {
		params.Date = v.(OptSendingFilterGetDate)
	}
	if v, ok := packed["settlements"]; ok {
		params.Settlements = v.(OptSendingFilterGetSettlements)
	}
	if v, ok := packed["length"]; ok {
		params.Length = v.(OptSendingFilterGetLength)
	}
	if v, ok := packed["width"]; ok {
		params.Width = v.(OptSendingFilterGetWidth)
	}
	if v, ok := packed["height"]; ok {
		params.Height = v.(OptSendingFilterGetHeight)
	}
	if v, ok := packed["weight"]; ok {
		params.Weight = v.(OptSendingFilterGetWeight)
	}
	if v, ok := packed["sort"]; ok {
		params.Sort = v.(OptSendingSort)
	}
	return params
}

func decodeSendingFilterGetParams(args [0]string, r *http.Request) (params SendingFilterGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page = Page(paramsDotPageVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
			if err := func() error {
				if err := params.Page.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: invalid")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: elems_on_page.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "elems_on_page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotElemsOnPageVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					paramsDotElemsOnPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ElemsOnPage = ElementsOnPage(paramsDotElemsOnPageVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: elems_on_page: parse")
			}
			if err := func() error {
				if err := params.ElemsOnPage.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: elems_on_page: invalid")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: order_id.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderID.SetTo(paramsDotOrderIDVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: order_id: parse")
			}
			if err := func() error {
				if params.OrderID.Set {
					if err := func() error {
						if err := (validate.String{
							MinLength:    1,
							MinLengthSet: true,
							MaxLength:    36,
							MaxLengthSet: true,
							Email:        false,
							Hostname:     false,
							Regex:        nil,
						}).Validate(string(params.OrderID.Value)); err != nil {
							return errors.Wrap(err, "string")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: order_id: invalid")
			}
		}
	}
	// Decode query: type.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotTypeVal SendingType
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotTypeVal = SendingType(c)
						return nil
					}(); err != nil {
						return err
					}
					params.Type = append(params.Type, paramsDotTypeVal)
					return nil
				})
			}); err != nil {
				return params, errors.Wrap(err, "query: type: parse")
			}
			if err := func() error {
				if err := (validate.Array{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
				}).ValidateLength(len(params.Type)); err != nil {
					return errors.Wrap(err, "array")
				}
				var failures []validate.FieldError
				for i, elem := range params.Type {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: type: invalid")
			}
		}
	}
	// Decode query: status.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotStatusVal SendingStatus
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotStatusVal = SendingStatus(c)
						return nil
					}(); err != nil {
						return err
					}
					params.Status = append(params.Status, paramsDotStatusVal)
					return nil
				})
			}); err != nil {
				return params, errors.Wrap(err, "query: status: parse")
			}
			if err := func() error {
				if err := (validate.Array{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
				}).ValidateLength(len(params.Status)); err != nil {
					return errors.Wrap(err, "array")
				}
				var failures []validate.FieldError
				for i, elem := range params.Status {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: status: invalid")
			}
		}
	}
	// Decode query: date.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "date",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"date_start", false}, {"date_finish", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDateVal SendingFilterGetDate
				if err := func() error {
					return paramsDotDateVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Date.SetTo(paramsDotDateVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: date: parse")
			}
		}
	}
	// Decode query: settlements.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "settlements",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"sender_settlement", false}, {"receiver_settlement", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSettlementsVal SendingFilterGetSettlements
				if err := func() error {
					return paramsDotSettlementsVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Settlements.SetTo(paramsDotSettlementsVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: settlements: parse")
			}
			if err := func() error {
				if params.Settlements.Set {
					if err := func() error {
						if err := params.Settlements.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: settlements: invalid")
			}
		}
	}
	// Decode query: length.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "length",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"length_min", false}, {"length_max", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLengthVal SendingFilterGetLength
				if err := func() error {
					return paramsDotLengthVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Length.SetTo(paramsDotLengthVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: length: parse")
			}
			if err := func() error {
				if params.Length.Set {
					if err := func() error {
						if err := params.Length.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: length: invalid")
			}
		}
	}
	// Decode query: width.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "width",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"width_min", false}, {"width_max", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotWidthVal SendingFilterGetWidth
				if err := func() error {
					return paramsDotWidthVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Width.SetTo(paramsDotWidthVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: width: parse")
			}
			if err := func() error {
				if params.Width.Set {
					if err := func() error {
						if err := params.Width.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: width: invalid")
			}
		}
	}
	// Decode query: height.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "height",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"height_min", false}, {"height_max", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotHeightVal SendingFilterGetHeight
				if err := func() error {
					return paramsDotHeightVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Height.SetTo(paramsDotHeightVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: height: parse")
			}
			if err := func() error {
				if params.Height.Set {
					if err := func() error {
						if err := params.Height.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: height: invalid")
			}
		}
	}
	// Decode query: weight.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "weight",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"weight_min", false}, {"weight_max", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotWeightVal SendingFilterGetWeight
				if err := func() error {
					return paramsDotWeightVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Weight.SetTo(paramsDotWeightVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: weight: parse")
			}
			if err := func() error {
				if params.Weight.Set {
					if err := func() error {
						if err := params.Weight.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: weight: invalid")
			}
		}
	}
	// Decode query: sort.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"sort_type", false}, {"sort_field", false}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal SendingSort
				if err := func() error {
					return paramsDotSortVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: sort: parse")
			}
			if err := func() error {
				if params.Sort.Set {
					if err := func() error {
						if err := params.Sort.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: sort: invalid")
			}
		}
	}
	return params, nil
}

// SendingGetParams is parameters of GET /sending operation.
type SendingGetParams struct {
	// ID of sending.
	OrderID SendingOrderID
}

func unpackSendingGetParams(packed map[string]any) (params SendingGetParams) {
	params.OrderID = packed["order_id"].(SendingOrderID)
	return params
}

func decodeSendingGetParams(args [0]string, r *http.Request) (params SendingGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: order_id.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderIDVal uuid.UUID
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToUUID(val)
					if err != nil {
						return err
					}

					paramsDotOrderIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.OrderID = SendingOrderID(paramsDotOrderIDVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: order_id: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	return params, nil
}
