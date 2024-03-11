package errcodes

type Code string

var codesMessages = map[Code]string{
	CodeIngredientNotFound:      codeIngredientNotFoundMessage,
	CodeIngredientInvalid:       codeIngredientInvalidMessage,
	CodeIngredientAlreadyExists: codeIngredientAlreadyExistsMessage,
	CodeUnknown:                 codeUnknownMessage,
	CodeUnauthorized:            codeUnauthorizedMessage,
	CodePermissionDenied:        codePermissionDeniedMessage,
}

// ingredient service statuses
const (
	// CodeIngredientInvalid - request validation error
	CodeIngredientInvalid Code = "IG22"
	// CodeIngredientNotFound - where is no Ingredient with this ID found
	CodeIngredientNotFound Code = "IG04"
	// CodeIngredientAlreadyExists - where is already same ingredient in it.
	CodeIngredientAlreadyExists Code = "IG09"
)

const (
	codeIngredientInvalidMessage       = "ingredient is not valid"
	codeIngredientNotFoundMessage      = "no such ingredient"
	codeIngredientAlreadyExistsMessage = "ingredient already exists"
	codeUnknownMessage                 = "unknown status"
	codeUnauthorizedMessage            = "unauthorized"
	codePermissionDeniedMessage        = "permission denied"
)

const (
	CodeUnknown          Code = "IT00"
	CodeUnauthorized     Code = "IT03"
	CodePermissionDenied Code = "IT01"
)
