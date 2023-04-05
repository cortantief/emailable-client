package emailableclient

type Emailable struct {
	apiKey string
}

func NewEmailable(apiKey string) Emailable {
	return Emailable{
		apiKey: apiKey,
	}
}

func (e Emailable) GetApiKey() string {
	return e.apiKey
}

func (e *Emailable) SetApiKey(key string) {
	e.apiKey = key
}

// Verify a single email. If a verification request takes longer than the timeout, you may retry this request for up to 5 minutes.
// After 5 minutes, further requests will count against your usage.
// The verification result will be returned when it is available.
// When a test key is used, a random sample response will be returned.
func (e Emailable) VerifyEmailWithOption(email string, o UniqueOption) (*VerifyEmailResult, error) {
	return verifyEmail(email, e.apiKey, o)
}

// Verify a single email. If a verification request takes longer than the timeout, you may retry this request for up to 5 minutes.
// After 5 minutes, further requests will count against your usage.
// The verification result will be returned when it is available.
func (e Emailable) VerifyEmail(email string) (*VerifyEmailResult, error) {
	return verifyEmail(email, e.apiKey, defaultUniqueOption())
}

// Verify a batch of emails. The emails should be sent as a parameter emails and should be separated by commas. Up to 50,000 emails can be sent per batch.
// For enterprise accounts, up to 1,000,000 emails can be sent per batch.
func (e Emailable) VerifyBatch(emails []string) (*BatchResult, error) {
	return verifyBatch(e.apiKey, emails, defaultBatchOption())
}

// Verify a batch of emails. The emails should be sent as a parameter emails and should be separated by commas. Up to 50,000 emails can be sent per batch.
// For enterprise accounts, up to 1,000,000 emails can be sent per batch.
func (e Emailable) VerifyBatchWithOption(emails []string, o BatchOption) (*BatchResult, error) {
	return verifyBatch(e.apiKey, emails, o)
}

// Get general account information like the email of the account owner and available credits.
func (e Emailable) GetAccountInfo() (*AccountResult, error) {
	return getAccountInfo(e.apiKey)
}

// GET requests to the batch endpoint will get the current status of the batch verification job specified in the id parameter.
// When a credit card transaction is necessary to obtain enough credits to verify a batch, billing related messages will be returned if there is an error.
// These will be sent with a 402 response code.
// When a test key is used, a random sample response will be returned for each email in the batch.
// Additionally, it is possible to simulate certain API responses when using a test key by utilizing the simulate parameter.
func (e Emailable) GetBatchStatus(id string) (*BatchStatusResult, error) {
	return getBatchStatus(e.apiKey, id, defaultBatchStatusOption())
}
func (e Emailable) GetBatchStatusWithOption(id string, o BatchStatusOption) (*BatchStatusResult, error) {
	return getBatchStatus(e.apiKey, id, o)
}
