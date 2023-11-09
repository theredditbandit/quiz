package customerrors

import "errors"

var InvalidSchemaError = errors.New("invalidSchema")
var InvalidFileTypeError = errors.New("invalidFileType")
