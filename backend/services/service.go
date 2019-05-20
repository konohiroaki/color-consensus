package services

import (
	"fmt"
	"strconv"
	"strings"
)

type ValidationError struct {
	message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{message}
}

func (e ValidationError) Error() string {
	return e.message
}

type InternalServerError struct {
	message string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message}
}

func (e InternalServerError) Error() string {
	return e.message
}

var util = utilFuncs{}

type utilFuncs struct{}

// ff -> 255
func (utilFuncs) fromHex(hex string) int {
	num, _ := strconv.ParseInt(hex, 16, 64)
	return int(num)
}

// 255 -> ff
func (utilFuncs) toHex(num int) string {
	return fmt.Sprintf("%02x", num)
}

func (utilFuncs) abs(num int) int {
	if (num < 0) {
		return -num;
	}
	return num;
}

// #abc -> #aabbcc
func (u utilFuncs) shortToLowerLongHex(code string) string {
	if (len(code) == 4) {
		code = "#" + code[1:2] + code[1:2] + code[2:3] + code[2:3] + code[3:4] + code[3:4]
	}
	return strings.ToLower(code)
}
