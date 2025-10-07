package github

import (
	"fmt"
)

func CheckRepoMaintainability(owner string, repo string) ([]interface{}, error) {
	basicHealthPath := fmt.Sprintf("/repos/%s/%s", owner, repo)
	basicHealthResponse, err := CallGithubApi(basicHealthPath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking health in maintainability: %s", err)
	}

	contributorsPath := fmt.Sprintf("/repos/%s/%s/stats/contributors", owner, repo)
	contributorsResponse, err := CallGithubApi(contributorsPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking contributors in maintainability: %s", err)
	}

	commitPath := fmt.Sprintf("/repos/%s/%s/stats/commit_activity", owner, repo)
	commitResponse, err := CallGithubApi(commitPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking commits in maintainability: %s", err)
	}

	codeFrequencyPath := fmt.Sprintf("/repos/%s/%s/stats/code_frequency", owner, repo)
	codeFrequencyResponse, err := CallGithubApi(codeFrequencyPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking code frequency in maintainability: %s", err)
	}

	participationPath := fmt.Sprintf("/repos/%s/%s/stats/participation", owner, repo)
	participationResponse, err := CallGithubApi(participationPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking participation in maintainability: %s", err)
	}

	issuesPath := fmt.Sprintf("/repos/%s/%s/stats/issues?state=closed", owner, repo)
	issuesResponse, err := CallGithubApi(issuesPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking issues in maintainability: %s", err)
	}

	pullRequestsPath := fmt.Sprintf("/repos/%s/%s/stats/pulls?state=open", owner, repo)
	pullRequestsResponse, err := CallGithubApi(pullRequestsPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while checking pull requests in maintainability: %s", err)
	}

	responses := []interface{}{basicHealthResponse, contributorsResponse, commitResponse, codeFrequencyResponse, participationResponse, issuesResponse, pullRequestsResponse}
	return responses, nil
}
