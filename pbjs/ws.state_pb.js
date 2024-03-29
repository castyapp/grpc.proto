// source: ws.state.proto
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

var ws_enums_pb = require('./ws.enums_pb.js');
goog.object.extend(proto, ws_enums_pb);
var grpc_user_pb = require('./grpc.user_pb.js');
goog.object.extend(proto, grpc_user_pb);
goog.exportSymbol('proto.proto.PersonalActivityMsgEvent', null, global);
goog.exportSymbol('proto.proto.PersonalStateMsgEvent', null, global);
goog.exportSymbol('proto.proto.PingMsgEvent', null, global);
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
proto.proto.PingMsgEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.PingMsgEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.PingMsgEvent.displayName = 'proto.proto.PingMsgEvent';
}
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
proto.proto.PersonalStateMsgEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.PersonalStateMsgEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.PersonalStateMsgEvent.displayName = 'proto.proto.PersonalStateMsgEvent';
}
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
proto.proto.PersonalActivityMsgEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.PersonalActivityMsgEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.PersonalActivityMsgEvent.displayName = 'proto.proto.PersonalActivityMsgEvent';
}



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
proto.proto.PingMsgEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.PingMsgEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.PingMsgEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PingMsgEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    state: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.proto.PingMsgEvent}
 */
proto.proto.PingMsgEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.PingMsgEvent;
  return proto.proto.PingMsgEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.PingMsgEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.PingMsgEvent}
 */
proto.proto.PingMsgEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.proto.PERSONAL_STATE} */ (reader.readEnum());
      msg.setState(value);
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
proto.proto.PingMsgEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.PingMsgEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.PingMsgEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PingMsgEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * optional PERSONAL_STATE state = 1;
 * @return {!proto.proto.PERSONAL_STATE}
 */
proto.proto.PingMsgEvent.prototype.getState = function() {
  return /** @type {!proto.proto.PERSONAL_STATE} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.proto.PERSONAL_STATE} value
 * @return {!proto.proto.PingMsgEvent} returns this
 */
proto.proto.PingMsgEvent.prototype.setState = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};





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
proto.proto.PersonalStateMsgEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.PersonalStateMsgEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.PersonalStateMsgEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PersonalStateMsgEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    user: (f = msg.getUser()) && grpc_user_pb.User.toObject(includeInstance, f),
    state: jspb.Message.getFieldWithDefault(msg, 2, 0),
    activity: (f = msg.getActivity()) && grpc_user_pb.Activity.toObject(includeInstance, f)
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
 * @return {!proto.proto.PersonalStateMsgEvent}
 */
proto.proto.PersonalStateMsgEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.PersonalStateMsgEvent;
  return proto.proto.PersonalStateMsgEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.PersonalStateMsgEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.PersonalStateMsgEvent}
 */
proto.proto.PersonalStateMsgEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new grpc_user_pb.User;
      reader.readMessage(value,grpc_user_pb.User.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 2:
      var value = /** @type {!proto.proto.PERSONAL_STATE} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 3:
      var value = new grpc_user_pb.Activity;
      reader.readMessage(value,grpc_user_pb.Activity.deserializeBinaryFromReader);
      msg.setActivity(value);
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
proto.proto.PersonalStateMsgEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.PersonalStateMsgEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.PersonalStateMsgEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PersonalStateMsgEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      grpc_user_pb.User.serializeBinaryToWriter
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getActivity();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      grpc_user_pb.Activity.serializeBinaryToWriter
    );
  }
};


/**
 * optional User user = 1;
 * @return {?proto.proto.User}
 */
proto.proto.PersonalStateMsgEvent.prototype.getUser = function() {
  return /** @type{?proto.proto.User} */ (
    jspb.Message.getWrapperField(this, grpc_user_pb.User, 1));
};


/**
 * @param {?proto.proto.User|undefined} value
 * @return {!proto.proto.PersonalStateMsgEvent} returns this
*/
proto.proto.PersonalStateMsgEvent.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.PersonalStateMsgEvent} returns this
 */
