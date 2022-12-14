// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// EncodeURI encodes EmployeeFilterGetBirthDate as URI form.
func (s EmployeeFilterGetBirthDate) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("birth_date_start", func(e uri.Encoder) error {
		if val, ok := s.BirthDateStart.Get(); ok {
			return e.EncodeValue(conv.DateToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"birth_date_start\"")
	}
	if err := e.EncodeField("birth_date_finish", func(e uri.Encoder) error {
		if val, ok := s.BirthDateFinish.Get(); ok {
			return e.EncodeValue(conv.DateToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"birth_date_finish\"")
	}
	return nil
}

var uriFieldsNameOfEmployeeFilterGetBirthDate = [2]string{
	0: "birth_date_start",
	1: "birth_date_finish",
}

// DecodeURI decodes EmployeeFilterGetBirthDate from URI form.
func (s *EmployeeFilterGetBirthDate) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode EmployeeFilterGetBirthDate to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "birth_date_start":
			if err := func() error {
				var sDotBirthDateStartVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDate(val)
					if err != nil {
						return err
					}

					sDotBirthDateStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.BirthDateStart.SetTo(sDotBirthDateStartVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birth_date_start\"")
			}
		case "birth_date_finish":
			if err := func() error {
				var sDotBirthDateFinishVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDate(val)
					if err != nil {
						return err
					}

					sDotBirthDateFinishVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.BirthDateFinish.SetTo(sDotBirthDateFinishVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birth_date_finish\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode EmployeeFilterGetBirthDate")
	}

	return nil
}

// EncodeURI encodes EmployeeFilterGetFullName as URI form.
func (s EmployeeFilterGetFullName) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("name", func(e uri.Encoder) error {
		if val, ok := s.Name.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"name\"")
	}
	if err := e.EncodeField("surname", func(e uri.Encoder) error {
		if val, ok := s.Surname.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"surname\"")
	}
	if err := e.EncodeField("middle_name", func(e uri.Encoder) error {
		if val, ok := s.MiddleName.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"middle_name\"")
	}
	return nil
}

var uriFieldsNameOfEmployeeFilterGetFullName = [3]string{
	0: "name",
	1: "surname",
	2: "middle_name",
}

// DecodeURI decodes EmployeeFilterGetFullName from URI form.
func (s *EmployeeFilterGetFullName) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode EmployeeFilterGetFullName to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "name":
			if err := func() error {
				var sDotNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Name.SetTo(sDotNameVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "surname":
			if err := func() error {
				var sDotSurnameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSurnameVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Surname.SetTo(sDotSurnameVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"surname\"")
			}
		case "middle_name":
			if err := func() error {
				var sDotMiddleNameVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotMiddleNameVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.MiddleName.SetTo(sDotMiddleNameVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"middle_name\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode EmployeeFilterGetFullName")
	}

	return nil
}

// EncodeURI encodes EmployeeSort as URI form.
func (s EmployeeSort) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("sort_type", func(e uri.Encoder) error {
		if val, ok := s.SortType.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"sort_type\"")
	}
	if err := e.EncodeField("sort_field", func(e uri.Encoder) error {
		if val, ok := s.SortField.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"sort_field\"")
	}
	return nil
}

var uriFieldsNameOfEmployeeSort = [2]string{
	0: "sort_type",
	1: "sort_field",
}

// DecodeURI decodes EmployeeSort from URI form.
func (s *EmployeeSort) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode EmployeeSort to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "sort_type":
			if err := func() error {
				var sDotSortTypeVal EmployeeSortSortType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSortTypeVal = EmployeeSortSortType(c)
					return nil
				}(); err != nil {
					return err
				}
				s.SortType.SetTo(sDotSortTypeVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sort_type\"")
			}
		case "sort_field":
			if err := func() error {
				var sDotSortFieldVal EmployeeSortSortField
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSortFieldVal = EmployeeSortSortField(c)
					return nil
				}(); err != nil {
					return err
				}
				s.SortField.SetTo(sDotSortFieldVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sort_field\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode EmployeeSort")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetDate as URI form.
func (s SendingFilterGetDate) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("date_start", func(e uri.Encoder) error {
		if val, ok := s.DateStart.Get(); ok {
			return e.EncodeValue(conv.DateToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"date_start\"")
	}
	if err := e.EncodeField("date_finish", func(e uri.Encoder) error {
		if val, ok := s.DateFinish.Get(); ok {
			return e.EncodeValue(conv.DateToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"date_finish\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetDate = [2]string{
	0: "date_start",
	1: "date_finish",
}

// DecodeURI decodes SendingFilterGetDate from URI form.
func (s *SendingFilterGetDate) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetDate to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "date_start":
			if err := func() error {
				var sDotDateStartVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDate(val)
					if err != nil {
						return err
					}

					sDotDateStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.DateStart.SetTo(sDotDateStartVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date_start\"")
			}
		case "date_finish":
			if err := func() error {
				var sDotDateFinishVal time.Time
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToDate(val)
					if err != nil {
						return err
					}

					sDotDateFinishVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.DateFinish.SetTo(sDotDateFinishVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date_finish\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetDate")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetHeight as URI form.
func (s SendingFilterGetHeight) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("height_min", func(e uri.Encoder) error {
		if val, ok := s.HeightMin.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"height_min\"")
	}
	if err := e.EncodeField("height_max", func(e uri.Encoder) error {
		if val, ok := s.HeightMax.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"height_max\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetHeight = [2]string{
	0: "height_min",
	1: "height_max",
}

// DecodeURI decodes SendingFilterGetHeight from URI form.
func (s *SendingFilterGetHeight) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetHeight to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "height_min":
			if err := func() error {
				var sDotHeightMinVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotHeightMinVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.HeightMin.SetTo(sDotHeightMinVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height_min\"")
			}
		case "height_max":
			if err := func() error {
				var sDotHeightMaxVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotHeightMaxVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.HeightMax.SetTo(sDotHeightMaxVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height_max\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetHeight")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetLength as URI form.
func (s SendingFilterGetLength) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("length_min", func(e uri.Encoder) error {
		if val, ok := s.LengthMin.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"length_min\"")
	}
	if err := e.EncodeField("length_max", func(e uri.Encoder) error {
		if val, ok := s.LengthMax.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"length_max\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetLength = [2]string{
	0: "length_min",
	1: "length_max",
}

// DecodeURI decodes SendingFilterGetLength from URI form.
func (s *SendingFilterGetLength) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetLength to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "length_min":
			if err := func() error {
				var sDotLengthMinVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotLengthMinVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.LengthMin.SetTo(sDotLengthMinVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"length_min\"")
			}
		case "length_max":
			if err := func() error {
				var sDotLengthMaxVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotLengthMaxVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.LengthMax.SetTo(sDotLengthMaxVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"length_max\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetLength")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetSettlements as URI form.
func (s SendingFilterGetSettlements) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("sender_settlement", func(e uri.Encoder) error {
		if val, ok := s.SenderSettlement.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"sender_settlement\"")
	}
	if err := e.EncodeField("receiver_settlement", func(e uri.Encoder) error {
		if val, ok := s.ReceiverSettlement.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"receiver_settlement\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetSettlements = [2]string{
	0: "sender_settlement",
	1: "receiver_settlement",
}

// DecodeURI decodes SendingFilterGetSettlements from URI form.
func (s *SendingFilterGetSettlements) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetSettlements to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "sender_settlement":
			if err := func() error {
				var sDotSenderSettlementVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSenderSettlementVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.SenderSettlement.SetTo(sDotSenderSettlementVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sender_settlement\"")
			}
		case "receiver_settlement":
			if err := func() error {
				var sDotReceiverSettlementVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotReceiverSettlementVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.ReceiverSettlement.SetTo(sDotReceiverSettlementVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"receiver_settlement\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetSettlements")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetWeight as URI form.
func (s SendingFilterGetWeight) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("weight_min", func(e uri.Encoder) error {
		if val, ok := s.WeightMin.Get(); ok {
			if unwrapped := int64(val); true {
				return e.EncodeValue(conv.Int64ToString(unwrapped))
			}
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"weight_min\"")
	}
	if err := e.EncodeField("weight_max", func(e uri.Encoder) error {
		if val, ok := s.WeightMax.Get(); ok {
			if unwrapped := int64(val); true {
				return e.EncodeValue(conv.Int64ToString(unwrapped))
			}
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"weight_max\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetWeight = [2]string{
	0: "weight_min",
	1: "weight_max",
}

// DecodeURI decodes SendingFilterGetWeight from URI form.
func (s *SendingFilterGetWeight) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetWeight to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "weight_min":
			if err := func() error {
				var sDotWeightMinVal SendingWeight
				if err := func() error {
					var sDotWeightMinValVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						sDotWeightMinValVal = c
						return nil
					}(); err != nil {
						return err
					}
					sDotWeightMinVal = SendingWeight(sDotWeightMinValVal)
					return nil
				}(); err != nil {
					return err
				}
				s.WeightMin.SetTo(sDotWeightMinVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"weight_min\"")
			}
		case "weight_max":
			if err := func() error {
				var sDotWeightMaxVal SendingWeight
				if err := func() error {
					var sDotWeightMaxValVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						sDotWeightMaxValVal = c
						return nil
					}(); err != nil {
						return err
					}
					sDotWeightMaxVal = SendingWeight(sDotWeightMaxValVal)
					return nil
				}(); err != nil {
					return err
				}
				s.WeightMax.SetTo(sDotWeightMaxVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"weight_max\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetWeight")
	}

	return nil
}

// EncodeURI encodes SendingFilterGetWidth as URI form.
func (s SendingFilterGetWidth) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("width_min", func(e uri.Encoder) error {
		if val, ok := s.WidthMin.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"width_min\"")
	}
	if err := e.EncodeField("width_max", func(e uri.Encoder) error {
		if val, ok := s.WidthMax.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"width_max\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilterGetWidth = [2]string{
	0: "width_min",
	1: "width_max",
}

// DecodeURI decodes SendingFilterGetWidth from URI form.
func (s *SendingFilterGetWidth) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilterGetWidth to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "width_min":
			if err := func() error {
				var sDotWidthMinVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotWidthMinVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.WidthMin.SetTo(sDotWidthMinVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width_min\"")
			}
		case "width_max":
			if err := func() error {
				var sDotWidthMaxVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotWidthMaxVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.WidthMax.SetTo(sDotWidthMaxVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width_max\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilterGetWidth")
	}

	return nil
}

// EncodeURI encodes SendingSort as URI form.
func (s SendingSort) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("sort_type", func(e uri.Encoder) error {
		if val, ok := s.SortType.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"sort_type\"")
	}
	if err := e.EncodeField("sort_field", func(e uri.Encoder) error {
		if val, ok := s.SortField.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"sort_field\"")
	}
	return nil
}

var uriFieldsNameOfSendingSort = [2]string{
	0: "sort_type",
	1: "sort_field",
}

// DecodeURI decodes SendingSort from URI form.
func (s *SendingSort) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingSort to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "sort_type":
			if err := func() error {
				var sDotSortTypeVal SendingSortSortType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSortTypeVal = SendingSortSortType(c)
					return nil
				}(); err != nil {
					return err
				}
				s.SortType.SetTo(sDotSortTypeVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sort_type\"")
			}
		case "sort_field":
			if err := func() error {
				var sDotSortFieldVal SendingSortSortField
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotSortFieldVal = SendingSortSortField(c)
					return nil
				}(); err != nil {
					return err
				}
				s.SortField.SetTo(sDotSortFieldVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"sort_field\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingSort")
	}

	return nil
}
