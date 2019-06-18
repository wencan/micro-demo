#!/usr/bin/env python
# coding: utf-8

import grpc

import health_pb2
import health_pb2_grpc

def main():
    with grpc.insecure_channel("127.0.0.1:35375") as channel:
        stub = health_pb2_grpc.HealthStub(channel)
        resp = stub.Check(health_pb2.HealthCheckRequest(service="go-service"))
        print(resp.status)


if __name__ == "__main__":
    main()