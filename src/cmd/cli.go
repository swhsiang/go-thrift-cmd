/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"swhsiang/computing"
)

var debugClientProtocol bool

func init() {
	flag.BoolVar(&debugClientProtocol, "debug_client_protocol", false, "turn client protocol trace on")
}

func handleClient(client *computing.ComputingClient) (err error) {
	_, _ = client.Ping()
	fmt.Println("ping()!")
	return nil
}

// StartClient run client
func StartClient(
	host string,
	port int64,
	transport string,
	protocol string) (client *computing.ComputingClient, err error) {

	hostPort := fmt.Sprintf("%s:%d", host, port)

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	if debugClientProtocol {
		protocolFactory = thrift.NewTDebugProtocolFactory(protocolFactory, "client:")
	}
	var trans thrift.TTransport

	trans, err = thrift.NewTSocket(hostPort)

	if err != nil {
		return nil, err
	}
	trans = thrift.NewTBufferedTransport(trans, 8192)

	if err = trans.Open(); err != nil {
		return nil, err
	}
	client = computing.NewComputingClientFactory(trans, protocolFactory)
	return
}

func main() {
	client, err := StartClient("localhost", 9090, "", "")
	fmt.Println("[client]Ping()")
	_, _ = client.Ping()
	if err != nil {
		_ = fmt.Errorf("Unable to start client %v", err)
	}
	defer func ()  {
		_ = client.Transport.Close()
	}()

}
