package github

import (
	"fmt"

	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func CheckForQuality(owner string, repo string) (interface{}, error) {
	// documentation checks
	readmePath := fmt.Sprintf("/repos/%s/%s/contents/README.md", owner, repo)
	readmeCheckResponse, err := utils.CallGithubApi(readmePath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error while checking readme in quality: %s", err)
	}

	contributingPath := fmt.Sprintf("/repos/%s/%s/contents/CONTRIBUTING.md", owner, repo)
	contributingResponse, err := utils.CallGithubApi(contributingPath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error while checking contributing in quality: %s", err)
	}

	codeOfConductPath := fmt.Sprintf("/repos/%s/%s/contents/CODE_OF_CONDUCT.md", owner, repo)
	codeOfConductResponse, err := utils.CallGithubApi(codeOfConductPath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error while checking code of conduct in quality: %s", err)
	}

	// license checks
	licensePath := fmt.Sprintf("/repos/%s/%s/license", owner, repo)
	licenseResponse, err := utils.CallGithubApi(licensePath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error while checking license in quality: %s", err)
	}

	// hygiene check
	gitIgnorePath := fmt.Sprintf("/repos/%s/%s/contents/.gitignore", owner, repo)
	gitIgnoreResponse, err := utils.CallGithubApi(gitIgnorePath, "GET")
	if err != nil {
		return nil, fmt.Errorf("error while checking .gitignore in quality: %s", err)
	}

	responses := []interface{}{readmeCheckResponse, contributingResponse, codeOfConductResponse, licenseResponse, gitIgnoreResponse}

	return responses, nil
}
