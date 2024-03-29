// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package post

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Srv interface {
	CreatePost(ctx context.Context, req *CreatePostRequest) (r *CreatePostResponse, err error)
}

type SrvClient struct {
	c thrift.TClient
}

func NewSrvClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SrvClient {
	return &SrvClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSrvClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SrvClient {
	return &SrvClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSrvClient(c thrift.TClient) *SrvClient {
	return &SrvClient{
		c: c,
	}
}

func (p *SrvClient) Client_() thrift.TClient {
	return p.c
}

func (p *SrvClient) CreatePost(ctx context.Context, req *CreatePostRequest) (r *CreatePostResponse, err error) {
	var _args SrvCreatePostArgs
	_args.Req = req
	var _result SrvCreatePostResult
	if err = p.Client_().Call(ctx, "CreatePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type SrvProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Srv
}

func (p *SrvProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SrvProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SrvProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSrvProcessor(handler Srv) *SrvProcessor {
	self := &SrvProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("CreatePost", &srvProcessorCreatePost{handler: handler})
	return self
}
func (p *SrvProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type srvProcessorCreatePost struct {
	handler Srv
}

func (p *srvProcessorCreatePost) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SrvCreatePostArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CreatePost", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := SrvCreatePostResult{}
	var retval *CreatePostResponse
	if retval, err2 = p.handler.CreatePost(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CreatePost: "+err2.Error())
		oprot.WriteMessageBegin("CreatePost", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CreatePost", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type SrvCreatePostArgs struct {
	Req *CreatePostRequest `thrift:"req,1" frugal:"1,default,CreatePostRequest" json:"req"`
}

func NewSrvCreatePostArgs() *SrvCreatePostArgs {
	return &SrvCreatePostArgs{}
}

func (p *SrvCreatePostArgs) InitDefault() {
	*p = SrvCreatePostArgs{}
}

var SrvCreatePostArgs_Req_DEFAULT *CreatePostRequest

func (p *SrvCreatePostArgs) GetReq() (v *CreatePostRequest) {
	if !p.IsSetReq() {
		return SrvCreatePostArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *SrvCreatePostArgs) SetReq(val *CreatePostRequest) {
	p.Req = val
}

var fieldIDToName_SrvCreatePostArgs = map[int16]string{
	1: "req",
}

func (p *SrvCreatePostArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SrvCreatePostArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SrvCreatePostArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SrvCreatePostArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewCreatePostRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SrvCreatePostArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreatePost_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SrvCreatePostArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SrvCreatePostArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SrvCreatePostArgs(%+v)", *p)

}

func (p *SrvCreatePostArgs) DeepEqual(ano *SrvCreatePostArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *SrvCreatePostArgs) Field1DeepEqual(src *CreatePostRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type SrvCreatePostResult struct {
	Success *CreatePostResponse `thrift:"success,0,optional" frugal:"0,optional,CreatePostResponse" json:"success,omitempty"`
}

func NewSrvCreatePostResult() *SrvCreatePostResult {
	return &SrvCreatePostResult{}
}

func (p *SrvCreatePostResult) InitDefault() {
	*p = SrvCreatePostResult{}
}

var SrvCreatePostResult_Success_DEFAULT *CreatePostResponse

func (p *SrvCreatePostResult) GetSuccess() (v *CreatePostResponse) {
	if !p.IsSetSuccess() {
		return SrvCreatePostResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SrvCreatePostResult) SetSuccess(x interface{}) {
	p.Success = x.(*CreatePostResponse)
}

var fieldIDToName_SrvCreatePostResult = map[int16]string{
	0: "success",
}

func (p *SrvCreatePostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SrvCreatePostResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SrvCreatePostResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SrvCreatePostResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewCreatePostResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SrvCreatePostResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreatePost_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SrvCreatePostResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *SrvCreatePostResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SrvCreatePostResult(%+v)", *p)

}

func (p *SrvCreatePostResult) DeepEqual(ano *SrvCreatePostResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *SrvCreatePostResult) Field0DeepEqual(src *CreatePostResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
