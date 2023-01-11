package vismanet

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewEmployeesPost() EmployeesPost {
	r := EmployeesPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesPost struct {
	client      *Client
	queryParams *EmployeesPostQueryParams
	pathParams  *EmployeesPostPathParams
	method      string
	headers     http.Header
	requestBody EmployeesPostBody
}

func (r EmployeesPost) NewQueryParams() *EmployeesPostQueryParams {
	return &EmployeesPostQueryParams{}
}

type EmployeesPostQueryParams struct {
}

func (p EmployeesPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r EmployeesPost) NewPathParams() *EmployeesPostPathParams {
	return &EmployeesPostPathParams{}
}

type EmployeesPostPathParams struct {
}

func (p *EmployeesPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeesPost) PathParams() *EmployeesPostPathParams {
	return r.pathParams
}

func (r *EmployeesPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesPost) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesPost) Method() string {
	return r.method
}

func (r EmployeesPost) NewRequestBody() EmployeesPostBody {
	return EmployeesPostBody{}
}

type EmployeesPostBody struct {
}

func (r *EmployeesPost) RequestBody() *EmployeesPostBody {
	return nil
}

func (r *EmployeesPost) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeesPost) SetRequestBody(body EmployeesPostBody) {
	r.requestBody = body
}

func (r *EmployeesPost) NewResponseBody() *EmployeesPostResponseBody {
	return &EmployeesPostResponseBody{}
}

type EmployeesPostResponseBody struct {
	JobID string `json:"jobId"`
}

func (r *EmployeesPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/command/nl/hrm/employees", r.PathParams())
	return &u
}

func (r *EmployeesPost) Do() (EmployeesPostResponseBody, error) {
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

func (r *EmployeesPost) DoCSV() (Employees, error) {
	employees := Employees{}
	postresp, err := r.Do()
	if err != nil {
		return employees, errors.WithStack(err)
	}

	waiting := true
	gresp := EmployeesGetResponseBody{}
	i := 0
	for waiting == true {
		i = i + 1
		greq := r.client.NewEmployeesGet()
		greq.PathParams().JobID = postresp.JobID
		gresp, err = greq.Do()
		if err != nil {
			return employees, errors.WithStack(err)
		}

		waiting = gresp.Status == "InProgress"
		time.Sleep(time.Duration(i) * time.Second)
	}

	for _, u := range gresp.EmployeeFileUris {
		r, err := http.Get(u)
		if err != nil {
			return employees, errors.WithStack(err)
		}

		ee := Employees{}
		err = gocsv.Unmarshal(r.Body, &ee)
		if err != nil {
			return employees, errors.WithStack(err)
		}
		employees = append(employees, ee...)
	}

	return employees, nil
}
