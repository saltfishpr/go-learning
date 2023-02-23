/**
 * @fileoverview gRPC-Web generated client stub for saltfishpr.demo.user.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v0.0.0
// source: user/v1/user.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as user_v1_user_pb from '../../user/v1/user_pb';


export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorRegister = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/Register',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.RegisterRequest,
    user_v1_user_pb.RegisterResponse,
    (request: user_v1_user_pb.RegisterRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.RegisterResponse.deserializeBinary
  );

  register(
    request: user_v1_user_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.RegisterResponse>;

  register(
    request: user_v1_user_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.RegisterResponse) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.RegisterResponse>;

  register(
    request: user_v1_user_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.RegisterResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/Register',
        request,
        metadata || {},
        this.methodDescriptorRegister,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/Register',
    request,
    metadata || {},
    this.methodDescriptorRegister);
  }

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/Login',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.LoginRequest,
    user_v1_user_pb.LoginResponse,
    (request: user_v1_user_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.LoginResponse.deserializeBinary
  );

  login(
    request: user_v1_user_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.LoginResponse>;

  login(
    request: user_v1_user_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.LoginResponse) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.LoginResponse>;

  login(
    request: user_v1_user_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.LoginResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorCreateUser = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/CreateUser',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.CreateUserRequest,
    user_v1_user_pb.User,
    (request: user_v1_user_pb.CreateUserRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.User.deserializeBinary
  );

  createUser(
    request: user_v1_user_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.User>;

  createUser(
    request: user_v1_user_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.User>;

  createUser(
    request: user_v1_user_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/CreateUser',
        request,
        metadata || {},
        this.methodDescriptorCreateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/CreateUser',
    request,
    metadata || {},
    this.methodDescriptorCreateUser);
  }

  methodDescriptorUpdateUser = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/UpdateUser',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.UpdateUserRequest,
    user_v1_user_pb.User,
    (request: user_v1_user_pb.UpdateUserRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.User.deserializeBinary
  );

  updateUser(
    request: user_v1_user_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.User>;

  updateUser(
    request: user_v1_user_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.User>;

  updateUser(
    request: user_v1_user_pb.UpdateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/UpdateUser',
        request,
        metadata || {},
        this.methodDescriptorUpdateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/UpdateUser',
    request,
    metadata || {},
    this.methodDescriptorUpdateUser);
  }

  methodDescriptorDeleteUser = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/DeleteUser',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.DeleteUserRequest,
    google_protobuf_empty_pb.Empty,
    (request: user_v1_user_pb.DeleteUserRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteUser(
    request: user_v1_user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteUser(
    request: user_v1_user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteUser(
    request: user_v1_user_pb.DeleteUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/DeleteUser',
        request,
        metadata || {},
        this.methodDescriptorDeleteUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/DeleteUser',
    request,
    metadata || {},
    this.methodDescriptorDeleteUser);
  }

  methodDescriptorGetUser = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/GetUser',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.GetUserRequest,
    user_v1_user_pb.User,
    (request: user_v1_user_pb.GetUserRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.User.deserializeBinary
  );

  getUser(
    request: user_v1_user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.User>;

  getUser(
    request: user_v1_user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.User>;

  getUser(
    request: user_v1_user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/GetUser',
        request,
        metadata || {},
        this.methodDescriptorGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/GetUser',
    request,
    metadata || {},
    this.methodDescriptorGetUser);
  }

  methodDescriptorListUser = new grpcWeb.MethodDescriptor(
    '/saltfishpr.demo.user.v1.UserService/ListUser',
    grpcWeb.MethodType.UNARY,
    user_v1_user_pb.ListUserRequest,
    user_v1_user_pb.ListUserResponse,
    (request: user_v1_user_pb.ListUserRequest) => {
      return request.serializeBinary();
    },
    user_v1_user_pb.ListUserResponse.deserializeBinary
  );

  listUser(
    request: user_v1_user_pb.ListUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user_v1_user_pb.ListUserResponse>;

  listUser(
    request: user_v1_user_pb.ListUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.ListUserResponse) => void): grpcWeb.ClientReadableStream<user_v1_user_pb.ListUserResponse>;

  listUser(
    request: user_v1_user_pb.ListUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user_v1_user_pb.ListUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/saltfishpr.demo.user.v1.UserService/ListUser',
        request,
        metadata || {},
        this.methodDescriptorListUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/saltfishpr.demo.user.v1.UserService/ListUser',
    request,
    metadata || {},
    this.methodDescriptorListUser);
  }

}
