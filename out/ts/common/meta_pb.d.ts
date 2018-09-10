// package: common
// file: common/meta.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Meta extends jspb.Message {
  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): MutationLog | undefined;
  setCreated(value?: MutationLog): void;

  hasUpdated(): boolean;
  clearUpdated(): void;
  getUpdated(): MutationLog | undefined;
  setUpdated(value?: MutationLog): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Meta.AsObject;
  static toObject(includeInstance: boolean, msg: Meta): Meta.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Meta, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Meta;
  static deserializeBinaryFromReader(message: Meta, reader: jspb.BinaryReader): Meta;
}

export namespace Meta {
  export type AsObject = {
    created?: MutationLog.AsObject,
    updated?: MutationLog.AsObject,
  }
}

export class MutationLog extends jspb.Message {
  hasDate(): boolean;
  clearDate(): void;
  getDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getUserId(): number;
  setUserId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MutationLog.AsObject;
  static toObject(includeInstance: boolean, msg: MutationLog): MutationLog.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MutationLog, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MutationLog;
  static deserializeBinaryFromReader(message: MutationLog, reader: jspb.BinaryReader): MutationLog;
}

export namespace MutationLog {
  export type AsObject = {
    date?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    userId: number,
  }
}

