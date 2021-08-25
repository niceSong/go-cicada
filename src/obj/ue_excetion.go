package obj

type UeException interface {
	// AuthenticationError @eb5gcErr(code = 201001, message = "认证错误")
	AuthenticationError()

	// UeIdFormatError @eb5gcErr(code = 201002, message = "UeId格式错误")
	UeIdFormatError()
}