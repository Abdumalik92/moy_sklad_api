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
			Meta Meta `json:"meta,omitempty"`
		} `json:"employee,omitempty"`
	} `json:"context,omitempty"`
	Meta Meta  `json:"meta,omitempty"`
	Rows []Row `json:"rows,omitempty"`
}

type CreateEmployee struct {
	ID          string `json:"-"`
	FirstName   string `json:"firstName,omitempty"`
	MiddleName  string `json:"middleName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Inn         string `json:"inn,omitempty"`
	Position    string `json:"position,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Description string `json:"description,omitempty"`
}
type Meta struct {
	Href         string `json:"href,omitempty"`
	MetadataHref string `json:"metadataHref,omitempty"`
	Type         string `json:"type,omitempty"`
	MediaType    string `json:"mediaType,omitempty"`
	UUIDHref     string `json:"uuidHref,omitempty"`
	DownloadHref string `json:"downloadHref,omitempty"`
	Size         int    `json:"size,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}
type Row struct {
	Meta      Meta   `json:"meta,omitempty"`
	ID        string `json:"id,omitempty"`
	AccountID string `json:"accountId,omitempty"`
	Owner     struct {
		Meta Meta `json:"meta,omitempty"`
	} `json:"owner,omitempty"`
	Shared bool `json:"shared,omitempty"`
	Group  struct {
		Meta Meta `json:"meta,omitempty"`
	} `json:"group,omitempty"`
	Updated      string `json:"updated,omitempty"`
	Name         string `json:"name,omitempty"`
	ExternalCode string `json:"externalCode,omitempty"`
	Archived     bool   `json:"archived,omitempty"`
	Created      string `json:"created,omitempty"`
	UID          string `json:"uid,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	FullName     string `json:"fullName,omitempty"`
	ShortFio     string `json:"shortFio,omitempty"`
	Inn          string `json:"inn,omitempty"`
	Position     string `json:"position,omitempty"`
	Cashiers     []struct {
		Meta Meta `json:"meta,omitempty"`
	} `json:"cashiers,omitempty"`
}
