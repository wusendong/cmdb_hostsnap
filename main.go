/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wusendong/cmdb_hostsnap/command"
	"gopkg.in/urfave/cli.v1"
)

// app info
const (
	AppName = "hostsnap"
	Usage   = "hostsnap"
)

// build info
var (
	Version     = "0.1.0"
	BuildCommit = ""
	BuildTime   = ""
	GoVersion   = ""
)

func cmdNotFound(c *cli.Context, command string) {
	panic(fmt.Errorf("Unrecognized command: %s", command))
}

func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	panic(fmt.Errorf("Usage error, please check your command"))
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	a := cli.NewApp()
	a.Version = Version
	a.Name = AppName
	a.Usage = Usage

	a.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, d",
			Usage:  "enable debug logging level",
			EnvVar: "CMDB_DEBUG",
		},
	}
	a.Commands = []cli.Command{
		command.DaemonCmd(),
		command.ReloadCmd(),
	}
	a.CommandNotFound = cmdNotFound
	a.OnUsageError = onUsageError

	if err := a.Run(os.Args); err != nil {
		message := fmt.Sprintf("Critical error: %v", err)
		// log to std error
		fmt.Fprintln(os.Stderr, message)
		// log to log file
		logrus.Fatal(message)
	}
}