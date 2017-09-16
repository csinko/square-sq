package main


type Webhook struct{

	Commits []struct{

		commitID string `json:"id"`
		commitMessage string `json:"message"`
		commitTimestamp time.Time `json:"timestamp"`

		Author struct {
			
			authorEmail    string `json:"email"`
			authorName     string `json:"name"`
			authorUsername string `json:"username"`
		
		} `json:"author"`


	} `json:"committer"`

	Repository struct{

		repoName  string `json:"name"`
		repoID int64 `json:"id"`

		Owner struct {

			Name string `json:"Name"`

		} `json:"owner"`



	} `json:"repository`"

}
