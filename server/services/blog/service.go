package blog

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"

	pb "github.com/asa-taka/practice-grpc-blog/out/go/blog"
	"github.com/asa-taka/practice-grpc-blog/server/services/util"
)

func RegisterNewServer(grpcServer *grpc.Server) {
	pb.RegisterBlogServiceServer(grpcServer, newServer())
}

type accountServiceServer struct {
	mu        sync.Mutex
	documents []*pb.Document
}

func newServer() *accountServiceServer {
	s := &accountServiceServer{}
	s.loadDocuments("server/data/documents.json")
	return s
}

func (s *accountServiceServer) loadDocuments(filePath string) {
	decoder := util.CreateArrayJSONDecoder(filePath)
	for decoder.More() {
		u := pb.Document{}
		if err := jsonpb.UnmarshalNext(decoder, &u); err != nil {
			log.Fatal(err)
		}
		s.documents = append(s.documents, &u)
	}
}

// rpc Handlers
// ------------

func (s *accountServiceServer) QueryDocument(ctx context.Context, in *pb.QueryDocumentsRequest) (*pb.QueryDocumentsResponse, error) {
	res := &pb.QueryDocumentsResponse{Documents: s.documents}
	return res, nil
}

func (s *accountServiceServer) GetDocument(ctx context.Context, in *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	d, err := s.getDocumentByID(in.Id)
	return &pb.GetDocumentResponse{Document: d}, err
}

func (s *accountServiceServer) CreateDocument(ctx context.Context, in *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	res := &pb.CreateDocumentResponse{Document: in.Document} // mock
	return res, nil
}

func (s *accountServiceServer) UpdateDocument(ctx context.Context, in *pb.UpdateDocumentRequest) (*pb.UpdateDocumentResponse, error) {
	res := &pb.UpdateDocumentResponse{} // mock
	return res, nil
}

func (s *accountServiceServer) DeleteDocument(ctx context.Context, in *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	d, err := s.getDocumentByID(in.Id) // mock
	return &pb.DeleteDocumentResponse{Document: d}, err
}

// Utilities

func (s *accountServiceServer) getDocumentByID(id int32) (*pb.Document, error) {
	for _, d := range s.documents {
		if d.Id == id {
			return d, nil
		}
	}
	return nil, fmt.Errorf("Document not found: %v", id)
}
