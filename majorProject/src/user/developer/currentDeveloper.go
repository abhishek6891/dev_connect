package developer

type Developer struct {
	Name          string   `json:"name,omitempty"`
	Phone         string   `json:"phone,omitempty"`
	Password      string   `json:"password,omitempty"`
	Email         string   `json:"email,omitempty"`
	Age           int      `json:"age,omitempty"`
	Gender        string   `json:"gender,omitempty"`
	Qualification []string `json:"qualification,omitempty"`
	Skills        []string `json:"skills,omitempty"`
	ProjectsDone  []string `json:"projects_done,omitempty"`
	Experience    string   `json:"experience,omitempty"`
	Availability  string   `json:"availability,omitempty"`
	RatePerHour   float32  `json:"rate_per_hour,omitempty"`
}
