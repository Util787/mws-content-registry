import grpc
from concurrent import futures
from llm_grpc import llm_service_pb2_grpc, llm_service_pb2
from .model_runner import ModelRunner
from .config import GRPC_PORT

model = ModelRunner()


class LLMService(llm_service_pb2_grpc.LLMServiceServicer):
    def Analyze(self, request, context):
        result = model.process(
            request.user_query, request.stats_json, request.context)
        return llm_service_pb2.AnalyzeResponse(**result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    llm_service_pb2_grpc.add_LLMServiceServicer_to_server(LLMService(), server)
    server.add_insecure_port(f"[::]:{GRPC_PORT}")
    server.start()
    server.wait_for_termination()
