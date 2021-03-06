// Copyright 2019 Yunion
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

package common

import (
	"fmt"
	"os"

	"yunion.io/x/log"
	"yunion.io/x/pkg/util/reflectutils"
	"yunion.io/x/structarg"
)

func ParseOptions(optStruct interface{}, args []string, configFileName string) {
	parser, err := structarg.NewArgumentParser(optStruct,
		"email-sender", "", "")
	if err != nil {
		log.Fatalf("Error define argument parser: %v", err)
	}

	err = parser.ParseArgs2(args[1:], false, false)
	if err != nil {
		log.Fatalf("Parse arguments error: %v", err)
	}

	var optionsRef *structarg.BaseOptions

	err = reflectutils.FindAnonymouStructPointer(optStruct, &optionsRef)
	if err != nil {
		log.Fatalf("Find common options fail %s", err)
	}

	if optionsRef.Help {
		fmt.Println(parser.HelpString())
		os.Exit(0)
	}

	if len(optionsRef.Config) > 0 {
		log.Infof("Use configuration file: %s", optionsRef.Config)
		err = parser.ParseFile(optionsRef.Config)
		if err != nil {
			log.Fatalf("Parse configuration file: %v", err)
		}
	}
	parser.SetDefault()
}

type IServiceOptions interface {
	GetLogLevel() string
	GetSockFileDir() string
	GetSenderNum() int
	GetOthers() interface{}
}

type SBaseOptions struct {
	SockFileDir string `help:"socket file directory" default:"/etc/yunion/notify"`
	SenderNum   int    `default:"50" help:"number of sender"`
	LogLevel    string `help:"log level" default:"info" choices:"debug|info|warn|error"`

	structarg.BaseOptions
}

func (o SBaseOptions) GetLogLevel() string {
	return o.LogLevel
}

func (o SBaseOptions) GetSockFileDir() string {
	return o.SockFileDir
}

func (o SBaseOptions) GetSenderNum() int {
	return o.SenderNum
}

func (s SBaseOptions) GetOthers() interface{} {
	return nil
}
