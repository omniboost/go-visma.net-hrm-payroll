package vismanet

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewContractsGet() ContractsGet {
	r := ContractsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContractsGet struct {
	client      *Client
	queryParams *ContractsGetQueryParams
	pathParams  *ContractsGetPathParams
	method      string
	headers     http.Header
	requestBody ContractsGetBody
}

func (r ContractsGet) NewQueryParams() *ContractsGetQueryParams {
	return &ContractsGetQueryParams{}
}

type ContractsGetQueryParams struct{}

func (p ContractsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ContractsGet) QueryParams() *ContractsGetQueryParams {
	return r.queryParams
}

func (r ContractsGet) NewPathParams() *ContractsGetPathParams {
	return &ContractsGetPathParams{}
}

type ContractsGetPathParams struct {
	JobID string `schema:"jobId"`
}

func (p *ContractsGetPathParams) Params() map[string]string {
	return map[string]string{
		"jobId": p.JobID,
	}
}

func (r *ContractsGet) PathParams() *ContractsGetPathParams {
	return r.pathParams
}

func (r *ContractsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ContractsGet) SetMethod(method string) {
	r.method = method
}

func (r *ContractsGet) Method() string {
	return r.method
}

func (r ContractsGet) NewRequestBody() ContractsGetBody {
	return ContractsGetBody{}
}

type ContractsGetBody struct {
}

func (r *ContractsGet) RequestBody() *ContractsGetBody {
	return nil
}

func (r *ContractsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *ContractsGet) SetRequestBody(body ContractsGetBody) {
	r.requestBody = body
}

func (r *ContractsGet) NewResponseBody() *ContractsGetResponseBody {
	return &ContractsGetResponseBody{}
}

type ContractsGetResponseBody struct {
	ChangeTimestampBefore time.Time `json:"changeTimestampBefore"`
	ContractFileUris      []string  `json:"contractFileUris"`
	Status                string    `json:"status"`
}

func (r *ContractsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/query/nl/hrm/contracts/{{.jobId}}", r.PathParams())
	return &u
}

func (r *ContractsGet) Do() (ContractsGetResponseBody, error) {
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
