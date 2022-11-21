// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

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
