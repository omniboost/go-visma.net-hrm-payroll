package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewCostCentersGet() CostCentersGet {
	r := CostCentersGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CostCentersGet struct {
	client      *Client
	queryParams *CostCentersGetQueryParams
	pathParams  *CostCentersGetPathParams
	method      string
	headers     http.Header
	requestBody CostCentersGetBody
}

func (r CostCentersGet) NewQueryParams() *CostCentersGetQueryParams {
	return &CostCentersGetQueryParams{}
}

type CostCentersGetQueryParams struct{}

func (p CostCentersGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CostCentersGet) QueryParams() *CostCentersGetQueryParams {
	return r.queryParams
}

func (r CostCentersGet) NewPathParams() *CostCentersGetPathParams {
	return &CostCentersGetPathParams{}
}

type CostCentersGetPathParams struct {
	JobID string `schema:"jobId"`
}

func (p *CostCentersGetPathParams) Params() map[string]string {
	return map[string]string{
		"jobId": p.JobID,
	}
}

func (r *CostCentersGet) PathParams() *CostCentersGetPathParams {
	return r.pathParams
}

func (r *CostCentersGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CostCentersGet) SetMethod(method string) {
	r.method = method
}

func (r *CostCentersGet) Method() string {
	return r.method
}

func (r CostCentersGet) NewRequestBody() CostCentersGetBody {
	return CostCentersGetBody{}
}

type CostCentersGetBody struct {
}

func (r *CostCentersGet) RequestBody() *CostCentersGetBody {
	return nil
}

func (r *CostCentersGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *CostCentersGet) SetRequestBody(body CostCentersGetBody) {
	r.requestBody = body
}

func (r *CostCentersGet) NewResponseBody() *CostCentersGetResponseBody {
	return &CostCentersGetResponseBody{}
}

type CostCentersGetResponseBody struct {
	CostCentersFileUris []string `json:"costCentersFileUris"`
	Status              string   `json:"status"`
}

func (r *CostCentersGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/query/nl/hrm/metadata/cost-centers/{{.jobId}}", r.PathParams())
	return &u
}

func (r *CostCentersGet) Do() (CostCentersGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
