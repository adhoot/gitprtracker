package gitapi

import (
	"net/http"
	"fmt"
	"encoding/json"
)

const PullRequestsUrl = "https://git.soma.salesforce.com/api/v3/repos/csc-health/report-collector/pulls"

func SearchPRs() (*PRSeachResult, error) {
	resp, err := http.Get(PullRequestsUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result PRSeachResult
	if err := json.NewDecoder(resp.Body).Decode(&result.Items); err != nil {
		resp.Body.Close()
		return nil, err
	}

	result.TotalCount = len(result.Items)

	resp.Body.Close()
	return &result, nil
}