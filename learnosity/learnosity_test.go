package learnosity

import (
	"encoding/json"
	"fmt"
	"github.com/jgroeneveld/trial/assert"
	"github.com/nsf/jsondiff"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	ConsumerKey = os.Getenv("LEARNOSITY_CONSUMER_KEY")
	Domain = os.Getenv("LEARNOSITY_DOMAIN")
	ConsumerSecret = os.Getenv("LEARNOSITY_CONSUMER_SECRET")

	time.Local = time.UTC
}

func TestAuth(t *testing.T) {
	t.Run("Test 2022-05-23 3:38PM", func(t *testing.T) {
		expected := `{"security":{"consumer_key":"PUT_YOUR_KEY_HERE","domain":"localhost","timestamp":"20220523-2038","user_id":"PUT_YOUR_USER_ID_HERE","signature":"PUT_EXPECTED_SIGNATURE_HERE"},"request":{"user_id":"PUT_YOUR_USER_ID_HERE","activity_template_id":"quickstart_examples_activity_template_001","session_id":"PUT_YOUR_SESSION_ID_HERE","activity_id":"quickstart_examples_activity_001","rendering_type":"assess","type":"submit_practice","name":"Items API Quickstart","state":"initial"}}`
		req := BuildLearnosityReq{overrideTimestamp: "20220523-2038", overrideUserID: "PUT_YOUR_USER_ID_HERE", overrideSessionID: "PUT_YOUR_SESSION_ID_HERE"}
		res, _, err := BuildLearnosityRequest(&req)
		assert.Nil(t, err)
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		options := jsondiff.DefaultConsoleOptions()
		diff, output := jsondiff.Compare([]byte(expected), resBytes, &options)
		fmt.Println(fmt.Sprintf("Difference %+v", diff))
		fmt.Println(fmt.Sprintf("Output %+v", output))
	})

	t.Run("Test 2022-05-23 3:39PM", func(t *testing.T) {
		expected := `{"security":{"consumer_key":"PUT_YOUR_KEY_HERE","domain":"localhost","timestamp":"20220523-2039","user_id":"PUT_YOUR_USER_ID_HERE","signature":"PUT_EXPECTED_SIGNATURE_HERE"},"request":{"user_id":"PUT_YOUR_USER_ID_HERE","activity_template_id":"quickstart_examples_activity_template_001","session_id":"PUT_YOUR_SESSION_ID_HERE","activity_id":"quickstart_examples_activity_001","rendering_type":"assess","type":"submit_practice","name":"Items API Quickstart","state":"initial"}}`
		req := BuildLearnosityReq{overrideTimestamp: "20220523-2039", overrideUserID: "PUT_YOUR_USER_ID_HERE", overrideSessionID: "PUT_YOUR_SESSION_ID_HERE"}
		res, _, err := BuildLearnosityRequest(&req)
		assert.Nil(t, err)
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		options := jsondiff.DefaultConsoleOptions()
		diff, output := jsondiff.Compare([]byte(expected), resBytes, &options)
		fmt.Println(fmt.Sprintf("Difference %+v", diff))
		fmt.Println(fmt.Sprintf("Output %+v", output))
		fmt.Println(expected)
	})

	t.Run("Test 2022-05-23 3:40PM", func(t *testing.T) {
		expected := `{"security":{"consumer_key":"PUT_YOUR_KEY_HERE","domain":"localhost","timestamp":"20220523-2040","user_id":"PUT_YOUR_USER_ID_HERE","signature":"PUT_EXPECTED_SIGNATURE_HERE"},"request":{"user_id":"PUT_YOUR_USER_ID_HERE","activity_template_id":"quickstart_examples_activity_template_001","session_id":"PUT_YOUR_SESSION_ID_HERE","activity_id":"quickstart_examples_activity_001","rendering_type":"assess","type":"submit_practice","name":"Items API Quickstart","state":"initial"}}`
		req := BuildLearnosityReq{overrideTimestamp: "20220523-2040", overrideUserID: "PUT_YOUR_USER_ID_HERE", overrideSessionID: "PUT_YOUR_SESSION_ID_HERE"}
		res, _, err := BuildLearnosityRequest(&req)
		assert.Nil(t, err)
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		options := jsondiff.DefaultConsoleOptions()
		diff, output := jsondiff.Compare([]byte(expected), resBytes, &options)
		fmt.Println(fmt.Sprintf("Difference %+v", diff))
		fmt.Println(fmt.Sprintf("Output %+v", output))
		fmt.Println(expected)
	})

	t.Run("Test 2022-05-23 3:41PM", func(t *testing.T) {
		expected := `{"security":{"consumer_key":"PUT_YOUR_KEY_HERE","domain":"localhost","timestamp":"20220523-2041","user_id":"PUT_YOUR_USER_ID_HERE","signature":"PUT_EXPECTED_SIGNATURE_HERE"},"request":{"user_id":"PUT_YOUR_USER_ID_HERE","activity_template_id":"quickstart_examples_activity_template_001","session_id":"PUT_YOUR_SESSION_ID_HERE","activity_id":"quickstart_examples_activity_001","rendering_type":"assess","type":"submit_practice","name":"Items API Quickstart","state":"initial"}}`
		req := BuildLearnosityReq{overrideTimestamp: "20220523-2041", overrideUserID: "PUT_YOUR_USER_ID_HERE", overrideSessionID: "PUT_YOUR_SESSION_ID_HERE"}
		res, _, err := BuildLearnosityRequest(&req)
		assert.Nil(t, err)
		resBytes, err := json.Marshal(res)
		assert.Nil(t, err)

		options := jsondiff.DefaultConsoleOptions()
		diff, output := jsondiff.Compare([]byte(expected), resBytes, &options)
		fmt.Println(fmt.Sprintf("Difference %+v", diff))
		fmt.Println(fmt.Sprintf("Output %+v", output))
		fmt.Println(expected)
	})
}
