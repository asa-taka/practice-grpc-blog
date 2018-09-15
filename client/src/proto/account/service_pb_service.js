// package: blog.account
// file: account/service.proto

var account_service_pb = require("../account/service_pb");
var grpc = require("grpc-web-client").grpc;

var AccountService = (function () {
  function AccountService() {}
  AccountService.serviceName = "blog.account.AccountService";
  return AccountService;
}());

AccountService.QueryUsers = {
  methodName: "QueryUsers",
  service: AccountService,
  requestStream: false,
  responseStream: false,
  requestType: account_service_pb.QueryUsersRequest,
  responseType: account_service_pb.QueryUsersResponse
};

AccountService.GetUser = {
  methodName: "GetUser",
  service: AccountService,
  requestStream: false,
  responseStream: false,
  requestType: account_service_pb.GetUserRequest,
  responseType: account_service_pb.GetUserResponse
};

AccountService.CreateUser = {
  methodName: "CreateUser",
  service: AccountService,
  requestStream: false,
  responseStream: false,
  requestType: account_service_pb.CreateUserRequest,
  responseType: account_service_pb.CreateUserResponse
};

AccountService.DeleteUser = {
  methodName: "DeleteUser",
  service: AccountService,
  requestStream: false,
  responseStream: false,
  requestType: account_service_pb.DeleteUserRequest,
  responseType: account_service_pb.DeleteUserResponse
};

AccountService.UpdateUser = {
  methodName: "UpdateUser",
  service: AccountService,
  requestStream: false,
  responseStream: false,
  requestType: account_service_pb.UpdateUserRequest,
  responseType: account_service_pb.UpdateUserResponse
};

exports.AccountService = AccountService;

function AccountServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AccountServiceClient.prototype.queryUsers = function queryUsers(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(AccountService.QueryUsers, {
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

AccountServiceClient.prototype.getUser = function getUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(AccountService.GetUser, {
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

AccountServiceClient.prototype.createUser = function createUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(AccountService.CreateUser, {
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

AccountServiceClient.prototype.deleteUser = function deleteUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(AccountService.DeleteUser, {
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

AccountServiceClient.prototype.updateUser = function updateUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(AccountService.UpdateUser, {
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

exports.AccountServiceClient = AccountServiceClient;

