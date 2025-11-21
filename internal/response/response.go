package response

import "io"

type Resposne interface {
	Marshal(r io.Writer) error
}
