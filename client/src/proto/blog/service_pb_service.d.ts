// package: blog
// file: blog/service.proto

import * as blog_service_pb from "../blog/service_pb";
import {grpc} from "grpc-web-client";

type BlogServiceQueryDocument = {
  readonly methodName: string;
  readonly service: typeof BlogService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_service_pb.QueryDocumentsRequest;
  readonly responseType: typeof blog_service_pb.QueryDocumentsResponse;
};

type BlogServiceGetDocument = {
  readonly methodName: string;
  readonly service: typeof BlogService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_service_pb.GetDocumentRequest;
  readonly responseType: typeof blog_service_pb.GetDocumentResponse;
};

type BlogServiceCreateDocument = {
  readonly methodName: string;
  readonly service: typeof BlogService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_service_pb.CreateDocumentRequest;
  readonly responseType: typeof blog_service_pb.CreateDocumentResponse;
};

type BlogServiceDeleteDocument = {
  readonly methodName: string;
  readonly service: typeof BlogService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_service_pb.DeleteDocumentRequest;
  readonly responseType: typeof blog_service_pb.DeleteDocumentResponse;
};

type BlogServiceUpdateDocument = {
  readonly methodName: string;
  readonly service: typeof BlogService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof blog_service_pb.UpdateDocumentRequest;
  readonly responseType: typeof blog_service_pb.UpdateDocumentResponse;
};

export class BlogService {
  static readonly serviceName: string;
  static readonly QueryDocument: BlogServiceQueryDocument;
  static readonly GetDocument: BlogServiceGetDocument;
  static readonly CreateDocument: BlogServiceCreateDocument;
  static readonly DeleteDocument: BlogServiceDeleteDocument;
  static readonly UpdateDocument: BlogServiceUpdateDocument;
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

export class BlogServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  queryDocument(
    requestMessage: blog_service_pb.QueryDocumentsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: blog_service_pb.QueryDocumentsResponse|null) => void
  ): void;
  queryDocument(
    requestMessage: blog_service_pb.QueryDocumentsRequest,
    callback: (error: ServiceError, responseMessage: blog_service_pb.QueryDocumentsResponse|null) => void
  ): void;
  getDocument(
    requestMessage: blog_service_pb.GetDocumentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: blog_service_pb.GetDocumentResponse|null) => void
  ): void;
  getDocument(
    requestMessage: blog_service_pb.GetDocumentRequest,
    callback: (error: ServiceError, responseMessage: blog_service_pb.GetDocumentResponse|null) => void
  ): void;
  createDocument(
    requestMessage: blog_service_pb.CreateDocumentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: blog_service_pb.CreateDocumentResponse|null) => void
  ): void;
  createDocument(
    requestMessage: blog_service_pb.CreateDocumentRequest,
    callback: (error: ServiceError, responseMessage: blog_service_pb.CreateDocumentResponse|null) => void
  ): void;
  deleteDocument(
    requestMessage: blog_service_pb.DeleteDocumentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: blog_service_pb.DeleteDocumentResponse|null) => void
  ): void;
  deleteDocument(
    requestMessage: blog_service_pb.DeleteDocumentRequest,
    callback: (error: ServiceError, responseMessage: blog_service_pb.DeleteDocumentResponse|null) => void
  ): void;
  updateDocument(
    requestMessage: blog_service_pb.UpdateDocumentRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: blog_service_pb.UpdateDocumentResponse|null) => void
  ): void;
  updateDocument(
    requestMessage: blog_service_pb.UpdateDocumentRequest,
    callback: (error: ServiceError, responseMessage: blog_service_pb.UpdateDocumentResponse|null) => void
  ): void;
}

