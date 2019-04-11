package roboat

type GitHubUser struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
}

type GitHubComment struct {
	Id   int        `json:"id"`
	Body string     `json:"body"`
	User GitHubUser `json:"user"`
}

type GitHubIssue struct {
	Id     int        `json:"id"`
	Number int        `json:"number"`
	Title  string     `json:"title"`
	State  string     `json:"state"`
	Body   string     `json:"body"`
	User   GitHubUser `json:"user"`
}

type GitHubRepository struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Owner       GitHubUser `json:"owner"`
	Private     bool       `json:"private"`
	Fork        bool       `json:"fork"`
	Description string     `json:"decription"`
	GitUrl      string     `json:"git_url"`
	SshUrl      string     `json:"ssh_url"`
	CloneUrl    string     `json:"clone_url"`
}

type GitHubIssueCommentPayload struct {
	Action     string           `json:"action"`
	Issue      GitHubIssue      `json:"issue"`
	Comment    GitHubComment    `json:"comment"`
	Repository GitHubRepository `json:"repository"`
	Sender     GitHubUser       `json:"sender"`
}

type Roboat struct {
	GitHubToken string
	BotName     string
}
