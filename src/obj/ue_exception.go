package obj

type UeException interface {
	// AuthenticationError @CicadaError(code = 201001, message = "认证错误")
	AuthenticationError()

	// UeIdFormatError @CicadaError(code = 201002, message = "UeId格式错误")
	UeIdFormatError()
}