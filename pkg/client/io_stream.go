package client

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/kitex-contrib/kitexcall/pkg/log"
)

type outputFormatter struct{}

func (f *outputFormatter) format(output any) (string, error) {
	if outputStr, ok := output.(string); ok {
		return log.FormatJson(outputStr)
	}

	return "", fmt.Errorf("output of %T is not supported, now only support string", output)
}

// ioStream is responsible for interacting with standard input/output or file input/output.
// it provides generic_client with streaming interface.
type ioStream struct {
	// deal with Recv input from
	decoder *json.Decoder
	err     error

	// deal with Send output to
	out       io.Writer
	formatter *outputFormatter
}

func newIoStream(in io.Reader, out io.Writer) *ioStream {
	return &ioStream{
		decoder:   json.NewDecoder(in),
		out:       out,
		formatter: &outputFormatter{},
	}
}

func (st *ioStream) Recv() (string, error) {
	if st.err != nil {
		return "", st.err
	}

	for {
		var jsonMsg json.RawMessage
		if err := st.decoder.Decode(&jsonMsg); err != nil {
			if err == io.EOF {
				return "", err
			}
			st.err = err
			return "", err
		}
		return string(jsonMsg), nil
	}
}

func (st *ioStream) Send(res any) error {
	resStr, err := st.formatter.format(res)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(st.out, "%s\n", resStr)
	return err
}
