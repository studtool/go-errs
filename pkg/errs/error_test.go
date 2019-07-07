package errs

import (
	"testing"

	"github.com/franela/goblin"

	"errors"
	"fmt"
)

func TestNew(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Wrap()", func() {
		g.It("should create internal error", func() {
			err := Wrap(nil)
			g.Assert(err.Type).Equal(Internal)
			g.Assert(err.Code).Equal(InternalErrorCode)
			g.Assert(err.Message).Equal(InternalErrorMessage)
		})
	})
}

const (
	testErrorCode    = 0
	testErrorMessage = "the error message"
)

func TestNewInternal(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, Internal,
		NewInternal, "NewInternal",
	)
}

func TestNewBadFormat(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, BadFormat,
		NewBadFormat, "NewBadFormat",
	)
}

func TestNewInvalidFormat(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, InvalidFormat,
		NewInvalidFormat, "NewInvalidFormat",
	)
}

func TestNewConflict(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, Conflict,
		NewConflict, "NewConflict",
	)
}

func TestNewNotFound(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, NotFound,
		NewNotFound, "NewNotFound",
	)
}

func TestNewNotAuthorized(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, NotAuthorized,
		NewNotAuthorized, "NewNotAuthorized",
	)
}

func TestNewPermissionDenied(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, PermissionDenied,
		NewPermissionDenied, "NewPermissionDenied",
	)
}

func TestNewNotImplemented(t *testing.T) {
	g := goblin.Goblin(t)
	testErrorConstructor(
		g, NotImplemented,
		NewNotImplemented, "NewNotImplemented",
	)
}

type errorConstructor func(code int, message string) *Error

func testErrorConstructor(
	g *goblin.G, t int,
	constructor errorConstructor,
	constructorName string,
) {
	g.Describe(constructorName+"()", func() {
		g.It("should create corresponding error", func() {
			err := constructor(
				testErrorCode, testErrorMessage,
			)

			g.Assert(err.Type).Equal(t)
			g.Assert(err.Code).Equal(testErrorCode)
			g.Assert(err.Message).Equal(testErrorMessage)
		})
	})
}

func TestError_Error(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Error.Error()", func() {
		g.It("should return error description", func() {
			err := NewCustom(
				Internal, InternalErrorCode,
				InternalErrorMessage, errors.New(InternalErrorMessage),
			)

			actual := err.Error()
			expected := fmt.Sprintf(
				"error: [type = '%d'; code: = '%d'; message = '%s'; cause = '%s']",
				Internal, InternalErrorCode, InternalErrorMessage, InternalErrorMessage,
			)

			g.Assert(actual).Equal(expected)
		})
	})
}
