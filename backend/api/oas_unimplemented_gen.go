// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DataExportSendingGet implements GET /data_export_sending operation.
//
// Get data from collection `sendings` from database. Return json array.
//
// GET /data_export_sending
func (UnimplementedHandler) DataExportSendingGet(ctx context.Context) (r []Sending, _ error) {
	return r, ht.ErrNotImplemented
}

// DataImportSendingPost implements POST /data_import_sending operation.
//
// Import data into collection `sendings` in database. Require array of json.
//
// POST /data_import_sending
func (UnimplementedHandler) DataImportSendingPost(ctx context.Context, req []Sending) (r DataImportSendingPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EmployeeFilterGet implements GET /employee_filter operation.
//
// Get employees that fit the filter. Require `page` and `elems_on_page`. Return amount of employees
// that fit the filter and employees on the selected page.
//
// GET /employee_filter
func (UnimplementedHandler) EmployeeFilterGet(ctx context.Context, params EmployeeFilterGetParams) (r EmployeeFilterGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostcodesBySettlementGet implements GET /postcodes_by_settlement operation.
//
// Get information about postcodes in cities. Return map with `settlement` key and `postcode` array
// value.
//
// GET /postcodes_by_settlement
func (UnimplementedHandler) PostcodesBySettlementGet(ctx context.Context, params PostcodesBySettlementGetParams) (r PostcodesBySettlementGetResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SendingFilterGet implements GET /sending_filter operation.
//
// Get sendings that fit the filter. Require `page` and `elems_on_page`. Return amount of sendings
// that fit the filter and sendings on the selected page.
//
// GET /sending_filter
func (UnimplementedHandler) SendingFilterGet(ctx context.Context, params SendingFilterGetParams) (r SendingFilterGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SendingGet implements GET /sending operation.
//
// Get information about a sending by `order_id`. Require a complete match of `order_id`. Return
// `type`, `status` and `stages`.
//
// GET /sending
func (UnimplementedHandler) SendingGet(ctx context.Context, params SendingGetParams) (r SendingGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SendingPost implements POST /sending operation.
//
// Registration of a new sending. Require `type`, `sender`, `receiver`, `size`, `weight`. Return
// `order_id` of new sending.
//
// POST /sending
func (UnimplementedHandler) SendingPost(ctx context.Context, req SendingPostReq) (r SendingPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SendingStatisticsGet implements GET /sending_statistics operation.
//
// Get statistics of sendings. Return array of keys and statistic value.
//
// GET /sending_statistics
func (UnimplementedHandler) SendingStatisticsGet(ctx context.Context, params SendingStatisticsGetParams) (r SendingStatisticsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}
