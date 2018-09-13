package blog

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/asa-taka/practice-grpc-blog/out/go/blog"
)

func RegisterNewServer(grpcServer *grpc.Server) {
	pb.RegisterBlogServiceServer(grpcServer, newServer())
}

type accountServiceServer struct {
	mu        sync.Mutex
	documents []*pb.Document
}

func newServer() *accountServiceServer {
	return &accountServiceServer{
		documents: loadDocuments(),
	}
}

func loadDocuments() []*pb.Document {
	jsonFile, err := os.Open("server/data/documents.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var docs []*pb.Document
	json.Unmarshal(byteValue, &docs)

	return docs
}

// rpc Handlers
// ------------

func (s *accountServiceServer) QueryDocument(ctx context.Context, in *pb.QueryDocumentsRequest) (*pb.QueryDocumentsResponse, error) {
	res := &pb.QueryDocumentsResponse{Documents: s.documents}
	return res, nil
}

func (s *accountServiceServer) GetDocument(ctx context.Context, in *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	res := &pb.GetDocumentResponse{Document: s.documents[0]}
	return res, nil
}

func (s *accountServiceServer) CreateDocument(ctx context.Context, in *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	res := &pb.CreateDocumentResponse{}
	return res, nil
}

func (s *accountServiceServer) UpdateDocument(ctx context.Context, in *pb.UpdateDocumentRequest) (*pb.UpdateDocumentResponse, error) {
	res := &pb.UpdateDocumentResponse{}
	return res, nil
}

func (s *accountServiceServer) DeleteDocument(ctx context.Context, in *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	res := &pb.DeleteDocumentResponse{}
	return res, nil
}
