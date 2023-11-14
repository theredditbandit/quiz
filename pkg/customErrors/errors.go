package customerrors

import "errors"

var ErrInvalidSchema = errors.New("invalidSchema")
var ErrInvalidFileType = errors.New("invalidFileType")
