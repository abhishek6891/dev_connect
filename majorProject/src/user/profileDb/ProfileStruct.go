package profileDb

type ProfileUser struct {
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	Qualification []string `json:"qualification"`
	Skills        []string `json:"skills"`
	ProjectsDone  []string `json:"projects_done"`
	Experience    string   `json:"experience"`
}
