// Copyright 2025 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package streaming

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type Message struct {
	Msg string `thrift:"Msg,1" frugal:"1,default,string" json:"Msg"`
}

func NewMessage() *Message {
	return &Message{}
}

func (p *Message) InitDefault() {
}

func (p *Message) GetMsg() (v string) {
	return p.Msg
}
func (p *Message) SetMsg(val string) {
	p.Msg = val
}

func (p *Message) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Message(%+v)", *p)
}

var fieldIDToName_Message = map[int16]string{
	1: "Msg",
}

type BaseResp struct {
	StatusCode    int32  `thrift:"StatusCode,1" frugal:"1,default,i32" json:"StatusCode"`
	StatusMessage string `thrift:"StatusMessage,2" frugal:"2,default,string" json:"StatusMessage"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{}
}

func (p *BaseResp) InitDefault() {
}

func (p *BaseResp) GetStatusCode() (v int32) {
	return p.StatusCode
}

func (p *BaseResp) GetStatusMessage() (v string) {
	return p.StatusMessage
}
func (p *BaseResp) SetStatusCode(val int32) {
	p.StatusCode = val
}
func (p *BaseResp) SetStatusMessage(val string) {
	p.StatusMessage = val
}

func (p *BaseResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResp(%+v)", *p)
}

var fieldIDToName_BaseResp = map[int16]string{
	1: "StatusCode",
	2: "StatusMessage",
}

type StreamingResponse struct {
	Msg      *Message  `thrift:"Msg,1" frugal:"1,default,Message" json:"Msg"`
	BaseResp *BaseResp `thrift:"BaseResp,2" frugal:"2,default,BaseResp" json:"BaseResp"`
}

func NewStreamingResponse() *StreamingResponse {
	return &StreamingResponse{}
}

func (p *StreamingResponse) InitDefault() {
}

var StreamingResponse_Msg_DEFAULT *Message

func (p *StreamingResponse) GetMsg() (v *Message) {
	if !p.IsSetMsg() {
		return StreamingResponse_Msg_DEFAULT
	}
	return p.Msg
}

var StreamingResponse_BaseResp_DEFAULT *BaseResp

func (p *StreamingResponse) GetBaseResp() (v *BaseResp) {
	if !p.IsSetBaseResp() {
		return StreamingResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *StreamingResponse) SetMsg(val *Message) {
	p.Msg = val
}
func (p *StreamingResponse) SetBaseResp(val *BaseResp) {
	p.BaseResp = val
}

func (p *StreamingResponse) IsSetMsg() bool {
	return p.Msg != nil
}

func (p *StreamingResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *StreamingResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingResponse(%+v)", *p)
}

var fieldIDToName_StreamingResponse = map[int16]string{
	1: "Msg",
	2: "BaseResp",
}

type StreamingService interface {
	BidirectionalStream(stream StreamingService_BidirectionalStreamServer) (err error)

	ServerStream(req *Message, stream StreamingService_ServerStreamServer) (err error)

	ClientStream(stream StreamingService_ClientStreamServer) (err error)
}

type StreamingServiceBidirectionalStreamArgs struct {
	Req *Message `thrift:"req,1" frugal:"1,default,Message" json:"req"`
}

func NewStreamingServiceBidirectionalStreamArgs() *StreamingServiceBidirectionalStreamArgs {
	return &StreamingServiceBidirectionalStreamArgs{}
}

func (p *StreamingServiceBidirectionalStreamArgs) InitDefault() {
}

var StreamingServiceBidirectionalStreamArgs_Req_DEFAULT *Message

func (p *StreamingServiceBidirectionalStreamArgs) GetReq() (v *Message) {
	if !p.IsSetReq() {
		return StreamingServiceBidirectionalStreamArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *StreamingServiceBidirectionalStreamArgs) SetReq(val *Message) {
	p.Req = val
}

func (p *StreamingServiceBidirectionalStreamArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *StreamingServiceBidirectionalStreamArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceBidirectionalStreamArgs(%+v)", *p)
}

var fieldIDToName_StreamingServiceBidirectionalStreamArgs = map[int16]string{
	1: "req",
}

type StreamingServiceBidirectionalStreamResult struct {
	Success *StreamingResponse `thrift:"success,0,optional" frugal:"0,optional,StreamingResponse" json:"success,omitempty"`
}

func NewStreamingServiceBidirectionalStreamResult() *StreamingServiceBidirectionalStreamResult {
	return &StreamingServiceBidirectionalStreamResult{}
}

func (p *StreamingServiceBidirectionalStreamResult) InitDefault() {
}

var StreamingServiceBidirectionalStreamResult_Success_DEFAULT *StreamingResponse

func (p *StreamingServiceBidirectionalStreamResult) GetSuccess() (v *StreamingResponse) {
	if !p.IsSetSuccess() {
		return StreamingServiceBidirectionalStreamResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StreamingServiceBidirectionalStreamResult) SetSuccess(x interface{}) {
	p.Success = x.(*StreamingResponse)
}

func (p *StreamingServiceBidirectionalStreamResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StreamingServiceBidirectionalStreamResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceBidirectionalStreamResult(%+v)", *p)
}

var fieldIDToName_StreamingServiceBidirectionalStreamResult = map[int16]string{
	0: "success",
}

type StreamingService_BidirectionalStreamServer interface {
	streaming.Stream

	Recv() (*Message, error)

	Send(*StreamingResponse) error
}

type StreamingServiceServerStreamArgs struct {
	Req *Message `thrift:"req,1" frugal:"1,default,Message" json:"req"`
}

func NewStreamingServiceServerStreamArgs() *StreamingServiceServerStreamArgs {
	return &StreamingServiceServerStreamArgs{}
}

func (p *StreamingServiceServerStreamArgs) InitDefault() {
}

var StreamingServiceServerStreamArgs_Req_DEFAULT *Message

func (p *StreamingServiceServerStreamArgs) GetReq() (v *Message) {
	if !p.IsSetReq() {
		return StreamingServiceServerStreamArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *StreamingServiceServerStreamArgs) SetReq(val *Message) {
	p.Req = val
}

func (p *StreamingServiceServerStreamArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *StreamingServiceServerStreamArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceServerStreamArgs(%+v)", *p)
}

var fieldIDToName_StreamingServiceServerStreamArgs = map[int16]string{
	1: "req",
}

type StreamingServiceServerStreamResult struct {
	Success *StreamingResponse `thrift:"success,0,optional" frugal:"0,optional,StreamingResponse" json:"success,omitempty"`
}

func NewStreamingServiceServerStreamResult() *StreamingServiceServerStreamResult {
	return &StreamingServiceServerStreamResult{}
}

func (p *StreamingServiceServerStreamResult) InitDefault() {
}

var StreamingServiceServerStreamResult_Success_DEFAULT *StreamingResponse

func (p *StreamingServiceServerStreamResult) GetSuccess() (v *StreamingResponse) {
	if !p.IsSetSuccess() {
		return StreamingServiceServerStreamResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StreamingServiceServerStreamResult) SetSuccess(x interface{}) {
	p.Success = x.(*StreamingResponse)
}

func (p *StreamingServiceServerStreamResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StreamingServiceServerStreamResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceServerStreamResult(%+v)", *p)
}

var fieldIDToName_StreamingServiceServerStreamResult = map[int16]string{
	0: "success",
}

type StreamingService_ServerStreamServer interface {
	streaming.Stream

	Send(*StreamingResponse) error
}

type StreamingServiceClientStreamArgs struct {
	Req *Message `thrift:"req,1" frugal:"1,default,Message" json:"req"`
}

func NewStreamingServiceClientStreamArgs() *StreamingServiceClientStreamArgs {
	return &StreamingServiceClientStreamArgs{}
}

func (p *StreamingServiceClientStreamArgs) InitDefault() {
}

var StreamingServiceClientStreamArgs_Req_DEFAULT *Message

func (p *StreamingServiceClientStreamArgs) GetReq() (v *Message) {
	if !p.IsSetReq() {
		return StreamingServiceClientStreamArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *StreamingServiceClientStreamArgs) SetReq(val *Message) {
	p.Req = val
}

func (p *StreamingServiceClientStreamArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *StreamingServiceClientStreamArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceClientStreamArgs(%+v)", *p)
}

var fieldIDToName_StreamingServiceClientStreamArgs = map[int16]string{
	1: "req",
}

type StreamingServiceClientStreamResult struct {
	Success *StreamingResponse `thrift:"success,0,optional" frugal:"0,optional,StreamingResponse" json:"success,omitempty"`
}

func NewStreamingServiceClientStreamResult() *StreamingServiceClientStreamResult {
	return &StreamingServiceClientStreamResult{}
}

func (p *StreamingServiceClientStreamResult) InitDefault() {
}

var StreamingServiceClientStreamResult_Success_DEFAULT *StreamingResponse

func (p *StreamingServiceClientStreamResult) GetSuccess() (v *StreamingResponse) {
	if !p.IsSetSuccess() {
		return StreamingServiceClientStreamResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StreamingServiceClientStreamResult) SetSuccess(x interface{}) {
	p.Success = x.(*StreamingResponse)
}

func (p *StreamingServiceClientStreamResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StreamingServiceClientStreamResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StreamingServiceClientStreamResult(%+v)", *p)
}

var fieldIDToName_StreamingServiceClientStreamResult = map[int16]string{
	0: "success",
}

type StreamingService_ClientStreamServer interface {
	streaming.Stream

	Recv() (*Message, error)

	SendAndClose(*StreamingResponse) error
}
