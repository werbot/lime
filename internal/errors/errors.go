package errors

import "errors"

const (
	MsgBadRequest      = "bad request"
	MsgNotFound        = "not found"
	MsgWrongPassword   = "wrong password"
	MsgMissingMetadata = "missing metadata"

	// MsgFailedToParseKeyPair     = "failed to parse key pair"
	// MsgFailedToParseCertificate = "failed to parse certificate"
	// MsgFailedToDialServer       = "failed to dial server"
	// MsgTokenIsInvalid           = "token is invalid"

	MsgSettingNotFound = "setting record not found"

	MsgPaymentNotFound = "payment record not found"
	MsgAuditNotFound   = "audit record not found"

	MsgLicenseNotFound        = "license not found"
	MsgLicenseLinkedToPayment = "license is already linked to this payment"

	MsgPatternNotFound   = "pattern not found"
	MsgPatternNotDeleted = "pattern cannot be deleted because there are associated licenses"

	MsgCustomerNotFound   = "customer not found"
	MsgCustomerNotDeleted = "customer cannot be deleted because there are associated payments"
)

var (
	ErrBadRequest      = errors.New(MsgBadRequest)
	ErrNotFound        = errors.New(MsgNotFound)
	ErrWrongPassword   = errors.New(MsgWrongPassword)
	ErrMissingMetadata = errors.New(MsgMissingMetadata)

	// ErrFailedToParseKeyPair     = errors.New(MsgFailedToParseKeyPair)
	// ErrFailedToParseCertificate = errors.New(MsgFailedToParseCertificate)
	// ErrFailedToDialServer       = errors.New(MsgFailedToDialServer)
	// ErrTokenIsInvalid           = errors.New(MsgTokenIsInvalid)

	ErrSettingNotFound = errors.New(MsgSettingNotFound)
	ErrPaymentNotFound = errors.New(MsgPaymentNotFound)
	ErrAuditNotFound   = errors.New(MsgAuditNotFound)

	ErrLicenseNotFound        = errors.New(MsgLicenseNotFound)
	ErrLicenseLinkedToPayment = errors.New(MsgLicenseLinkedToPayment)

	ErrPatternNotFound   = errors.New(MsgPatternNotFound)
	ErrPatternNotDeleted = errors.New(MsgPatternNotDeleted)

	ErrCustomerNotFound   = errors.New(MsgCustomerNotFound)
	ErrCustomerNotDeleted = errors.New(MsgCustomerNotDeleted)
)
