package kibana

const savedObjectsPath = "/api/saved_objects/"

type SavedObjectRequest struct {
	Type    string   `json:"type" url:"type"`
	Fields  []string `json:"fields" url:"fields"`
	Page    int      `json:"page" url:"page"`
	PerPage int      `json:"per_page" url:"per_page"`
}

type SavedObjectRequestBuilder struct {
	objectType string
	fields     []string
	page       int
	perPage    int
}

type SavedObjectsClient interface {
	GetByType(request *SavedObjectRequest) (*SavedObjectResponse, error)
}

type SavedObjectResponse struct {
	Page         int            `json:"page"`
	PerPage      int            `json:"per_page"`
	Total        int            `json:"total"`
	SavedObjects []*SavedObject `json:"saved_objects"`
}

type SavedObject struct {
	Id         string                 `json:"id"`
	Type       string                 `json:"type"`
	Version    version                `json:"version"`
	Attributes map[string]interface{} `json:"attributes"`
}

func NewSavedObjectRequestBuilder() *SavedObjectRequestBuilder {
	return &SavedObjectRequestBuilder{perPage: 20, page: 1}
}

func (builder *SavedObjectRequestBuilder) WithType(objectType string) *SavedObjectRequestBuilder {
	builder.objectType = objectType
	return builder
}

func (builder *SavedObjectRequestBuilder) WithFields(fields []string) *SavedObjectRequestBuilder {
	builder.fields = fields
	return builder
}

func (builder *SavedObjectRequestBuilder) WithPage(page int) *SavedObjectRequestBuilder {
	builder.page = page
	return builder
}

func (builder *SavedObjectRequestBuilder) WithPerPage(perPage int) *SavedObjectRequestBuilder {
	builder.perPage = perPage
	return builder
}

func (builder *SavedObjectRequestBuilder) Build() *SavedObjectRequest {
	return &SavedObjectRequest{
		Fields:  builder.fields,
		Type:    builder.objectType,
		PerPage: builder.perPage,
	}
}
