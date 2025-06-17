package src

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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

func (s *API) GetRank(user string) string {
	return "100"
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
