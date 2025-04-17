package developer

type Developer struct {
	DeveloperName    string   `json:"developer_name"`
	Email            string   `json:"email"`
	Skills           []string `json:"skills"`
	ExperienceYears  int      `json:"experience_years"`
	AssignedProjects []string `json:"assigned_projects"`
}
