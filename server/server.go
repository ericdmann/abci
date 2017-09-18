// Copyright 2015 Tendermint. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package server is used to start a new ABCI server.

It contains two server implementation:
 * gRPC server
 * socket server

*/

package server

import (
	"fmt"

	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
)

func NewServer(protoAddr, transport string, app types.Application) (cmn.Service, error) {
	var s cmn.Service
	var err error
	switch transport {
	case "socket":
		s = NewSocketServer(protoAddr, app)
	case "grpc":
		s = NewGRPCServer(protoAddr, types.NewGRPCApplication(app))
	default:
		err = fmt.Errorf("Unknown server type %s", transport)
	}
	return s, err
}
                  