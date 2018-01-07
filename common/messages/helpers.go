// Copyright 2018 Saurabh Badhwar. All Rights Reserved.
// The use of this code is governed by MIT License
// which can be found in the LICENSE file.

package messages

import (
	"encoding/hex"
	"io"
	"crypto/md5"
)

// generateChecksum returns the checksum for the provided payload
func generateChecksum(payload string) string {
	h := md5.New()
	io.WriteString(h, payload)
	return hex.EncodeToString(h.Sum(nil))
}