package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewCostCentersPost() CostCentersPost {
	r := CostCentersPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CostCentersPost struct {
	client      *Client
	queryParams *CostCentersPostQueryParams
	pathParams  *CostCentersPostPathParams
	method      string
	headers     http.Header
	requestBody CostCentersPostBody
}

func (r CostCentersPost) NewQueryParams() *CostCentersPostQueryParams {
	return &CostCentersPostQueryParams{}
}

type CostCentersPostQueryParams struct {
}

func (p CostCentersPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CostCentersPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r CostCentersPost) NewPathParams() *CostCentersPostPathParams {
	return &CostCentersPostPathParams{}
}

type CostCentersPostPathParams struct {
}

func (p *CostCentersPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostCentersPost) PathParams() *CostCentersPostPathParams {
	return r.pathParams
}

func (r *CostCentersPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CostCentersPost) SetMethod(method string) {
	r.method = method
}

func (r *CostCentersPost) Method() string {
	return r.method
}

func (r CostCentersPost) NewRequestBody() CostCentersPostBody {
	return CostCentersPostBody{}
}

type CostCentersPostBody struct {
}

func (r *CostCentersPost) RequestBody() *CostCentersPostBody {
	return nil
}

func (r *CostCentersPost) RequestBodyInterface() interface{} {
	return nil
}

func (r *CostCentersPost) SetRequestBody(body CostCentersPostBody) {
	r.requestBody = body
}

func (r *CostCentersPost) NewResponseBody() *CostCentersPostResponseBody {
	return &CostCentersPostResponseBody{}
}

type CostCentersPostResponseBody struct {
	JobID string `json:"jobId"`
}

func (r *CostCentersPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/command/nl/hrm/metadata/cost-centers", r.PathParams())
	return &u
}

func (r *CostCentersPost) Do() (CostCentersPostResponseBody, error) {
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
