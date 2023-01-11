package vismanet

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewEmployeesGet() EmployeesGet {
	r := EmployeesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesGet struct {
	client      *Client
	queryParams *EmployeesGetQueryParams
	pathParams  *EmployeesGetPathParams
	method      string
	headers     http.Header
	requestBody EmployeesGetBody
}

func (r EmployeesGet) NewQueryParams() *EmployeesGetQueryParams {
	return &EmployeesGetQueryParams{}
}

type EmployeesGetQueryParams struct{}

func (p EmployeesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesGet) QueryParams() *EmployeesGetQueryParams {
	return r.queryParams
}

func (r EmployeesGet) NewPathParams() *EmployeesGetPathParams {
	return &EmployeesGetPathParams{}
}

type EmployeesGetPathParams struct {
	JobID string `schema:"jobId"`
}

func (p *EmployeesGetPathParams) Params() map[string]string {
	return map[string]string{
		"jobId": p.JobID,
	}
}

func (r *EmployeesGet) PathParams() *EmployeesGetPathParams {
	return r.pathParams
}

func (r *EmployeesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesGet) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesGet) Method() string {
	return r.method
}

func (r EmployeesGet) NewRequestBody() EmployeesGetBody {
	return EmployeesGetBody{}
}

type EmployeesGetBody struct {
}

func (r *EmployeesGet) RequestBody() *EmployeesGetBody {
	return nil
}

func (r *EmployeesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeesGet) SetRequestBody(body EmployeesGetBody) {
	r.requestBody = body
}

func (r *EmployeesGet) NewResponseBody() *EmployeesGetResponseBody {
	return &EmployeesGetResponseBody{}
}

type EmployeesGetResponseBody struct {
	ChangeTimestampBefore time.Time `json:"changeTimestampBefore"`
	EmployeeFileUris      []string  `json:"employeeFileUris"`
	Status                string    `json:"status"`
}

func (r *EmployeesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/query/nl/hrm/employees/{{.jobId}}", r.PathParams())
	return &u
}

func (r *EmployeesGet) Do() (EmployeesGetResponseBody, error) {
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
