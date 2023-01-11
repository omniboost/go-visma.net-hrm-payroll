package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewContractsPost() ContractsPost {
	r := ContractsPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContractsPost struct {
	client      *Client
	queryParams *ContractsPostQueryParams
	pathParams  *ContractsPostPathParams
	method      string
	headers     http.Header
	requestBody ContractsPostBody
}

func (r ContractsPost) NewQueryParams() *ContractsPostQueryParams {
	return &ContractsPostQueryParams{}
}

type ContractsPostQueryParams struct {
}

func (p ContractsPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ContractsPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r ContractsPost) NewPathParams() *ContractsPostPathParams {
	return &ContractsPostPathParams{}
}

type ContractsPostPathParams struct {
}

func (p *ContractsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContractsPost) PathParams() *ContractsPostPathParams {
	return r.pathParams
}

func (r *ContractsPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ContractsPost) SetMethod(method string) {
	r.method = method
}

func (r *ContractsPost) Method() string {
	return r.method
}

func (r ContractsPost) NewRequestBody() ContractsPostBody {
	return ContractsPostBody{}
}

type ContractsPostBody struct {
}

func (r *ContractsPost) RequestBody() *ContractsPostBody {
	return nil
}

func (r *ContractsPost) RequestBodyInterface() interface{} {
	return nil
}

func (r *ContractsPost) SetRequestBody(body ContractsPostBody) {
	r.requestBody = body
}

func (r *ContractsPost) NewResponseBody() *ContractsPostResponseBody {
	return &ContractsPostResponseBody{}
}

type ContractsPostResponseBody struct {
	JobID string `json:"jobId"`
}

func (r *ContractsPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/command/nl/hrm/contracts", r.PathParams())
	return &u
}

func (r *ContractsPost) Do() (ContractsPostResponseBody, error) {
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
