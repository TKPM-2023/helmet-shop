package usermodel

import (
	"LearnGo/common"
	"errors"
)

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already exits"),
		"email has already exits",
		"ErrEmailExist")

	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid")

	InvalidEmailFormat = common.NewCustomError(
		errors.New("invalid email format"),
		"invalid email format",
		"InvalidEmailFormat")

	InvalidPasswordFormat = common.NewCustomError(
		errors.New("invalid password format"),
		"invalid password format",
		"InvalidPasswordFormat")

	InvalidFirstNameFormat = common.NewCustomError(
		errors.New("invalid first name format"),
		"invalid first name format",
		"InvalidFirstNameFormat")

	InvalidLastNameFormat = common.NewCustomError(
		errors.New("invalid last name format"),
		"invalid last name format",
		"InvalidLastNameFormat")
)
