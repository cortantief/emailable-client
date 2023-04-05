package emailableclient

import "errors"

type ResponseField string

const (
	AcceptAll    ResponseField = "accept_all"
	DidYouMean   ResponseField = "did_you_mean"
	Disposable   ResponseField = "disposable"
	Domain       ResponseField = "domain"
	Email        ResponseField = "email"
	FirstName    ResponseField = "first_name"
	Free         ResponseField = "free"
	FullName     ResponseField = "full_name"
	Gender       ResponseField = "gender"
	LastName     ResponseField = "last_name"
	MxRecord     ResponseField = "mx_record"
	Reason       ResponseField = "reason"
	Role         ResponseField = "role"
	Score        ResponseField = "score"
	SmtpProvider ResponseField = "smtp_provider"
	State        ResponseField = "state"
	Tag          ResponseField = "tag"
	User         ResponseField = "user"
)

type BatchSimulate string

const (
	GenericError             BatchSimulate = "generic_error"
	InsufficientCreditsError BatchSimulate = "insufficient_credits_error"
	PaymentError             BatchSimulate = "payment_error"
	CardError                BatchSimulate = "card_error"
)

type StatusSimulate string

const (
	StatusGenericError StatusSimulate = "generic_error"
	StatusImporting    StatusSimulate = "importing"
	StatusVerifying    StatusSimulate = "verifying"
	StatusPaused       StatusSimulate = "paused"
)

type EmailableError error

var (

	// Error based on status code, Url for more information https://emailable.com/docs/api/#status-codes
	TimeoutError        EmailableError = errors.New("The specified resource does not exist.")
	BadRequest          EmailableError = errors.New("Your request is structured incorrectly.")
	ApiKeyMissing       EmailableError = errors.New("Your request is structured incorrectly.")
	LowCredit           EmailableError = errors.New("You don't have enough credits to complete this request.")
	InvalidApiKey       EmailableError = errors.New("Your API key is invalid.")
	NotFound            EmailableError = errors.New("The specified resource does not exist.")
	UnknownStatusCode   EmailableError = errors.New("unknown status code")
	TooManyRequest      EmailableError = errors.New("You're requesting an endpoint too often.")
	InternalServerError EmailableError = errors.New("A server error occurred. Please try again later, or contact support if you're having trouble.")
	ServiceUnavailable  EmailableError = errors.New("We're temporarily offline for maintenance. Please try again later.")

	// Emailbased on usage
	MoreEmail EmailableError = errors.New("Please send more than one email.")
)

type BatchState uint

const (
	Importing BatchState = iota
	Verifying
	Paused
)

func (s BatchState) String() string {
	if s == Paused {
		return "Your batch is paused. Please contact support at <support@emailable.com>."
	} else if s == Importing {
		return "Your batch is being processed."
	} else if s == Verifying {
		return "Your batch is being processed."
	}
	return "Unknown state"
}
