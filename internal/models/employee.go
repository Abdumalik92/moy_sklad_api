package models

type Auth struct {
	AccessToken string `json:"access_token"`
}

type Request struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Response struct {
	Context struct {
		Employee struct {
			Meta Meta `json:"meta"`
		} `json:"employee"`
	} `json:"context"`
	Meta Meta  `json:"meta"`
	Rows []Row `json:"rows"`
}

type CreateEmployee struct {
	ID          string `json:"-"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	Inn         string `json:"inn"`
	Position    string `json:"position"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
type Meta struct {
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref"`
	Type         string `json:"type"`
	MediaType    string `json:"mediaType"`
	UUIDHref     string `json:"uuidHref"`
	DownloadHref string `json:"downloadHref,omitempty"`
	Size         int    `json:"size"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}
type Row struct {
	Meta      Meta   `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     struct {
		Meta Meta `json:"meta"`
	} `json:"owner"`
	Shared bool `json:"shared"`
	Group  struct {
		Meta Meta `json:"meta"`
	} `json:"group"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	Created      string `json:"created"`
	UID          string `json:"uid"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	LastName     string `json:"lastName"`
	FullName     string `json:"fullName"`
	ShortFio     string `json:"shortFio"`
	Inn          string `json:"inn"`
	Position     string `json:"position"`
	Cashiers     []struct {
		Meta Meta `json:"meta"`
	} `json:"cashiers"`
}
