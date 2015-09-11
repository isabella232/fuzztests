// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/fuzztests
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package fuzztests

import (
	"bytes"
	"fmt"
	gogopop "github.com/gogo/fuzztests/gengogofuzztests/gogopop"
	gofast "github.com/gogo/fuzztests/gofast"
	gogo "github.com/gogo/fuzztests/gogo"
	gogofast "github.com/gogo/fuzztests/gogofast"
	golang "github.com/gogo/fuzztests/golang"
	gogoproto "github.com/gogo/protobuf/proto"
	goproto "github.com/golang/protobuf/proto"
)

func debug(s string, i int, data []byte, err error) {
	gostringer := gogopop.NewFuncs[i]()
	if err := gogoproto.Unmarshal(data, gostringer); err == nil {
		s += ":" + fmt.Sprintf("%#v", gostringer)
	}
	err = fmt.Errorf("[%d](%T):%s:%v", i, gostringer, s, err)
	panic(err)
}

func Fuzz(data []byte) int {
	score := 0
	for i, golangf := range golang.NewFuncs {
		golangpb := golangf()
		if err := goproto.Unmarshal(data, golangpb); err != nil {
			continue
		}
		data2, err := goproto.Marshal(golangpb)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(data, data2) {
			panic("golang proto is not idempotent")
		}
		score = 1
		gofastpb := gofast.NewFuncs[i]()
		if err := goproto.Unmarshal(data, gofastpb); err != nil {
			debug("gofastpb unmarshal", i, data, err)
		}
		if !goproto.Equal(golangpb, gofastpb) {
			panic("golangpb != gofastpb")
		}
		gogopb := gogo.NewFuncs[i]()
		if err := gogoproto.Unmarshal(data, gogopb); err != nil {
			debug("gogopb unmarshal", i, data, err)
		}
		gogofastpb := gogofast.NewFuncs[i]()
		if err := gogoproto.Unmarshal(data, gogofastpb); err != nil {
			debug("gogofastpb unmarshal", i, data, err)
		}
		if !gogoproto.Equal(gogopb, gogofastpb) {
			panic("gogopb != gogofastpb")
		}
	}
	return score
}
