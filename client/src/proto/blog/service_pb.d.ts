// package: blog
// file: blog/service.proto

import * as jspb from "google-protobuf";
import * as blog_data_pb from "../blog/data_pb";
import * as common_page_pb from "../common/page_pb";

export class QueryDocumentsRequest extends jspb.Message {
  hasPage(): boolean;
  clearPage(): void;
  getPage(): common_page_pb.PageRequest | undefined;
  setPage(value?: common_page_pb.PageRequest): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryDocumentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: QueryDocumentsRequest): QueryDocumentsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryDocumentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryDocumentsRequest;
  static deserializeBinaryFromReader(message: QueryDocumentsRequest, reader: jspb.BinaryReader): QueryDocumentsRequest;
}

export namespace QueryDocumentsRequest {
  export type AsObject = {
    page?: common_page_pb.PageRequest.AsObject,
  }
}

export class QueryDocumentsResponse extends jspb.Message {
  clearDocumentsList(): void;
  getDocumentsList(): Array<blog_data_pb.Document>;
  setDocumentsList(value: Array<blog_data_pb.Document>): void;
  addDocuments(value?: blog_data_pb.Document, index?: number): blog_data_pb.Document;

  hasPage(): boolean;
  clearPage(): void;
  getPage(): common_page_pb.PageResponse | undefined;
  setPage(value?: common_page_pb.PageResponse): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryDocumentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: QueryDocumentsResponse): QueryDocumentsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryDocumentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryDocumentsResponse;
  static deserializeBinaryFromReader(message: QueryDocumentsResponse, reader: jspb.BinaryReader): QueryDocumentsResponse;
}

export namespace QueryDocumentsResponse {
  export type AsObject = {
    documentsList: Array<blog_data_pb.Document.AsObject>,
    page?: common_page_pb.PageResponse.AsObject,
  }
}

export class GetDocumentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentRequest): GetDocumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentRequest;
  static deserializeBinaryFromReader(message: GetDocumentRequest, reader: jspb.BinaryReader): GetDocumentRequest;
}

export namespace GetDocumentRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetDocumentResponse extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDocumentResponse): GetDocumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDocumentResponse;
  static deserializeBinaryFromReader(message: GetDocumentResponse, reader: jspb.BinaryReader): GetDocumentResponse;
}

export namespace GetDocumentResponse {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

export class CreateDocumentRequest extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentRequest): CreateDocumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentRequest;
  static deserializeBinaryFromReader(message: CreateDocumentRequest, reader: jspb.BinaryReader): CreateDocumentRequest;
}

export namespace CreateDocumentRequest {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

export class CreateDocumentResponse extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDocumentResponse): CreateDocumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDocumentResponse;
  static deserializeBinaryFromReader(message: CreateDocumentResponse, reader: jspb.BinaryReader): CreateDocumentResponse;
}

export namespace CreateDocumentResponse {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

export class UpdateDocumentRequest extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentRequest): UpdateDocumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentRequest;
  static deserializeBinaryFromReader(message: UpdateDocumentRequest, reader: jspb.BinaryReader): UpdateDocumentRequest;
}

export namespace UpdateDocumentRequest {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

export class UpdateDocumentResponse extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateDocumentResponse): UpdateDocumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateDocumentResponse;
  static deserializeBinaryFromReader(message: UpdateDocumentResponse, reader: jspb.BinaryReader): UpdateDocumentResponse;
}

export namespace UpdateDocumentResponse {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

export class DeleteDocumentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentRequest): DeleteDocumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteDocumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentRequest;
  static deserializeBinaryFromReader(message: DeleteDocumentRequest, reader: jspb.BinaryReader): DeleteDocumentRequest;
}

export namespace DeleteDocumentRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteDocumentResponse extends jspb.Message {
  hasDocument(): boolean;
  clearDocument(): void;
  getDocument(): blog_data_pb.Document | undefined;
  setDocument(value?: blog_data_pb.Document): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDocumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDocumentResponse): DeleteDocumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteDocumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDocumentResponse;
  static deserializeBinaryFromReader(message: DeleteDocumentResponse, reader: jspb.BinaryReader): DeleteDocumentResponse;
}

export namespace DeleteDocumentResponse {
  export type AsObject = {
    document?: blog_data_pb.Document.AsObject,
  }
}

