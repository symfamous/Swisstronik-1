package rpc

import (
	"encoding/json"
	"errors"
)

var ErrNotImplemented = errors.New("not implemented")

const (
	InvalidJSON    = -32700 // Invalid JSON was received by the server.
	InvalidRequest = -32600 // The JSON sent is not a valid Request object.
	MethodNotFound = -32601 // The method does not exist / is not available.
	InvalidParams  = -32602 // Invalid method parameter(s).
	InternalError  = -32603 // Internal JSON-RPC error.
)

// Err returns an RPCError based on the given code and data.
//
// Parameters:
// - code: an integer representing the error code.
// - data: any data associated with the error.
// Returns
// - *RPCError: a pointer to an RPCError object.
func Err(code int, data any) *RPCError {
	switch code {
	case InvalidJSON:
		return &RPCError{code: InvalidJSON, message: "Parse error", data: data}
	case InvalidRequest:
		return &RPCError{code: InvalidRequest, message: "Invalid Request", data: data}
	case MethodNotFound:
		return &RPCError{code: MethodNotFound, message: "Method Not Found", data: data}
	case InvalidParams:
		return &RPCError{code: InvalidParams, message: "Invalid Params", data: data}
	default:
		return &RPCError{code: InternalError, message: "Internal Error", data: data}
	}
}

// tryUnwrapToRPCErr unwraps the error and checks if it matches any of the given RPC errors.
// If a match is found, the corresponding RPC error is returned.
// If no match is found, the function returns an InternalError with the original error.
//
// Parameters:
// - err: The error to be unwrapped
// - rpcErrors: variadic list of *RPCError objects to be checked
// Returns:
// - error: the original error
func tryUnwrapToRPCErr(err error, rpcErrors ...*RPCError) error {
	var nodeErr *RPCError
	if json.Unmarshal([]byte(err.Error()), nodeErr) != nil {
		return err
	}

	dataErr := isErrorWithData(nodeErr)
	if dataErr != nil {
		return dataErr
	}

	for _, rpcErr := range rpcErrors {
		if errors.Is(nodeErr, rpcErr) {
			return rpcErr
		}
	}
	return Err(InternalError, err)
}

// isErrorWithData checks if the error is the type of error that might contain information in the data field.
// In the case it is, it adds this information to the returned error.
//
// Parameters:
// - nodeErr: The error to be checked
// Returns:
// - *RPCError: a pointer to the RPCError resulting object
func isErrorWithData(nodeErr *RPCError) *RPCError {
	switch nodeErr.code {
	case ErrUnexpectedError.code:
		unexpectedErr := *ErrUnexpectedError
		unexpectedErr.data = nodeErr.data
		return &unexpectedErr
	case ErrNoTraceAvailable.code:
		noTraceAvailableError := *ErrNoTraceAvailable
		noTraceAvailableError.data = nodeErr.data
		return &noTraceAvailableError
	case ErrContractError.code:
		contractError := *ErrContractError
		contractError.data = nodeErr.data
		return &contractError
	}
	return nil
}

type RPCError struct {
	code    int
	message string
	data    any
}

// Error returns the error message associated with the RPCError.
//
// Parameters:
// - none
// Returns:
// - string: the error message
func (e *RPCError) Error() string {
	return e.message
}

// Code returns the code of the RPCError.
//
// Parameters:
//
//	none
//
// Returns:
//   - int: the code
func (e *RPCError) Code() int {
	return e.code
}

// Data returns the data associated with the RPCError.
//
// Parameters:
//
//	none
//
// Returns:
//   - any: the data
func (e *RPCError) Data() any {
	return e.data
}

var (
	ErrFailedToReceiveTxn = &RPCError{
		code:    1,
		message: "Failed to write transaction",
	}
	ErrNoTraceAvailable = &RPCError{
		code:    10,
		message: "No trace available for transaction",
	}
	ErrContractNotFound = &RPCError{
		code:    20,
		message: "Contract not found",
	}
	ErrBlockNotFound = &RPCError{
		code:    24,
		message: "Block not found",
	}
	ErrInvalidTxnHash = &RPCError{
		code:    25,
		message: "Invalid transaction hash",
	}
	ErrInvalidBlockHash = &RPCError{
		code:    26,
		message: "Invalid block hash",
	}
	ErrInvalidTxnIndex = &RPCError{
		code:    27,
		message: "Invalid transaction index in a block",
	}
	ErrClassHashNotFound = &RPCError{
		code:    28,
		message: "Class hash not found",
	}
	ErrHashNotFound = &RPCError{
		code:    29,
		message: "Transaction hash not found",
	}
	ErrPageSizeTooBig = &RPCError{
		code:    31,
		message: "Requested page size is too big",
	}
	ErrNoBlocks = &RPCError{
		code:    32,
		message: "There are no blocks",
	}
	ErrInvalidContinuationToken = &RPCError{
		code:    33,
		message: "The supplied continuation token is invalid or unknown",
	}
	ErrTooManyKeysInFilter = &RPCError{
		code:    34,
		message: "Too many keys provided in a filter",
	}
	ErrContractError = &RPCError{
		code:    40,
		message: "Contract error",
	}
	ErrTxnExec = &RPCError{
		code:    41,
		message: "Transaction execution error",
	}
	ErrInvalidContractClass = &RPCError{
		code:    50,
		message: "Invalid contract class",
	}
	ErrClassAlreadyDeclared = &RPCError{
		code:    51,
		message: "Class already declared",
	}
	ErrInvalidTransactionNonce = &RPCError{
		code:    52,
		message: "Invalid transaction nonce",
	}
	ErrInsufficientMaxFee = &RPCError{
		code:    53,
		message: "Max fee is smaller than the minimal transaction cost (validation plus fee transfer)",
	}
	ErrInsufficientAccountBalance = &RPCError{
		code:    54,
		message: "Account balance is smaller than the transaction's max_fee",
	}
	ErrValidationFailure = &RPCError{
		code:    55,
		message: "Account validation failed",
	}
	ErrCompilationFailed = &RPCError{
		code:    56,
		message: "Compilation failed",
	}
	ErrContractClassSizeTooLarge = &RPCError{
		code:    57,
		message: "Contract class size is too large",
	}
	ErrNonAccount = &RPCError{
		code:    58,
		message: "Sender address is not an account contract",
	}
	ErrDuplicateTx = &RPCError{
		code:    59,
		message: "A transaction with the same hash already exists in the mempool",
	}
	ErrCompiledClassHashMismatch = &RPCError{
		code:    60,
		message: "The compiled class hash did not match the one supplied in the transaction",
	}
	ErrUnsupportedTxVersion = &RPCError{
		code:    61,
		message: "The transaction version is not supported",
	}
	ErrUnsupportedContractClassVersion = &RPCError{
		code:    62,
		message: "The contract class version is not supported",
	}
	ErrUnexpectedError = &RPCError{
		code:    63,
		message: "An unexpected error occurred",
	}
)
