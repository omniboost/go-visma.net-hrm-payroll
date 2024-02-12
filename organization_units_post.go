package vismanet

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/omniboost/go-visma.net-hrm-payroll/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewOrganizationUnitsPost() OrganizationUnitsPost {
	r := OrganizationUnitsPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrganizationUnitsPost struct {
	client      *Client
	queryParams *OrganizationUnitsPostQueryParams
	pathParams  *OrganizationUnitsPostPathParams
	method      string
	headers     http.Header
	requestBody OrganizationUnitsPostBody
}

func (r OrganizationUnitsPost) NewQueryParams() *OrganizationUnitsPostQueryParams {
	return &OrganizationUnitsPostQueryParams{}
}

type OrganizationUnitsPostQueryParams struct {
}

func (p OrganizationUnitsPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OrganizationUnitsPost) QueryParams() QueryParams {
	return r.queryParams
}

func (r OrganizationUnitsPost) NewPathParams() *OrganizationUnitsPostPathParams {
	return &OrganizationUnitsPostPathParams{}
}

type OrganizationUnitsPostPathParams struct {
}

func (p *OrganizationUnitsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrganizationUnitsPost) PathParams() *OrganizationUnitsPostPathParams {
	return r.pathParams
}

func (r *OrganizationUnitsPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrganizationUnitsPost) SetMethod(method string) {
	r.method = method
}

func (r *OrganizationUnitsPost) Method() string {
	return r.method
}

func (r OrganizationUnitsPost) NewRequestBody() OrganizationUnitsPostBody {
	return OrganizationUnitsPostBody{}
}

type OrganizationUnitsPostBody struct {
}

func (r *OrganizationUnitsPost) RequestBody() *OrganizationUnitsPostBody {
	return nil
}

func (r *OrganizationUnitsPost) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrganizationUnitsPost) SetRequestBody(body OrganizationUnitsPostBody) {
	r.requestBody = body
}

func (r *OrganizationUnitsPost) NewResponseBody() *OrganizationUnitsPostResponseBody {
	return &OrganizationUnitsPostResponseBody{}
}

type OrganizationUnitsPostResponseBody struct {
	JobID string `json:"jobId"`
}

func (r *OrganizationUnitsPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/command/nl/hrm/metadata/organization-units", r.PathParams())
	return &u
}

func (r *OrganizationUnitsPost) Do() (OrganizationUnitsPostResponseBody, error) {
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

func (r *OrganizationUnitsPost) DoCSV() (OrganizationUnits, error) {
	units := OrganizationUnits{}
	postresp, err := r.Do()
	if err != nil {
		return units, errors.WithStack(err)
	}

	waiting := true
	gresp := OrganizationUnitsGetResponseBody{}
	i := 0
	for waiting == true {
		i = i + 1
		greq := r.client.NewOrganizationUnitsGet()
		greq.PathParams().JobID = postresp.JobID
		gresp, err = greq.Do()
		if err != nil {
			return units, errors.WithStack(err)
		}

		waiting = gresp.Status == "InProgress"
		time.Sleep(time.Duration(i) * time.Second)
	}

	for _, u := range gresp.OrganizationUnitsFileUris {
		r, err := http.Get(u)
		if err != nil {
			return units, errors.WithStack(err)
		}

		ee := OrganizationUnits{}
		err = gocsv.Unmarshal(r.Body, &ee)
		if err != nil {
			return units, errors.WithStack(err)
		}
		units = append(units, ee...)
	}

	return units, nil
}

