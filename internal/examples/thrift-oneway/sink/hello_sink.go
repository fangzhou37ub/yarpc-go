// Code generated by thriftrw v1.0.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sink

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Hello_Sink_Args struct {
	Snk *SinkRequest `json:"snk,omitempty"`
}

func (v *Hello_Sink_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Snk != nil {
		w, err = v.Snk.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SinkRequest_Read(w wire.Value) (*SinkRequest, error) {
	var v SinkRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *Hello_Sink_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Snk, err = _SinkRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Hello_Sink_Args) String() string {
	var fields [1]string
	i := 0
	if v.Snk != nil {
		fields[i] = fmt.Sprintf("Snk: %v", v.Snk)
		i++
	}
	return fmt.Sprintf("Hello_Sink_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Hello_Sink_Args) MethodName() string {
	return "sink"
}

func (v *Hello_Sink_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

var Hello_Sink_Helper = struct {
	Args func(snk *SinkRequest) *Hello_Sink_Args
}{}

func init() {
	Hello_Sink_Helper.Args = func(snk *SinkRequest) *Hello_Sink_Args {
		return &Hello_Sink_Args{Snk: snk}
	}
}