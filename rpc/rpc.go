//
// This file is part of arduino-cli.
//
// Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to modify or
// otherwise use the software for commercial activities involving the Arduino
// software without disclosing the source code of your own applications. To purchase
// a commercial license, send an email to license@arduino.cc.
//

package rpc

//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. board.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. commands.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. common.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. compile.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. core.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. lib.proto
//go:generate protoc -I . -I .. --go_out=plugins=grpc:../../../.. upload.proto

// protoc --plugin=protoc-gen-grpc-java=path/to/protoc-gen-grpc-java --grpc-java_out=java --proto_path=. board.proto
// find . -name "*.proto" -exec protoc -I=. --java_out=java {} \;
