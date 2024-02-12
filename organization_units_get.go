package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
)

func (c *Client) NewOrganizationUnitsGet() OrganizationUnitsGet {
	r := OrganizationUnitsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrganizationUnitsGet struct {
	client      *Client
	queryParams *OrganizationUnitsGetQueryParams
	pathParams  *OrganizationUnitsGetPathParams
	method      string
	headers     http.Header
	requestBody OrganizationUnitsGetBody
}

func (r OrganizationUnitsGet) NewQueryParams() *OrganizationUnitsGetQueryParams {
	return &OrganizationUnitsGetQueryParams{}
}

type OrganizationUnitsGetQueryParams struct{}

func (p OrganizationUnitsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OrganizationUnitsGet) QueryParams() *OrganizationUnitsGetQueryParams {
	return r.queryParams
}

func (r OrganizationUnitsGet) NewPathParams() *OrganizationUnitsGetPathParams {
	return &OrganizationUnitsGetPathParams{}
}

type OrganizationUnitsGetPathParams struct {
	JobID string `schema:"jobId"`
}

func (p *OrganizationUnitsGetPathParams) Params() map[string]string {
	return map[string]string{
		"jobId": p.JobID,
	}
}

func (r *OrganizationUnitsGet) PathParams() *OrganizationUnitsGetPathParams {
	return r.pathParams
}

func (r *OrganizationUnitsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrganizationUnitsGet) SetMethod(method string) {
	r.method = method
}

func (r *OrganizationUnitsGet) Method() string {
	return r.method
}

func (r OrganizationUnitsGet) NewRequestBody() OrganizationUnitsGetBody {
	return OrganizationUnitsGetBody{}
}

type OrganizationUnitsGetBody struct {
}

func (r *OrganizationUnitsGet) RequestBody() *OrganizationUnitsGetBody {
	return nil
}

func (r *OrganizationUnitsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrganizationUnitsGet) SetRequestBody(body OrganizationUnitsGetBody) {
	r.requestBody = body
}

func (r *OrganizationUnitsGet) NewResponseBody() *OrganizationUnitsGetResponseBody {
	return &OrganizationUnitsGetResponseBody{}
}

type OrganizationUnitsGetResponseBody struct {
	OrganizationUnitsFileUris []string `json:"organizationUnitsFileUris"`
	Status              string   `json:"status"`
}

func (r *OrganizationUnitsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/query/nl/hrm/metadata/organization-units/{{.jobId}}", r.PathParams())
	return &u
}

func (r *OrganizationUnitsGet) Do() (OrganizationUnitsGetResponseBody, error) {
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

