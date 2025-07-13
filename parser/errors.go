package parser

import "errors"

var (
	ErrTextTooShort    = errors.New("text is too short")
	ErrUnsupportedLang = errors.New("unsupported language")
)
