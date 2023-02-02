package learnosity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

var ConsumerKey = ""
var ConsumerSecret = ""
var Domain = ""

type AuthRes struct {
	Security Security `json:"security,omitempty"`
	Request  Request  `json:"request,omitempty"`
}

type BuildLearnosityReq struct {
	overrideTimestamp string // field to allow debugging BuildLearnosityRequest output
	overrideSessionID string // field to allow debugging BuildLearnosityRequest output
	overrideUserID    string // field to allow debugging BuildLearnosityRequest output
}

type Security struct {
	ConsumerKey string `json:"consumer_key,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Signature   string `json:"signature,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	UserID      string `json:"user_id,omitempty"`
}

type Request struct {
	UserID             string `json:"user_id,omitempty"`
	ActivityTemplateID string `json:"activity_template_id,omitempty"`
	SessionID          string `json:"session_id,omitempty"`
	ActivityID         string `json:"activity_id,omitempty"`
	RenderingType      string `json:"rendering_type,omitempty"`
	Type               string `json:"type,omitempty"`
	Name               string `json:"name,omitempty"`
	State              string `json:"state,omitempty"`
}

func BuildLearnosityRequest(bReq *BuildLearnosityReq) (*AuthRes, int, error) {
	timestamp := GetTimestamp(bReq.overrideTimestamp)
	hashComponents := []string{ConsumerKey, Domain, timestamp, bReq.overrideUserID, ConsumerSecret}

	authResponse := GetAuthResponse(timestamp, bReq.overrideSessionID, bReq.overrideUserID, hashComponents)

	return authResponse, 200, nil
}

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func GetTimestamp(timestampOverride string) string {
	if timestampOverride != "" {
		return timestampOverride
	}

	return time.Now().Format("yyyymmdd-hhmm")
}

func GetAuthResponse(timestamp, overrideSessionID, overrideUserID string, hashComponents []string) *AuthRes {
	if overrideSessionID == "" {
		overrideSessionID = uuid.New().String()
	}

	if overrideUserID == "" {
		overrideUserID = uuid.New().String()
	}

	request := Request{
		UserID:             overrideUserID,
		ActivityTemplateID: "quickstart_examples_activity_template_001",
		SessionID:          overrideSessionID,
		ActivityID:         "quickstart_examples_activity_001",
		RenderingType:      "assess",
		Type:               "submit_practice",
		Name:               "Items API Quickstart",
		State:              "initial",
	}

	requestJSONBytes, _ := json.Marshal(request)

	hashComponents = append(hashComponents, string(requestJSONBytes))
	signature := strings.Join(hashComponents, "_")

	result := NewSHA256([]byte(signature))
	signature = hex.EncodeToString(result)

	return &AuthRes{
		Security: Security{
			ConsumerKey: os.Getenv("LEARNOSITY_CONSUMER_KEY"),
			Domain:      os.Getenv("LEARNOSITY_DOMAIN"),
			Signature:   signature,
			Timestamp:   timestamp,
			UserID:      overrideUserID,
		},
		Request: request,
	}
}
