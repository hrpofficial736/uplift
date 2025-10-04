package github

import (
	"fmt"

	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func CheckRepoSecurity(owner string, repo string) ([]interface{}, error) {
	dependabotPath := fmt.Sprintf("/repos/%s/%s/dependabot/alerts", owner, repo)
	dependabotAlertsResponse, err := utils.CallGithubApi(dependabotPath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error from the github api for dependabot alerts: %s", err)
	}

	codeScanningPath := fmt.Sprintf("/repos/%s/%s/code-scanning/alerts", owner, repo)
	codeScanningAlertsResponse, err := utils.CallGithubApi(codeScanningPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while code scanning: %s", err)
	}

	secretScanningPath := fmt.Sprintf("/repos/%s/%s/code-scanning/alerts", owner, repo)
	secretScanningAlertsResponse, err := utils.CallGithubApi(secretScanningPath, "GET")

	if err != nil {
		return nil, fmt.Errorf("error from the github api while secret scanning: %s", err)
	}

	responses := []interface{}{dependabotAlertsResponse, codeScanningAlertsResponse, secretScanningAlertsResponse}
	return responses, nil
}
