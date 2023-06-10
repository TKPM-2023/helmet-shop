package usermodel

import (
	"TKPM-Go/common"
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

	ErrEmailIsRequired = common.NewCustomError(
		errors.New("email is required"),
		"email is required",
		"ErrEmailIsRequired")

	InvalidEmailFormat = common.NewCustomError(
		errors.New("invalid email format"),
		"invalid email format",
		"InvalidEmailFormat")

	ErrPasswordIsRequired = common.NewCustomError(
		errors.New("password is required"),
		"password is required",
		"ErrPasswordIsRequired")

	InvalidPasswordFormat = common.NewCustomError(
		errors.New("invalid password format"),
		"invalid password format",
		"InvalidPasswordFormat")

	ErrNewPasswordIsRequired = common.NewCustomError(
		errors.New("password new is required"),
		"password new is required",
		"ErrNewPasswordIsRequired")

	InvalidNewPasswordFormat = common.NewCustomError(
		errors.New("invalid new password format"),
		"invalid new password format",
		"InvalidNewPasswordFormat")

	PasswordIncorrect = common.NewCustomError(
		errors.New("password incorrect"),
		"password incorrect",
		"PasswordIncorrect")

	ErrFirstNameIsRequired = common.NewCustomError(
		errors.New("first name is required"),
		"first name is required",
		"ErrFirstNameIsRequired")

	InvalidFirstNameFormat = common.NewCustomError(
		errors.New("invalid first name format"),
		"invalid first name format",
		"InvalidFirstNameFormat")

	ErrLastNameIsRequired = common.NewCustomError(
		errors.New("last name is required"),
		"last name is required",
		"ErrLastNameIsRequired")

	InvalidLastNameFormat = common.NewCustomError(
		errors.New("invalid last name format"),
		"invalid last name format",
		"InvalidLastNameFormat")

	InvalidPhoneFormat = common.NewCustomError(
		errors.New("invalid phone format"),
		"invalid phone format",
		"InvalidPhoneFormat")

	InvalidRoleFormat = common.NewCustomError(
		errors.New("invalid role format"),
		"invalid role format",
		"InvalidRoleFormat")
)
