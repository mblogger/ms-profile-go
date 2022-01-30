package model

type Profile struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
}

type Contact struct {
	HouseNumber       string `json:"house_number,omitempty"`
	Street            string `json:"street,omitempty"`
	Address1          string `json:"address_1,omitempty"`
	Address2          string `json:"address_2,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	PinCode           int    `json:"pin_code,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Country           string `json:"country,omitempty"`
	IsSameAsPermanent bool   `json:"is_same_as_permanent,omitempty"`
}

type Experience struct {
	StartDate      string   `json:"start_date,omitempty"`
	EndDate        string   `json:"end_date,omitempty"`
	OrgName        string   `json:"org_name,omitempty"`
	OrgContact     Contact  `json:"org_contact,omitempty"`
	OrgDescription string   `json:"org_description,omitempty"`
	OrgDetailList  []string `json:"org_detail_list,omitempty"`
}
