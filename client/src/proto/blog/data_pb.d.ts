// package: blog
// file: blog/data.proto

import * as jspb from "google-protobuf";
import * as common_meta_pb from "../common/meta_pb";

export class Document extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  hasMeta(): boolean;
  clearMeta(): void;
  getMeta(): common_meta_pb.Meta | undefined;
  setMeta(value?: common_meta_pb.Meta): void;

  getTitle(): string;
  setTitle(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  clearTagsList(): void;
  getTagsList(): Array<string>;
  setTagsList(value: Array<string>): void;
  addTags(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Document.AsObject;
  static toObject(includeInstance: boolean, msg: Document): Document.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Document, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Document;
  static deserializeBinaryFromReader(message: Document, reader: jspb.BinaryReader): Document;
}

export namespace Document {
  export type AsObject = {
    id: number,
    meta?: common_meta_pb.Meta.AsObject,
    title: string,
    body: string,
    tagsList: Array<string>,
  }
}

