// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Operations defines a contract for the operations described by OpenAPI v3 specification.
type defaultOperations interface {
	// OnlyForm implements onlyForm operation.
	//
	// POST /onlyForm
	OnlyForm(ctx context.Context, req *OnlyFormReq) error
	// OnlyMultipartFile implements onlyMultipartFile operation.
	//
	// POST /onlyMultipartFile
	OnlyMultipartFile(ctx context.Context, req *OnlyMultipartFileReq) error
	// OnlyMultipartForm implements onlyMultipartForm operation.
	//
	// POST /onlyMultipartForm
	OnlyMultipartForm(ctx context.Context, req *OnlyMultipartFormReq) error
	// TestFormURLEncoded implements testFormURLEncoded operation.
	//
	// POST /testFormURLEncoded
	TestFormURLEncoded(ctx context.Context, req *TestForm) error
	// TestMultipart implements testMultipart operation.
	//
	// POST /testMultipart
	TestMultipart(ctx context.Context, req *TestFormMultipart) error
	// TestMultipartUpload implements testMultipartUpload operation.
	//
	// POST /testMultipartUpload
	TestMultipartUpload(ctx context.Context, req *TestMultipartUploadReq) (*TestMultipartUploadOK, error)
	// TestReuseFormOptionalSchema implements testReuseFormOptionalSchema operation.
	//
	// POST /testReuseFormOptionalSchema
	TestReuseFormOptionalSchema(ctx context.Context, req OptSharedRequestMultipart) error
	// TestReuseFormSchema implements testReuseFormSchema operation.
	//
	// POST /testReuseFormSchema
	TestReuseFormSchema(ctx context.Context, req *SharedRequestMultipart) error
	// TestShareFormSchema implements testShareFormSchema operation.
	//
	// POST /testShareFormSchema
	TestShareFormSchema(ctx context.Context, req TestShareFormSchemaReq) error
}

type Operations interface {
	defaultOperations
}