proto.proto.PersonalStateMsgEvent.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.PersonalStateMsgEvent.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional PERSONAL_STATE state = 2;
 * @return {!proto.proto.PERSONAL_STATE}
 */
proto.proto.PersonalStateMsgEvent.prototype.getState = function() {
  return /** @type {!proto.proto.PERSONAL_STATE} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.proto.PERSONAL_STATE} value
 * @return {!proto.proto.PersonalStateMsgEvent} returns this
 */
proto.proto.PersonalStateMsgEvent.prototype.setState = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional Activity activity = 3;
 * @return {?proto.proto.Activity}
 */
proto.proto.PersonalStateMsgEvent.prototype.getActivity = function() {
  return /** @type{?proto.proto.Activity} */ (
    jspb.Message.getWrapperField(this, grpc_user_pb.Activity, 3));
};


/**
 * @param {?proto.proto.Activity|undefined} value
 * @return {!proto.proto.PersonalStateMsgEvent} returns this
*/
proto.proto.PersonalStateMsgEvent.prototype.setActivity = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.PersonalStateMsgEvent} returns this
 */
proto.proto.PersonalStateMsgEvent.prototype.clearActivity = function() {
  return this.setActivity(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.PersonalStateMsgEvent.prototype.hasActivity = function() {
  return jspb.Message.getField(this, 3) != null;
};





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
proto.proto.PersonalActivityMsgEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.PersonalActivityMsgEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.PersonalActivityMsgEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PersonalActivityMsgEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    user: (f = msg.getUser()) && grpc_user_pb.User.toObject(includeInstance, f),
    activity: (f = msg.getActivity()) && grpc_user_pb.Activity.toObject(includeInstance, f)
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
 * @return {!proto.proto.PersonalActivityMsgEvent}
 */
proto.proto.PersonalActivityMsgEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.PersonalActivityMsgEvent;
  return proto.proto.PersonalActivityMsgEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.PersonalActivityMsgEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.PersonalActivityMsgEvent}
 */
proto.proto.PersonalActivityMsgEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new grpc_user_pb.User;
      reader.readMessage(value,grpc_user_pb.User.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 2:
      var value = new grpc_user_pb.Activity;
      reader.readMessage(value,grpc_user_pb.Activity.deserializeBinaryFromReader);
      msg.setActivity(value);
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
proto.proto.PersonalActivityMsgEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.PersonalActivityMsgEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.PersonalActivityMsgEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.PersonalActivityMsgEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      grpc_user_pb.User.serializeBinaryToWriter
    );
  }
  f = message.getActivity();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      grpc_user_pb.Activity.serializeBinaryToWriter
    );
  }
};


/**
 * optional User user = 1;
 * @return {?proto.proto.User}
 */
proto.proto.PersonalActivityMsgEvent.prototype.getUser = function() {
  return /** @type{?proto.proto.User} */ (
    jspb.Message.getWrapperField(this, grpc_user_pb.User, 1));
};


/**
 * @param {?proto.proto.User|undefined} value
 * @return {!proto.proto.PersonalActivityMsgEvent} returns this
*/
proto.proto.PersonalActivityMsgEvent.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.PersonalActivityMsgEvent} returns this
 */
proto.proto.PersonalActivityMsgEvent.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.PersonalActivityMsgEvent.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Activity activity = 2;
 * @return {?proto.proto.Activity}
 */
proto.proto.PersonalActivityMsgEvent.prototype.getActivity = function() {
  return /** @type{?proto.proto.Activity} */ (
    jspb.Message.getWrapperField(this, grpc_user_pb.Activity, 2));
};


/**
 * @param {?proto.proto.Activity|undefined} value
 * @return {!proto.proto.PersonalActivityMsgEvent} returns this
*/
proto.proto.PersonalActivityMsgEvent.prototype.setActivity = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.PersonalActivityMsgEvent} returns this
 */
proto.proto.PersonalActivityMsgEvent.prototype.clearActivity = function() {
  return this.setActivity(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.PersonalActivityMsgEvent.prototype.hasActivity = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.proto);
