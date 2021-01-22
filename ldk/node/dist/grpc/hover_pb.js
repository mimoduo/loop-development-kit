"use strict";
// source: hover.proto
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
var session_pb = require('./session_pb.js');
goog.object.extend(proto, session_pb);
goog.exportSymbol('proto.proto.HoverReadRequest', null, global);
goog.exportSymbol('proto.proto.HoverReadResponse', null, global);
goog.exportSymbol('proto.proto.HoverReadStreamRequest', null, global);
goog.exportSymbol('proto.proto.HoverReadStreamResponse', null, global);
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
proto.proto.HoverReadRequest = function (opt_data) {
    jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.HoverReadRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
    /**
     * @public
     * @override
     */
    proto.proto.HoverReadRequest.displayName = 'proto.proto.HoverReadRequest';
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
proto.proto.HoverReadStreamRequest = function (opt_data) {
    jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.HoverReadStreamRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
    /**
     * @public
     * @override
     */
    proto.proto.HoverReadStreamRequest.displayName = 'proto.proto.HoverReadStreamRequest';
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
proto.proto.HoverReadResponse = function (opt_data) {
    jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.HoverReadResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
    /**
     * @public
     * @override
     */
    proto.proto.HoverReadResponse.displayName = 'proto.proto.HoverReadResponse';
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
proto.proto.HoverReadStreamResponse = function (opt_data) {
    jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.HoverReadStreamResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
    /**
     * @public
     * @override
     */
    proto.proto.HoverReadStreamResponse.displayName = 'proto.proto.HoverReadStreamResponse';
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
    proto.proto.HoverReadRequest.prototype.toObject = function (opt_includeInstance) {
        return proto.proto.HoverReadRequest.toObject(opt_includeInstance, this);
    };
    /**
     * Static version of the {@see toObject} method.
     * @param {boolean|undefined} includeInstance Deprecated. Whether to include
     *     the JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @param {!proto.proto.HoverReadRequest} msg The msg instance to transform.
     * @return {!Object}
     * @suppress {unusedLocalVariables} f is only used for nested messages
     */
    proto.proto.HoverReadRequest.toObject = function (includeInstance, msg) {
        var f, obj = {
            session: (f = msg.getSession()) && session_pb.Session.toObject(includeInstance, f),
            xfromcenter: jspb.Message.getFieldWithDefault(msg, 2, 0),
            yfromcenter: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.proto.HoverReadRequest}
 */
proto.proto.HoverReadRequest.deserializeBinary = function (bytes) {
    var reader = new jspb.BinaryReader(bytes);
    var msg = new proto.proto.HoverReadRequest;
    return proto.proto.HoverReadRequest.deserializeBinaryFromReader(msg, reader);
};
/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.HoverReadRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.HoverReadRequest}
 */
proto.proto.HoverReadRequest.deserializeBinaryFromReader = function (msg, reader) {
    while (reader.nextField()) {
        if (reader.isEndGroup()) {
            break;
        }
        var field = reader.getFieldNumber();
        switch (field) {
            case 1:
                var value = new session_pb.Session;
                reader.readMessage(value, session_pb.Session.deserializeBinaryFromReader);
                msg.setSession(value);
                break;
            case 2:
                var value = /** @type {number} */ (reader.readUint32());
                msg.setXfromcenter(value);
                break;
            case 3:
                var value = /** @type {number} */ (reader.readUint32());
                msg.setYfromcenter(value);
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
proto.proto.HoverReadRequest.prototype.serializeBinary = function () {
    var writer = new jspb.BinaryWriter();
    proto.proto.HoverReadRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
};
/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.HoverReadRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HoverReadRequest.serializeBinaryToWriter = function (message, writer) {
    var f = undefined;
    f = message.getSession();
    if (f != null) {
        writer.writeMessage(1, f, session_pb.Session.serializeBinaryToWriter);
    }
    f = message.getXfromcenter();
    if (f !== 0) {
        writer.writeUint32(2, f);
    }
    f = message.getYfromcenter();
    if (f !== 0) {
        writer.writeUint32(3, f);
    }
};
/**
 * optional Session session = 1;
 * @return {?proto.proto.Session}
 */
proto.proto.HoverReadRequest.prototype.getSession = function () {
    return /** @type{?proto.proto.Session} */ (jspb.Message.getWrapperField(this, session_pb.Session, 1));
};
/**
 * @param {?proto.proto.Session|undefined} value
 * @return {!proto.proto.HoverReadRequest} returns this
*/
proto.proto.HoverReadRequest.prototype.setSession = function (value) {
    return jspb.Message.setWrapperField(this, 1, value);
};
/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.HoverReadRequest} returns this
 */
proto.proto.HoverReadRequest.prototype.clearSession = function () {
    return this.setSession(undefined);
};
/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.HoverReadRequest.prototype.hasSession = function () {
    return jspb.Message.getField(this, 1) != null;
};
/**
 * optional uint32 xFromCenter = 2;
 * @return {number}
 */
proto.proto.HoverReadRequest.prototype.getXfromcenter = function () {
    return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};
/**
 * @param {number} value
 * @return {!proto.proto.HoverReadRequest} returns this
 */
proto.proto.HoverReadRequest.prototype.setXfromcenter = function (value) {
    return jspb.Message.setProto3IntField(this, 2, value);
};
/**
 * optional uint32 yFromCenter = 3;
 * @return {number}
 */
proto.proto.HoverReadRequest.prototype.getYfromcenter = function () {
    return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};
/**
 * @param {number} value
 * @return {!proto.proto.HoverReadRequest} returns this
 */
proto.proto.HoverReadRequest.prototype.setYfromcenter = function (value) {
    return jspb.Message.setProto3IntField(this, 3, value);
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
    proto.proto.HoverReadStreamRequest.prototype.toObject = function (opt_includeInstance) {
        return proto.proto.HoverReadStreamRequest.toObject(opt_includeInstance, this);
    };
    /**
     * Static version of the {@see toObject} method.
     * @param {boolean|undefined} includeInstance Deprecated. Whether to include
     *     the JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @param {!proto.proto.HoverReadStreamRequest} msg The msg instance to transform.
     * @return {!Object}
     * @suppress {unusedLocalVariables} f is only used for nested messages
     */
    proto.proto.HoverReadStreamRequest.toObject = function (includeInstance, msg) {
        var f, obj = {
            session: (f = msg.getSession()) && session_pb.Session.toObject(includeInstance, f),
            xfromcenter: jspb.Message.getFieldWithDefault(msg, 2, 0),
            yfromcenter: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.proto.HoverReadStreamRequest}
 */
proto.proto.HoverReadStreamRequest.deserializeBinary = function (bytes) {
    var reader = new jspb.BinaryReader(bytes);
    var msg = new proto.proto.HoverReadStreamRequest;
    return proto.proto.HoverReadStreamRequest.deserializeBinaryFromReader(msg, reader);
};
/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.HoverReadStreamRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.HoverReadStreamRequest}
 */
proto.proto.HoverReadStreamRequest.deserializeBinaryFromReader = function (msg, reader) {
    while (reader.nextField()) {
        if (reader.isEndGroup()) {
            break;
        }
        var field = reader.getFieldNumber();
        switch (field) {
            case 1:
                var value = new session_pb.Session;
                reader.readMessage(value, session_pb.Session.deserializeBinaryFromReader);
                msg.setSession(value);
                break;
            case 2:
                var value = /** @type {number} */ (reader.readUint32());
                msg.setXfromcenter(value);
                break;
            case 3:
                var value = /** @type {number} */ (reader.readUint32());
                msg.setYfromcenter(value);
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
proto.proto.HoverReadStreamRequest.prototype.serializeBinary = function () {
    var writer = new jspb.BinaryWriter();
    proto.proto.HoverReadStreamRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
};
/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.HoverReadStreamRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HoverReadStreamRequest.serializeBinaryToWriter = function (message, writer) {
    var f = undefined;
    f = message.getSession();
    if (f != null) {
        writer.writeMessage(1, f, session_pb.Session.serializeBinaryToWriter);
    }
    f = message.getXfromcenter();
    if (f !== 0) {
        writer.writeUint32(2, f);
    }
    f = message.getYfromcenter();
    if (f !== 0) {
        writer.writeUint32(3, f);
    }
};
/**
 * optional Session session = 1;
 * @return {?proto.proto.Session}
 */
proto.proto.HoverReadStreamRequest.prototype.getSession = function () {
    return /** @type{?proto.proto.Session} */ (jspb.Message.getWrapperField(this, session_pb.Session, 1));
};
/**
 * @param {?proto.proto.Session|undefined} value
 * @return {!proto.proto.HoverReadStreamRequest} returns this
*/
proto.proto.HoverReadStreamRequest.prototype.setSession = function (value) {
    return jspb.Message.setWrapperField(this, 1, value);
};
/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.HoverReadStreamRequest} returns this
 */
proto.proto.HoverReadStreamRequest.prototype.clearSession = function () {
    return this.setSession(undefined);
};
/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.HoverReadStreamRequest.prototype.hasSession = function () {
    return jspb.Message.getField(this, 1) != null;
};
/**
 * optional uint32 xFromCenter = 2;
 * @return {number}
 */
proto.proto.HoverReadStreamRequest.prototype.getXfromcenter = function () {
    return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};
/**
 * @param {number} value
 * @return {!proto.proto.HoverReadStreamRequest} returns this
 */
proto.proto.HoverReadStreamRequest.prototype.setXfromcenter = function (value) {
    return jspb.Message.setProto3IntField(this, 2, value);
};
/**
 * optional uint32 yFromCenter = 3;
 * @return {number}
 */
proto.proto.HoverReadStreamRequest.prototype.getYfromcenter = function () {
    return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};
/**
 * @param {number} value
 * @return {!proto.proto.HoverReadStreamRequest} returns this
 */
proto.proto.HoverReadStreamRequest.prototype.setYfromcenter = function (value) {
    return jspb.Message.setProto3IntField(this, 3, value);
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
    proto.proto.HoverReadResponse.prototype.toObject = function (opt_includeInstance) {
        return proto.proto.HoverReadResponse.toObject(opt_includeInstance, this);
    };
    /**
     * Static version of the {@see toObject} method.
     * @param {boolean|undefined} includeInstance Deprecated. Whether to include
     *     the JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @param {!proto.proto.HoverReadResponse} msg The msg instance to transform.
     * @return {!Object}
     * @suppress {unusedLocalVariables} f is only used for nested messages
     */
    proto.proto.HoverReadResponse.toObject = function (includeInstance, msg) {
        var f, obj = {
            text: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.proto.HoverReadResponse}
 */
proto.proto.HoverReadResponse.deserializeBinary = function (bytes) {
    var reader = new jspb.BinaryReader(bytes);
    var msg = new proto.proto.HoverReadResponse;
    return proto.proto.HoverReadResponse.deserializeBinaryFromReader(msg, reader);
};
/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.HoverReadResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.HoverReadResponse}
 */
proto.proto.HoverReadResponse.deserializeBinaryFromReader = function (msg, reader) {
    while (reader.nextField()) {
        if (reader.isEndGroup()) {
            break;
        }
        var field = reader.getFieldNumber();
        switch (field) {
            case 1:
                var value = /** @type {string} */ (reader.readString());
                msg.setText(value);
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
proto.proto.HoverReadResponse.prototype.serializeBinary = function () {
    var writer = new jspb.BinaryWriter();
    proto.proto.HoverReadResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
};
/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.HoverReadResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HoverReadResponse.serializeBinaryToWriter = function (message, writer) {
    var f = undefined;
    f = message.getText();
    if (f.length > 0) {
        writer.writeString(1, f);
    }
};
/**
 * optional string text = 1;
 * @return {string}
 */
proto.proto.HoverReadResponse.prototype.getText = function () {
    return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};
/**
 * @param {string} value
 * @return {!proto.proto.HoverReadResponse} returns this
 */
proto.proto.HoverReadResponse.prototype.setText = function (value) {
    return jspb.Message.setProto3StringField(this, 1, value);
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
    proto.proto.HoverReadStreamResponse.prototype.toObject = function (opt_includeInstance) {
        return proto.proto.HoverReadStreamResponse.toObject(opt_includeInstance, this);
    };
    /**
     * Static version of the {@see toObject} method.
     * @param {boolean|undefined} includeInstance Deprecated. Whether to include
     *     the JSPB instance for transitional soy proto support:
     *     http://goto/soy-param-migration
     * @param {!proto.proto.HoverReadStreamResponse} msg The msg instance to transform.
     * @return {!Object}
     * @suppress {unusedLocalVariables} f is only used for nested messages
     */
    proto.proto.HoverReadStreamResponse.toObject = function (includeInstance, msg) {
        var f, obj = {
            text: jspb.Message.getFieldWithDefault(msg, 1, ""),
            error: jspb.Message.getFieldWithDefault(msg, 15, "")
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
 * @return {!proto.proto.HoverReadStreamResponse}
 */
proto.proto.HoverReadStreamResponse.deserializeBinary = function (bytes) {
    var reader = new jspb.BinaryReader(bytes);
    var msg = new proto.proto.HoverReadStreamResponse;
    return proto.proto.HoverReadStreamResponse.deserializeBinaryFromReader(msg, reader);
};
/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.HoverReadStreamResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.HoverReadStreamResponse}
 */
proto.proto.HoverReadStreamResponse.deserializeBinaryFromReader = function (msg, reader) {
    while (reader.nextField()) {
        if (reader.isEndGroup()) {
            break;
        }
        var field = reader.getFieldNumber();
        switch (field) {
            case 1:
                var value = /** @type {string} */ (reader.readString());
                msg.setText(value);
                break;
            case 15:
                var value = /** @type {string} */ (reader.readString());
                msg.setError(value);
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
proto.proto.HoverReadStreamResponse.prototype.serializeBinary = function () {
    var writer = new jspb.BinaryWriter();
    proto.proto.HoverReadStreamResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
};
/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.HoverReadStreamResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.HoverReadStreamResponse.serializeBinaryToWriter = function (message, writer) {
    var f = undefined;
    f = message.getText();
    if (f.length > 0) {
        writer.writeString(1, f);
    }
    f = message.getError();
    if (f.length > 0) {
        writer.writeString(15, f);
    }
};
/**
 * optional string text = 1;
 * @return {string}
 */
proto.proto.HoverReadStreamResponse.prototype.getText = function () {
    return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};
/**
 * @param {string} value
 * @return {!proto.proto.HoverReadStreamResponse} returns this
 */
proto.proto.HoverReadStreamResponse.prototype.setText = function (value) {
    return jspb.Message.setProto3StringField(this, 1, value);
};
/**
 * optional string error = 15;
 * @return {string}
 */
proto.proto.HoverReadStreamResponse.prototype.getError = function () {
    return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 15, ""));
};
/**
 * @param {string} value
 * @return {!proto.proto.HoverReadStreamResponse} returns this
 */
proto.proto.HoverReadStreamResponse.prototype.setError = function (value) {
    return jspb.Message.setProto3StringField(this, 15, value);
};
goog.object.extend(exports, proto.proto);
