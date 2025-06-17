package src

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const mincraftGameID = "j1npme6p"
const anyPercentGlitchlessCategoryID = "mkeyl926"

var ErrNoRunsFound = errors.New("no runs found")

type API struct {
	UserID string
}

func New(user string) (*API, error) {
	userID, err := getUserID(user)
	if err != nil {
		return nil, fmt.Errorf("could not get user id: %w", err)
	}
	return &API{UserID: userID}, nil
}

func (api *API) GetRun() (string, int, float64, error) {
	parameters := map[string]string{
		"userId": api.UserID,
	}
	encodedParameters, err := encodeParameters(parameters)
	if err != nil {
		return "", 0, 0, fmt.Errorf("encode parameters failed: %w", err)
	}
	response, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v2/GetUserLeaderboard?_r=%s", encodedParameters))
	if err != nil {
		return "", 0, 0, fmt.Errorf("get user summary failed: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", 0, 0, errors.New("get user summary failed with status: " + response.Status)
	}
	getUserLeaderboard := GetUserLeaderboardResponse{}
	err = json.NewDecoder(response.Body).Decode(&getUserLeaderboard)
	if err != nil {
		return "", 0, 0, fmt.Errorf("decode response body failed: %w", err)
	}
	if len(getUserLeaderboard.Runs) == 0 {
		return "", 0, 0, ErrNoRunsFound
	}
	for _, run := range getUserLeaderboard.Runs {
		if run.Obsolete || run.GameID != mincraftGameID || run.CategoryID != anyPercentGlitchlessCategoryID {
			continue
		}
		return run.ID, run.Place, run.Igt, nil
	}
	return "", 0, 0, ErrNoRunsFound
}

func encodeParameters(parameters map[string]string) (string, error) {
	json, err := json.Marshal(parameters)
	if err != nil {
		return "", fmt.Errorf("encode json failed: %w", err)
	}
	base64Encoded := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(json)
	return base64Encoded, nil
}

func getUserID(user string) (string, error) {
	parameters := map[string]string{
		"url": user,
	}
	encodedParameters, err := encodeParameters(parameters)
	if err != nil {
		return "", fmt.Errorf("encode parameters failed: %w", err)
	}
	response, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v2/GetUserSummary?_r=%s", encodedParameters))
	if err != nil {
		return "", fmt.Errorf("get user summary failed: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", errors.New("get user summary failed with status: " + response.Status)
	}
	getUserSummary := GetUserSummaryResponse{}
	err = json.NewDecoder(response.Body).Decode(&getUserSummary)
	if err != nil {
		return "", fmt.Errorf("decode response body failed: %w", err)
	}
	return getUserSummary.User.ID, nil
}
