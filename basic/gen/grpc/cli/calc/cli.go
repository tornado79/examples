// Code generated by goa v3.1.3, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package cli

import (
	"flag"
	"fmt"
	"os"

	calcc "goa.design/examples/basic/gen/grpc/calc/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `calc add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc add --message '{
      "a": 8399553735696626949,
      "b": 360622074634248926
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddMessageFlag = calcAddFlags.String("message", "", "")
	)
	calcFlags.Usage = calcUsage
	calcAddFlags.Usage = calcAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "calc":
			svcf = calcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "calc":
			switch epn {
			case "add":
				epf = calcAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcc.NewClient(cc, opts...)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers
Usage:
    %s [globalflags] calc COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc add -message JSON

Add implements add.
    -message JSON: 

Example:
    `+os.Args[0]+` calc add --message '{
      "a": 8399553735696626949,
      "b": 360622074634248926
   }'
`, os.Args[0])
}
