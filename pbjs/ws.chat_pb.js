// source: ws.chat.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var grpc_user_pb = require('./grpc.user_pb.js');
goog.object.extend(proto, grpc_user_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.proto.ChatMsgEvent', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.ChatMsgEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.ChatMsgEvent.repeatedFields_, null);
};
goog.inherits(proto.proto.ChatMsgEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.ChatMsgEvent.displayName = 'proto.proto.ChatMsgEvent';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.ChatMsgEvent.repeatedFields_ = [5,6];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.ChatMsgEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.ChatMsgEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.ChatMsgEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.ChatMsgEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    message: msg.getMessage_asB64(),
    from: jspb.Message.getFieldWithDefault(msg, 2, ""),
    to: jspb.Message.getFieldWithDefault(msg, 3, ""),
    user: (f = msg.getUser()) && grpc_user_pb.User.toObject(includeInstance, f),
    emojiesList: (f = jspb.Message.getRepeatedField(msg, 5)) == null ? undefined : f,
    mentionsList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
    createdAt: (f = msg.getCreatedAt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.ChatMsgEvent}
 */
proto.proto.ChatMsgEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.ChatMsgEvent;
  return proto.proto.ChatMsgEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.ChatMsgEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.ChatMsgEvent}
 */
proto.proto.ChatMsgEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setMessage(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFrom(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setTo(value);
      break;
    case 4:
      var value = new grpc_user_pb.User;
      reader.readMessage(value,grpc_user_pb.User.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 5:
      var value = /** @type {!Array<number>} */ (reader.readPackedUint64());
      msg.setEmojiesList(value);
      break;
    case 6:
      var value = /** @type {!Array<number>} */ (reader.readPackedUint64());
      msg.setMentionsList(value);
      break;
    case 7:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedAt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.ChatMsgEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.ChatMsgEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.ChatMsgEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.ChatMsgEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMessage_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getFrom();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTo();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      grpc_user_pb.User.serializeBinaryToWriter
    );
  }
  f = message.getEmojiesList();
  if (f.length > 0) {
    writer.writePackedUint64(
      5,
      f
    );
  }
  f = message.getMentionsList();
  if (f.length > 0) {
    writer.writePackedUint64(
      6,
      f
    );
  }
  f = message.getCreatedAt();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes message = 1;
 * @return {string}
 */
proto.proto.ChatMsgEvent.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes message = 1;
 * This is a type-conversion wrapper around `getMessage()`
 * @return {string}
 */
proto.proto.ChatMsgEvent.prototype.getMessage_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getMessage()));
};


/**
 * optional bytes message = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getMessage()`
 * @return {!Uint8Array}
 */
proto.proto.ChatMsgEvent.prototype.getMessage_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getMessage()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.setMessage = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional string from = 2;
 * @return {string}
 */
proto.proto.ChatMsgEvent.prototype.getFrom = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.setFrom = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string to = 3;
 * @return {string}
 */
proto.proto.ChatMsgEvent.prototype.getTo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.setTo = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional User user = 4;
 * @return {?proto.proto.User}
 */
proto.proto.ChatMsgEvent.prototype.getUser = function() {
  return /** @type{?proto.proto.User} */ (
    jspb.Message.getWrapperField(this, grpc_user_pb.User, 4));
};


/**
 * @param {?proto.proto.User|undefined} value
 * @return {!proto.proto.ChatMsgEvent} returns this
*/
proto.proto.ChatMsgEvent.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.ChatMsgEvent.prototype.hasUser = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * repeated uint64 emojies = 5;
 * @return {!Array<number>}
 */
proto.proto.ChatMsgEvent.prototype.getEmojiesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 5));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.setEmojiesList = function(value) {
  return jspb.Message.setField(this, 5, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.addEmojies = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 5, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.clearEmojiesList = function() {
  return this.setEmojiesList([]);
};


/**
 * repeated uint64 mentions = 6;
 * @return {!Array<number>}
 */
proto.proto.ChatMsgEvent.prototype.getMentionsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.setMentionsList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.addMentions = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.clearMentionsList = function() {
  return this.setMentionsList([]);
};


/**
 * optional google.protobuf.Timestamp created_at = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.proto.ChatMsgEvent.prototype.getCreatedAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 7));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.proto.ChatMsgEvent} returns this
*/
proto.proto.ChatMsgEvent.prototype.setCreatedAt = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.ChatMsgEvent} returns this
 */
proto.proto.ChatMsgEvent.prototype.clearCreatedAt = function() {
  return this.setCreatedAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.ChatMsgEvent.prototype.hasCreatedAt = function() {
  return jspb.Message.getField(this, 7) != null;
};


goog.object.extend(exports, proto.proto);