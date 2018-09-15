// package: blog.account
// file: account/service.proto

import * as account_service_pb from "../account/service_pb";
import {grpc} from "grpc-web-client";

type AccountServiceQueryUsers = {
  readonly methodName: string;
  readonly service: typeof AccountService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof account_service_pb.QueryUsersRequest;
  readonly responseType: typeof account_service_pb.QueryUsersResponse;
};

type AccountServiceGetUser = {
  readonly methodName: string;
  readonly service: typeof AccountService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof account_service_pb.GetUserRequest;
  readonly responseType: typeof account_service_pb.GetUserResponse;
};

type AccountServiceCreateUser = {
  readonly methodName: string;
  readonly service: typeof AccountService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof account_service_pb.CreateUserRequest;
  readonly responseType: typeof account_service_pb.CreateUserResponse;
};

type AccountServiceDeleteUser = {
  readonly methodName: string;
  readonly service: typeof AccountService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof account_service_pb.DeleteUserRequest;
  readonly responseType: typeof account_service_pb.DeleteUserResponse;
};

type AccountServiceUpdateUser = {
  readonly methodName: string;
  readonly service: typeof AccountService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof account_service_pb.UpdateUserRequest;
  readonly responseType: typeof account_service_pb.UpdateUserResponse;
};

export class AccountService {
  static readonly serviceName: string;
  static readonly QueryUsers: AccountServiceQueryUsers;
  static readonly GetUser: AccountServiceGetUser;
  static readonly CreateUser: AccountServiceCreateUser;
  static readonly DeleteUser: AccountServiceDeleteUser;
  static readonly UpdateUser: AccountServiceUpdateUser;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class AccountServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  queryUsers(
    requestMessage: account_service_pb.QueryUsersRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: account_service_pb.QueryUsersResponse|null) => void
  ): void;
  queryUsers(
    requestMessage: account_service_pb.QueryUsersRequest,
    callback: (error: ServiceError, responseMessage: account_service_pb.QueryUsersResponse|null) => void
  ): void;
  getUser(
    requestMessage: account_service_pb.GetUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: account_service_pb.GetUserResponse|null) => void
  ): void;
  getUser(
    requestMessage: account_service_pb.GetUserRequest,
    callback: (error: ServiceError, responseMessage: account_service_pb.GetUserResponse|null) => void
  ): void;
  createUser(
    requestMessage: account_service_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: account_service_pb.CreateUserResponse|null) => void
  ): void;
  createUser(
    requestMessage: account_service_pb.CreateUserRequest,
    callback: (error: ServiceError, responseMessage: account_service_pb.CreateUserResponse|null) => void
  ): void;
  deleteUser(
    requestMessage: account_service_pb.DeleteUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: account_service_pb.DeleteUserResponse|null) => void
  ): void;
  deleteUser(
    requestMessage: account_service_pb.DeleteUserRequest,
    callback: (error: ServiceError, responseMessage: account_service_pb.DeleteUserResponse|null) => void
  ): void;
  updateUser(
    requestMessage: account_service_pb.UpdateUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: account_service_pb.UpdateUserResponse|null) => void
  ): void;
  updateUser(
    requestMessage: account_service_pb.UpdateUserRequest,
    callback: (error: ServiceError, responseMessage: account_service_pb.UpdateUserResponse|null) => void
  ): void;
}

