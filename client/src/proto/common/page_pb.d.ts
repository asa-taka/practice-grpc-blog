// package: common
// file: common/page.proto

import * as jspb from "google-protobuf";

export class PageRequest extends jspb.Message {
  getSize(): number;
  setSize(value: number): void;

  getFrom(): number;
  setFrom(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PageRequest): PageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PageRequest;
  static deserializeBinaryFromReader(message: PageRequest, reader: jspb.BinaryReader): PageRequest;
}

export namespace PageRequest {
  export type AsObject = {
    size: number,
    from: number,
  }
}

export class PageResponse extends jspb.Message {
  getSize(): number;
  setSize(value: number): void;

  getFrom(): number;
  setFrom(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PageResponse): PageResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PageResponse;
  static deserializeBinaryFromReader(message: PageResponse, reader: jspb.BinaryReader): PageResponse;
}

export namespace PageResponse {
  export type AsObject = {
    size: number,
    from: number,
  }
}

