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

package echo

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Echo_Echo_Args struct {
	Ping *Ping `json:"ping,omitempty"`
}

func (v *Echo_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Ping != nil {
		w, err = v.Ping.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Ping_Read(w wire.Value) (*Ping, error) {
	var v Ping
	err := v.FromWire(w)
	return &v, err
}

func (v *Echo_Echo_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Ping, err = _Ping_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Echo_Echo_Args) String() string {
	var fields [1]string
	i := 0
	if v.Ping != nil {
		fields[i] = fmt.Sprintf("Ping: %v", v.Ping)
		i++
	}
	return fmt.Sprintf("Echo_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_Echo_Args) MethodName() string {
	return "echo"
}

func (v *Echo_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Echo_Echo_Helper = struct {
	Args           func(ping *Ping) *Echo_Echo_Args
	IsException    func(error) bool
	WrapResponse   func(*Pong, error) (*Echo_Echo_Result, error)
	UnwrapResponse func(*Echo_Echo_Result) (*Pong, error)
}{}

func init() {
	Echo_Echo_Helper.Args = func(ping *Ping) *Echo_Echo_Args {
		return &Echo_Echo_Args{Ping: ping}
	}
	Echo_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Echo_Echo_Helper.WrapResponse = func(success *Pong, err error) (*Echo_Echo_Result, error) {
		if err == nil {
			return &Echo_Echo_Result{Success: success}, nil
		}
		return nil, err
	}
	Echo_Echo_Helper.UnwrapResponse = func(result *Echo_Echo_Result) (success *Pong, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Echo_Echo_Result struct {
	Success *Pong `json:"success,omitempty"`
}

func (v *Echo_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Pong_Read(w wire.Value) (*Pong, error) {
	var v Pong
	err := v.FromWire(w)
	return &v, err
}

func (v *Echo_Echo_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Pong_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Echo_Echo_Result) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("Echo_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_Echo_Result) MethodName() string {
	return "echo"
}

func (v *Echo_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}