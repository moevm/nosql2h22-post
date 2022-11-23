// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// EncodeURI encodes SendingFilter as URI form.
func (s SendingFilter) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("order_id", func(e uri.Encoder) error {
		if val, ok := s.OrderID.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"order_id\"")
	}
	if err := e.EncodeField("type", func(e uri.Encoder) error {
		if val, ok := s.Type.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"type\"")
	}
	if err := e.EncodeField("status", func(e uri.Encoder) error {
		if val, ok := s.Status.Get(); ok {
			return e.EncodeValue(conv.StringToString(string(val)))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"status\"")
	}
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
	if err := e.EncodeField("length", func(e uri.Encoder) error {
		if val, ok := s.Length.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"length\"")
	}
	if err := e.EncodeField("width", func(e uri.Encoder) error {
		if val, ok := s.Width.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"width\"")
	}
	if err := e.EncodeField("height", func(e uri.Encoder) error {
		if val, ok := s.Height.Get(); ok {
			return e.EncodeValue(conv.Int64ToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"height\"")
	}
	if err := e.EncodeField("weight", func(e uri.Encoder) error {
		if val, ok := s.Weight.Get(); ok {
			if unwrapped := int64(val); true {
				return e.EncodeValue(conv.Int64ToString(unwrapped))
			}
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"weight\"")
	}
	return nil
}

var uriFieldsNameOfSendingFilter = [11]string{
	0:  "order_id",
	1:  "type",
	2:  "status",
	3:  "date_start",
	4:  "date_finish",
	5:  "sender_settlement",
	6:  "receiver_settlement",
	7:  "length",
	8:  "width",
	9:  "height",
	10: "weight",
}

// DecodeURI decodes SendingFilter from URI form.
func (s *SendingFilter) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendingFilter to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "order_id":
			if err := func() error {
				var sDotOrderIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotOrderIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.OrderID.SetTo(sDotOrderIDVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order_id\"")
			}
		case "type":
			if err := func() error {
				var sDotTypeVal SendingType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotTypeVal = SendingType(c)
					return nil
				}(); err != nil {
					return err
				}
				s.Type.SetTo(sDotTypeVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "status":
			if err := func() error {
				var sDotStatusVal SendingStatus
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotStatusVal = SendingStatus(c)
					return nil
				}(); err != nil {
					return err
				}
				s.Status.SetTo(sDotStatusVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
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
		case "length":
			if err := func() error {
				var sDotLengthVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotLengthVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Length.SetTo(sDotLengthVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"length\"")
			}
		case "width":
			if err := func() error {
				var sDotWidthVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotWidthVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Width.SetTo(sDotWidthVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"width\"")
			}
		case "height":
			if err := func() error {
				var sDotHeightVal int64
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt64(val)
					if err != nil {
						return err
					}

					sDotHeightVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Height.SetTo(sDotHeightVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"height\"")
			}
		case "weight":
			if err := func() error {
				var sDotWeightVal SendingWeight
				if err := func() error {
					var sDotWeightValVal int64
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt64(val)
						if err != nil {
							return err
						}

						sDotWeightValVal = c
						return nil
					}(); err != nil {
						return err
					}
					sDotWeightVal = SendingWeight(sDotWeightValVal)
					return nil
				}(); err != nil {
					return err
				}
				s.Weight.SetTo(sDotWeightVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"weight\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendingFilter")
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
