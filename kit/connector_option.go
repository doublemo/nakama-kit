// Copyright 2024 The Bombus Authors
//
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package kit

import "github.com/doublemo/nakama-kit/pb"

type (
	ConnectorWriteOption func(msg *pb.ResponseWriter)
)

func WithConnectorWriteNoCache() func(msg *pb.ResponseWriter) {
	return func(msg *pb.ResponseWriter) {
		if msg.Context == nil {
			msg.Context = make(map[string]string)
		}
		msg.Context["Cache-Control"] = "no-cache"
	}
}
