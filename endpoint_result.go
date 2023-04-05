package emailableclient

type VerifyEmailResult struct {
	// Whether the mail server used to verify indicates that all addresses are deliverable regardless of whether or not the email is valid.
	AcceptAll bool `json:"accept_all,omitempty"`

	// A suggested correction for a common misspelling.
	DidYouMean string `json:"did_you_mean,omitempty"`

	// Whether this email is hosted on a disposable or temporary email service.
	Disposable bool `json:"disposable"`

	// The domain of the email. (e.g. The domain for john.smith@gmail.com would be gmail.com)
	Domain string `json:"domain"`

	// The length of time (in seconds) spent verifying this email.
	Duration float64 `json:"duration"`

	// The email that was verified.
	Email string `json:"email"`

	// The possible first name of the user.
	FirstName string `json:"first_name,omitempty"`

	// Whether the email is hosted by a free email provider.
	Free bool `json:"free"`

	// The possible full name of the user.
	FullName string `json:"full_name,omitempty"`

	// The possible gender of the user.
	Gender string `json:"gender,omitempty"`

	// The possible last name of the user.
	LastName string `json:"last_name,omitempty"`

	// The mailbox is currently full and emails may not be delivered.
	MailboxFull bool `json:"mailbox_full"`

	// The address of the mail server used to verify the email.
	MxRecord string `json:"mx_record,omitempty"`

	// An address that indicates it should not be replied to.
	NoReply bool `json:"no_reply"`

	// The reason for the associated
	Reason string `json:"reason,omitempty"`

	// Whether the email is considered a role address. (e.g. support, info, etc.)
	Role bool `json:"role"`

	// The score of the verified email.
	Score int `json:"score"`

	// The SMTP provider of the verified email's domain.
	SmtpProvider string `json:"smtp_provider,omitempty"`

	// The state of the verified email. (e.g. deliverable, undeliverable, risky, unknown)
	State string `json:"state"`

	// The tag part of the verified email. (e.g. The tag for john.smith+example@gmail.com would be example)
	Tag string `json:"tag,omitempty"`

	// The user part of the verified email. (e.g. The user for john.smith@gmail.com would be john.smith)
	User string `json:"user,omitempty"`
}

// https://emailable.com/docs/api/#get-account-info
type AccountResult struct {
	// The email of the account owner.
	OwnerEmail string `json:"owner_email"`

	// The amount of credits remaining on the account.
	AvailableCredit uint `json:"available_credits"`
}

// https://emailable.com/docs/api/#verify-a-batch-of-emails
type BatchResult struct {
	// A message about your batch.
	Message string `json:"message"`

	// The unique ID of the batch.
	Id string `json:"id"`
}

// A hash with one key per possible reason attribute. The values are integers representing the number of emails with that reason.
type ReasonCounts struct {
	AcceptedEmail     uint64 `json:"accepted_email"`
	InvalidDomain     uint64 `json:"invalid_domain"`
	InvalidEmail      uint64 `json:"invalid_email"`
	InvalidSmtp       uint64 `json:"invalid_smtp"`
	LowDeliverability uint64 `json:"low_deliverability"`
	LowQuality        uint64 `json:"low_quality"`
	NoConnect         uint64 `json:"no_connect"`
	RejectedEmail     uint64 `json:"rejected_email"`
	Timeout           uint64 `json:"timeout"`
	UnavailableSmtp   uint64 `json:"unavailable_smtp"`
	UnexpectedError   uint64 `json:"unexpected_error"`
}

// A hash with one key per possible state attribute.
// The values are integers representing the number of emails with that state.
// In addition to the state keys, total_counts also contains keys processed and total, with values indicating the number of emails in the batch.
type TotalCounts struct {
	Deliverable   uint64 `json:"deliverable"`
	Processed     uint64 `json:"processed"`
	Risky         uint64 `json:"risky"`
	Total         uint64 `json:"total"`
	Undeliverable uint64 `json:"undeliverable"`
	Unknown       uint64 `json:"unknown"`
}

type BatchStatusResult struct {
	// A message about your batch.
	Message string `json:"message"`

	// The number of emails that have been verified in the batch.
	Processed uint64 `json:"processed,omitempty"`

	// The total number of emails in your batch.
	Total uint64 `json:"total,omitempty"`

	// An array containing responses for each email in the batch.
	// This field will only be returned for batches up to 1,000 emails.
	// (See [single email verification]: https://emailable.com/docs/api/#verify-an-email for more information on the response fields.)
	Emails       []VerifyEmailResult `json:"emails,omitempty"`
	DownloadFile string              `json:"download_file,omitempty"`
	Id           string              `json:"id,omitempty"`
	ReasonCounts ReasonCounts        `json:"reason_counts,omitempty"`
	TotalCounts  TotalCounts         `json:"total_counts,omitempty"`
}
