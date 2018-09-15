// package: blog.document
// file: document/service.proto

var document_service_pb = require("../document/service_pb");
var grpc = require("grpc-web-client").grpc;

var BlogService = (function () {
  function BlogService() {}
  BlogService.serviceName = "blog.document.BlogService";
  return BlogService;
}());

BlogService.QueryDocument = {
  methodName: "QueryDocument",
  service: BlogService,
  requestStream: false,
  responseStream: false,
  requestType: document_service_pb.QueryDocumentsRequest,
  responseType: document_service_pb.QueryDocumentsResponse
};

BlogService.GetDocument = {
  methodName: "GetDocument",
  service: BlogService,
  requestStream: false,
  responseStream: false,
  requestType: document_service_pb.GetDocumentRequest,
  responseType: document_service_pb.GetDocumentResponse
};

BlogService.CreateDocument = {
  methodName: "CreateDocument",
  service: BlogService,
  requestStream: false,
  responseStream: false,
  requestType: document_service_pb.CreateDocumentRequest,
  responseType: document_service_pb.CreateDocumentResponse
};

BlogService.DeleteDocument = {
  methodName: "DeleteDocument",
  service: BlogService,
  requestStream: false,
  responseStream: false,
  requestType: document_service_pb.DeleteDocumentRequest,
  responseType: document_service_pb.DeleteDocumentResponse
};

BlogService.UpdateDocument = {
  methodName: "UpdateDocument",
  service: BlogService,
  requestStream: false,
  responseStream: false,
  requestType: document_service_pb.UpdateDocumentRequest,
  responseType: document_service_pb.UpdateDocumentResponse
};

exports.BlogService = BlogService;

function BlogServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BlogServiceClient.prototype.queryDocument = function queryDocument(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BlogService.QueryDocument, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

BlogServiceClient.prototype.getDocument = function getDocument(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BlogService.GetDocument, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

BlogServiceClient.prototype.createDocument = function createDocument(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BlogService.CreateDocument, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

BlogServiceClient.prototype.deleteDocument = function deleteDocument(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BlogService.DeleteDocument, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

BlogServiceClient.prototype.updateDocument = function updateDocument(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BlogService.UpdateDocument, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.BlogServiceClient = BlogServiceClient;

