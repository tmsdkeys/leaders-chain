/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "tmsdkeys.leaders.leaders";

export interface TopRanked {
  clientId: string;
  address: string;
  score: string;
}

const baseTopRanked: object = { clientId: "", address: "", score: "" };

export const TopRanked = {
  encode(message: TopRanked, writer: Writer = Writer.create()): Writer {
    if (message.clientId !== "") {
      writer.uint32(10).string(message.clientId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.score !== "") {
      writer.uint32(26).string(message.score);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TopRanked {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTopRanked } as TopRanked;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.clientId = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.score = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TopRanked {
    const message = { ...baseTopRanked } as TopRanked;
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = String(object.clientId);
    } else {
      message.clientId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = String(object.score);
    } else {
      message.score = "";
    }
    return message;
  },

  toJSON(message: TopRanked): unknown {
    const obj: any = {};
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.address !== undefined && (obj.address = message.address);
    message.score !== undefined && (obj.score = message.score);
    return obj;
  },

  fromPartial(object: DeepPartial<TopRanked>): TopRanked {
    const message = { ...baseTopRanked } as TopRanked;
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = object.clientId;
    } else {
      message.clientId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
