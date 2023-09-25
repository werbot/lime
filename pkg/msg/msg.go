package msg

import "errors"

const (
	// MsgKeyNotFound is a ...
	MsgKeyNotFound = "key not found"

	// MsgLicenseNotFound is a ...
	MsgLicenseNotFound = "license not found"

	// MsgTariffNotFound is a ...
	MsgTariffNotFound = "tariff not found"

	// MsgCustomerNotFound is a ...
	MsgCustomerNotFound = "customer not found"

	// MsgFailedToParseKeyPair is ...
	MsgFailedToParseKeyPair = "failed to parse key pair"

	// MsgFailedToParseCertificate is ...
	MsgFailedToParseCertificate = "failed to parse certificate"

	// MsgFailedToDialServer is ...
	MsgFailedToDialServer = "failed to dial server"

	// MsgTokenIsInvalid is ...
	MsgTokenIsInvalid = "token is invalid"

	// MsgMissingMetadata is ...
	MsgMissingMetadata = "missing metadata"

	// MsgBadRequest is bad request
	MsgBadRequest = "bad request"

	// MsgNotFound is not Found
	MsgNotFound = "not found"
)

var (
	// ErrKeyNotFound is a ...
	ErrKeyNotFound = errors.New(MsgKeyNotFound)

	// ErrLicenseNotFound is a ...
	ErrLicenseNotFound = errors.New(MsgLicenseNotFound)

	// ErrTariffNotFound is a ...
	ErrTariffNotFound = errors.New(MsgTariffNotFound)

	// ErrCustomerNotFound is a ...
	ErrCustomerNotFound = errors.New(MsgCustomerNotFound)
)
