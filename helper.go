package emailableclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	emailField          = "email"
	multipleEmailField  = "emails"
	urlField            = "url"
	apiKeyField         = "api_key"
	fieldsResponseField = "response_fields"
	retriesField        = "retries"
	simulateField       = "simulate"
	timeoutField        = "timeout"
	acceptAllField      = "accept_all"
	smtpField           = "smtp"
	idField             = "id"
	partialField        = "partial"
)

var creatingUrlErr = errors.New("Error while creating url, shouldn't happen")

func validateStatusCode(statusCode int) EmailableError {
	if statusCode == http.StatusOK {
		return nil
	}

	if statusCode == 249 {
		return TimeoutError
	} else if statusCode == http.StatusBadRequest {
		return BadRequest
	} else if statusCode == http.StatusUnauthorized {
		return ApiKeyMissing
	} else if statusCode == http.StatusPaymentRequired {
		return LowCredit
	} else if statusCode == http.StatusForbidden {
		return InvalidApiKey
	} else if statusCode == http.StatusNotFound {
		return NotFound
	} else if statusCode == http.StatusTooManyRequests {
		return TooManyRequest
	} else if statusCode == http.StatusInternalServerError {
		return InternalServerError
	} else if statusCode == http.StatusServiceUnavailable {
		return ServiceUnavailable
	}

	return UnknownStatusCode
}

func defaultUniqueOption() UniqueOption {
	return UniqueOption{
		Smtp:      true,
		AcceptAll: false,
		Timeout:   5,
	}
}

func defaultBatchOption() BatchOption {
	return BatchOption{
		Url:            nil,
		ResponseFields: []ResponseField{},
		Retries:        true,
		Partial:        false,
		simulate:       nil,
	}
}

func defaultBatchStatusOption() BatchStatusOption {
	return BatchStatusOption{
		Partial:  false,
		simulate: nil,
	}
}

func createUniqueUrl(apiKey string, email string, o UniqueOption) *url.URL {
	u, err := url.Parse("https://api.emailable.com/v1/verify")

	if err != nil {
		return nil
	}

	query := u.Query()
	query.Add(emailField, email)
	query.Add(smtpField, strconv.FormatBool(o.Smtp))
	query.Add(acceptAllField, strconv.FormatBool(o.AcceptAll))
	query.Add(timeoutField, strconv.FormatUint(o.Timeout, 10))
	query.Add(apiKeyField, apiKey)
	u.RawQuery = query.Encode()

	return u
}

func verifyEmail(email string, apiKey string, o UniqueOption) (*VerifyEmailResult, error) {
	url := createUniqueUrl(apiKey, email, o)
	if url == nil {
		return nil, creatingUrlErr
	}

	result := VerifyEmailResult{}
	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := validateStatusCode(resp.StatusCode); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func addFormField(w *multipart.Writer, fname string, value []byte) error {
	f, err := w.CreateFormField(fname)

	if err != nil {
		return err
	}

	if _, err := f.Write(value); err != nil {
		return err
	}

	return nil
}

func setBatchWriter(writer *multipart.Writer, apiKey string, emails []string, o BatchOption) error {

	if err := writer.WriteField(apiKeyField, apiKey); err != nil {
		return err
	}

	if err := writer.WriteField(retriesField, strconv.FormatBool(o.Retries)); err != nil {
		return err
	}

	if o.Url != nil {
		if err := writer.WriteField(urlField, o.Url.String()); err != nil {
			return err
		}
	}

	if o.simulate != nil {
		if err := writer.WriteField(simulateField, string(*o.simulate)); err != nil {
			return err
		}
	}

	if len(o.ResponseFields) > 0 {
		fields := make([]string, 0, len(o.ResponseFields))
		for i := range o.ResponseFields {
			fields = append(fields, string(o.ResponseFields[i]))
		}

		if err := writer.WriteField(fieldsResponseField, strings.Join(fields, ",")); err != nil {
			return err
		}
	}

	if err := writer.WriteField(multipleEmailField, strings.Join(emails, ",")); err != nil {
		return err
	}

	return nil
}

func createBatchRequest(apiKey string, emails []string, o BatchOption) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := setBatchWriter(writer, apiKey, emails, o)
	writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.emailable.com/v1/batch", body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

func verifyBatch(apiKey string, emails []string, o BatchOption) (*BatchResult, error) {
	client := &http.Client{}
	req, err := createBatchRequest(apiKey, emails, o)
	if err != nil {
		return nil, err
	}

	result := BatchResult{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = validateStatusCode(resp.StatusCode); err != nil {
		return nil, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func getAccountInfo(apiKey string) (*AccountResult, error) {
	u, err := url.Parse("https://api.emailable.com/v1/account")

	if err != nil {
		return nil, creatingUrlErr
	}

	q := u.Query()
	q.Add(apiKeyField, apiKey)
	u.RawQuery = q.Encode()
	result := AccountResult{}
	resp, err := http.Get(u.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = validateStatusCode(resp.StatusCode); err != nil {
		return nil, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func getBatchStatus(apikey string, id string, o BatchStatusOption) (*BatchStatusResult, error) {
	u, err := url.Parse("https://api.emailable.com/v1/batch")

	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add(apiKeyField, apikey)
	q.Add(idField, id)
	q.Add(partialField, strconv.FormatBool(o.Partial))

	if o.simulate != nil {
		q.Add(simulateField, string(*o.simulate))
	}

	u.RawQuery = q.Encode()
	result := BatchStatusResult{}
	resp, err := http.Get(u.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
