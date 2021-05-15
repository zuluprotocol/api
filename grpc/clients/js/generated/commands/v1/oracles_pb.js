// source: commands/v1/oracles.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.vega.commands.v1.OracleDataSubmission', null, global);
goog.exportSymbol('proto.vega.commands.v1.OracleDataSubmission.OracleSource', null, global);
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
proto.vega.commands.v1.OracleDataSubmission = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.vega.commands.v1.OracleDataSubmission, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.vega.commands.v1.OracleDataSubmission.displayName = 'proto.vega.commands.v1.OracleDataSubmission';
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
proto.vega.commands.v1.OracleDataSubmission.prototype.toObject = function(opt_includeInstance) {
  return proto.vega.commands.v1.OracleDataSubmission.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.vega.commands.v1.OracleDataSubmission} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.vega.commands.v1.OracleDataSubmission.toObject = function(includeInstance, msg) {
  var f, obj = {
    source: jspb.Message.getFieldWithDefault(msg, 1, 0),
    payload: msg.getPayload_asB64()
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
 * @return {!proto.vega.commands.v1.OracleDataSubmission}
 */
proto.vega.commands.v1.OracleDataSubmission.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.vega.commands.v1.OracleDataSubmission;
  return proto.vega.commands.v1.OracleDataSubmission.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.vega.commands.v1.OracleDataSubmission} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.vega.commands.v1.OracleDataSubmission}
 */
proto.vega.commands.v1.OracleDataSubmission.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.vega.commands.v1.OracleDataSubmission.OracleSource} */ (reader.readEnum());
      msg.setSource(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setPayload(value);
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
proto.vega.commands.v1.OracleDataSubmission.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.vega.commands.v1.OracleDataSubmission.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.vega.commands.v1.OracleDataSubmission} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.vega.commands.v1.OracleDataSubmission.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSource();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPayload_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.vega.commands.v1.OracleDataSubmission.OracleSource = {
  ORACLE_SOURCE_UNSPECIFIED: 0,
  ORACLE_SOURCE_OPEN_ORACLE: 1
};

/**
 * optional OracleSource source = 1;
 * @return {!proto.vega.commands.v1.OracleDataSubmission.OracleSource}
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.getSource = function() {
  return /** @type {!proto.vega.commands.v1.OracleDataSubmission.OracleSource} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.vega.commands.v1.OracleDataSubmission.OracleSource} value
 * @return {!proto.vega.commands.v1.OracleDataSubmission} returns this
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.setSource = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional bytes payload = 2;
 * @return {!(string|Uint8Array)}
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.getPayload = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes payload = 2;
 * This is a type-conversion wrapper around `getPayload()`
 * @return {string}
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.getPayload_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getPayload()));
};


/**
 * optional bytes payload = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getPayload()`
 * @return {!Uint8Array}
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.getPayload_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getPayload()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.vega.commands.v1.OracleDataSubmission} returns this
 */
proto.vega.commands.v1.OracleDataSubmission.prototype.setPayload = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
};


goog.object.extend(exports, proto.vega.commands.v1);
