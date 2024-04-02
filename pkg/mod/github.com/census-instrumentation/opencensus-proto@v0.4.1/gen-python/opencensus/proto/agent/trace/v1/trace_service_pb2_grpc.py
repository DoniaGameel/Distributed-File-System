# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from opencensus.proto.agent.trace.v1 import trace_service_pb2 as opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2


class TraceServiceStub(object):
    """Service that can be used to push spans and configs between one Application
    instrumented with OpenCensus and an agent, or between an agent and a
    central collector or config service (in this case spans and configs are
    sent/received to/from multiple Applications).
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Config = channel.stream_stream(
                '/opencensus.proto.agent.trace.v1.TraceService/Config',
                request_serializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.CurrentLibraryConfig.SerializeToString,
                response_deserializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.UpdatedLibraryConfig.FromString,
                )
        self.Export = channel.stream_stream(
                '/opencensus.proto.agent.trace.v1.TraceService/Export',
                request_serializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceRequest.SerializeToString,
                response_deserializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceResponse.FromString,
                )


class TraceServiceServicer(object):
    """Service that can be used to push spans and configs between one Application
    instrumented with OpenCensus and an agent, or between an agent and a
    central collector or config service (in this case spans and configs are
    sent/received to/from multiple Applications).
    """

    def Config(self, request_iterator, context):
        """After initialization, this RPC must be kept alive for the entire life of
        the application. The agent pushes configs down to applications via a
        stream.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Export(self, request_iterator, context):
        """For performance reasons, it is recommended to keep this RPC
        alive for the entire life of the application.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TraceServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Config': grpc.stream_stream_rpc_method_handler(
                    servicer.Config,
                    request_deserializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.CurrentLibraryConfig.FromString,
                    response_serializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.UpdatedLibraryConfig.SerializeToString,
            ),
            'Export': grpc.stream_stream_rpc_method_handler(
                    servicer.Export,
                    request_deserializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceRequest.FromString,
                    response_serializer=opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'opencensus.proto.agent.trace.v1.TraceService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TraceService(object):
    """Service that can be used to push spans and configs between one Application
    instrumented with OpenCensus and an agent, or between an agent and a
    central collector or config service (in this case spans and configs are
    sent/received to/from multiple Applications).
    """

    @staticmethod
    def Config(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/opencensus.proto.agent.trace.v1.TraceService/Config',
            opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.CurrentLibraryConfig.SerializeToString,
            opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.UpdatedLibraryConfig.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Export(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/opencensus.proto.agent.trace.v1.TraceService/Export',
            opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceRequest.SerializeToString,
            opencensus_dot_proto_dot_agent_dot_trace_dot_v1_dot_trace__service__pb2.ExportTraceServiceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
