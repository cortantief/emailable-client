package emailableclient

import "net/url"

// https://emailable.com/docs/api/#emails
type UniqueOption struct {
	Smtp      bool
	AcceptAll bool
	Timeout   uint64 // time in seconds
}

// https://emailable.com/docs/api/#verify-a-batch-of-emails
type BatchOption struct {

	// A URL that will receive the batch results via HTTP POST.
	Url *url.URL

	// A comma separated list of fields to include in the response.
	// If nothing is specified, all fields will be returned.
	// Valid fields are accept_all, did_you_mean, disposable, domain, email, first_name, free, full_name, gender, last_name, mx_record, reason, role, score, smtp_provider, state, tag, and user.
	ResponseFields []ResponseField

	// Defaults to true.
	// Retries increase accuracy by automatically retrying verification when our system receives certain responses from mail servers.
	// To speed up verification, you can disable this by setting retries to false; however, doing so may increase the number of unknown responses.
	Retries bool

	// A boolean value indicating whether to include partial results when a batch is still verifying.
	// This option is only available for batches with up to 1,000 emails. Defaults to false.
	Partial bool
	// Used to simulate certain responses from the API while using a test key. Valid options are generic_error, insufficient_credits_error, payment_error, and card_error.
	simulate *BatchSimulate
}

func (o *BatchOption) AddSimulate(s BatchSimulate) {
	o.simulate = &s
}

func (o *BatchOption) RemoveSimulate() {
	o.simulate = nil
}

type BatchStatusOption struct {
	// A boolean value indicating whether to include partial results when a batch is still verifying.
	//This option is only available for batches with up to 1,000 emails. Defaults to false.
	Partial bool

	// Used to simulate certain responses from the API while using a test key.
	// Valid options are generic_error, importing, verifying, and paused.
	simulate *StatusSimulate
}

func (o *BatchStatusOption) AddSimulate(s StatusSimulate) {
	o.simulate = &s
}

func (o *BatchStatusOption) RemoveSimulate() {
	o.simulate = nil
}
