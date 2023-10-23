package kitkit

import "errors"

var ErrorNotPointer = errors.New("init val must be a pointer")
var ErrorUnsupportedType = errors.New("unsupported init val type")
