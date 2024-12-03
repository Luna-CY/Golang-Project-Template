package response

type Code int

const (
	Ok = Code(0) // OK

	// code format: [1-9]{1}    [10-98]{2}    [00-99]{2}    [01-99]{2}
	// position 1: code type prefix: 1 - web; 2-8 other server; 9 - common special error
	// position 2: error category: 10-19 - user request error; 20-29 - system error; 30-39 - business error; 99 is special error, not allow use
	// position 3: error sub-category: 00-99, default is 00
	// position 4: code number

	// example:
	// 1100001: email format error
	// 1200001: system error, database error
	// 1300001: business error, commodity not found

	// common error codes

	InvalidParams       = Code(9100001) // invalid parameters
	NeedAuthorization   = Code(9100002) // need authorization
	ServerInternalError = Code(9999999) // internal server error
)
