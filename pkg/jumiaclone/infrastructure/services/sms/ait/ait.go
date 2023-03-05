package ait

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/application"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"
	"github.com/sirupsen/logrus"
)

const (
	AITAuthenticationError = "The supplied authentication is invalid"
)

type AITService struct {
	Env string
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}

// GetAPIHost returns either sandbox or prod
func GetAPIHost(env string) string {
	return getHost(env, "api")
}

func getHost(env, service string) string {
	if env != "sandbox" {
		return fmt.Sprintf("https://%s.africastalking.com", service)
	}
	return fmt.Sprintf(
		"https://%s.sandbox.africastalking.com",
		service,
	)

}

func NewAITService() *AITService {
	env := application.MustGetEnvVar("AIT_ENVIRONMENT")
	return &AITService{
		Env: env,
	}
}

func (a *AITService) SendSMS(message, to string) (*dto.SendMessageResponse, error) {
	values := url.Values{}
	values.Set("username", application.MustGetEnvVar("AIT_ENVIRONMENT"))
	values.Set("to", to)
	values.Set("message", message)

	smsURL := GetSmsURL(a.Env)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, err := a.newPostRequest(smsURL, values, headers, application.MustGetEnvVar("AIT_API_KEY"))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// read the response body to a variable
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	smsMessageResponse := &dto.SendMessageResponse{}

	bodyAsString := string(bodyBytes)
	if strings.Contains(bodyAsString, AITAuthenticationError) {
		// return so that other processes don't break
		log.Println("AIT Authentication error encountered")
		return smsMessageResponse, nil
	}

	// reset the response body to the original unread state so that decode can
	// continue
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := json.NewDecoder(res.Body).Decode(smsMessageResponse); err != nil {
		return nil, errors.New(
			fmt.Errorf("SMS Error : unable to parse sms response ; %v", err).
				Error(),
		)
	}
	logrus.Print("message", smsMessageResponse.SMSMessageData.Message)

	return smsMessageResponse, nil

}

func (a *AITService) newPostRequest(
	url string,
	values url.Values,
	headers map[string]string,
	key string,
) (*http.Response, error) {
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", key)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}
