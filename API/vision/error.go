package vision

import (
	"fmt"

	"github.com/pkg/errors"
)

func wrapError(err error, op string, msg string) error {
	return errors.Wrap(err, fmt.Sprintf("%s: %s", op, msg))
}
