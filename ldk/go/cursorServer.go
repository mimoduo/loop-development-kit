package ldk

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/open-olive/loop-development-kit-go/proto"
)

type CursorServer struct {
	Impl CursorService
}

func (c *CursorServer) CursorPosition(ctx context.Context, empty *empty.Empty) (*proto.CursorPositionResponse, error) {
	value, err := c.Impl.Position()
	if err != nil {
		return nil, err
	}
	return &proto.CursorPositionResponse{
		X:      int32(value.X),
		Y:      int32(value.Y),
		Screen: 0,
	}, nil
}

func (c *CursorServer) CursorPositionStream(_ *empty.Empty, stream proto.Cursor_CursorPositionStreamServer) error {
	handler := func(msg CursorPosition, err error) {
		var errText string
		if err != nil {
			errText = err.Error()
		}
		if e := stream.Send(&proto.CursorPositionStreamResponse{
			X:      int32(msg.X),
			Y:      int32(msg.Y),
			Screen: 0,
			Error:  errText,
		}); e != nil {
			fmt.Println("error: ldk.CursorServer.CursorPositionStream -> stream.Send:", e)
		}
	}

	go func() {
		err := c.Impl.ListenPosition(stream.Context(), handler)
		if err != nil {
			fmt.Println("error: ldk.CursorServer.CursorPositionStream -> Listen:", err)
		}
	}()

	<-stream.Context().Done()
	return nil
}
