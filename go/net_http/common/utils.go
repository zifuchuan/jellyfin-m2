package common

import "io"

func safeClose(closer io.Closer) {
	_ = closer.Close()
}
