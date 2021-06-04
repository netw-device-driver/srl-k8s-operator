/*
	Copyright 2021 Wim Henderickx.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package controllers

import (
	"context"
	"net"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/netw-device-driver/netwdevpb"
	"google.golang.org/grpc"
)

type Deviation struct {
	netwdevpb.UnimplementedDeviationServer

	NewUpdates *bool
	Data       map[string]string
	Mutex      sync.RWMutex
}

// DeviationServer contains the device driver information
type DeviationServer struct {
	DeviationServerPort    *int
	DeviationServerAddress *string
	Deviation              *Deviation
	Debug                  *bool

	StopCh chan struct{}
	Ctx    context.Context
}

// Option is a function to initialize the options of the srl operator deviation server
type Option func(d *DeviationServer)

// WithDeviationServer initializes the deviation server in the srl operator
func WithDeviationServer(s *string) Option {
	return func(d *DeviationServer) {
		d.DeviationServerAddress = s
		p, _ := strconv.Atoi(strings.Split(*s, ":")[1])
		d.DeviationServerPort = &p
	}
}

func WithDebug(b *bool) Option {
	return func(d *DeviationServer) {
		d.Debug = b
	}
}

// NewDeviceDriver function defines a new device driver
func NewDeviationServer(opts ...Option) *DeviationServer {
	log.Info("initialize new SRL deviation server ...")

	d := &DeviationServer{
		DeviationServerAddress: new(string),
		DeviationServerPort:    new(int),
		Deviation: &Deviation{
			NewUpdates: new(bool),
			Data:       make(map[string]string),
		},
		StopCh: make(chan struct{}),
		Debug:  new(bool),
	}
	return d
}

// StartDeviationGRPCServer function starts the deviation server
func (d *DeviationServer) StartDeviationGRPCServer() {
	log.Info("Starting deviation GRPC server...")

	// create a listener on a specific address:port
	lis, err := net.Listen("tcp", *d.DeviationServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the gRPC service to the server
	netwdevpb.RegisterDeviationServer(grpcServer, d.Deviation)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// Update is a GRPC service that serves the deviation
func (d *Deviation) Update(ctx context.Context, req *netwdevpb.DeviationUpdate) (*netwdevpb.DeviationUpdateReply, error) {
	log.Infof("Deviation Update: Object: %v", req.Resource)
	reply := &netwdevpb.DeviationUpdateReply{}

	return reply, nil
}
