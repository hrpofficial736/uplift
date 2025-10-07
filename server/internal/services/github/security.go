package github

import (
	"fmt"
)

func CheckRepoSecurity(owner string, repo string) ([]interface{}, error) {
	fmt.Println("tool called on mcp server")
	dependabotPath := fmt.Sprintf("/repos/%s/%s/dependabot/alerts", owner, repo)
	dependabotAlertsResponse, err := CallGithubApi(dependabotPath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error from the github api for dependabot alerts: %s", err)
	}

	codeScanningPath := fmt.Sprintf("/repos/%s/%s/code-scanning/alerts", owner, repo)
	codeScanningAlertsResponse, err := CallGithubApi(codeScanningPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while code scanning: %s", err)
	}

	secretScanningPath := fmt.Sprintf("/repos/%s/%s/code-scanning/alerts", owner, repo)
	secretScanningAlertsResponse, err := CallGithubApi(secretScanningPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while secret scanning: %s", err)
	}

	responses := []interface{}{dependabotAlertsResponse, codeScanningAlertsResponse, secretScanningAlertsResponse}
	fmt.Printf("security responses in services: %v", responses)
	return responses, nil
}
