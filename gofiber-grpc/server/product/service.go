package product

import (
	"context"
	"fmt"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/proto"
	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/server/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductServiceServer service for the server
type ProductServiceServer struct{}

// CreateProduct service to CreateProduct
func (p *ProductServiceServer) CreateProduct(ctx context.Context, request *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	// get the protobuf product type from the request type
	product := request.GetProduct()

	// convert protobuf Product to ProductItem to add in the database
	productItem := &database.ProductItem{
		Name:  product.GetName(),
		Type:  product.GetType(),
		Prize: product.GetPrize(),
	}

	err := CreateProductQuery(productItem)
	if err != nil {
		// return grpc error
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// add product id
	product.Id = productItem.ID

	response := &proto.Response{Status: 200, Message: "success"}
	return &proto.CreateProductResponse{Product: product, Response: response}, nil
}

// UpdateProduct service to UpdateProduct
func (p *ProductServiceServer) UpdateProduct(ctx context.Context, request *proto.UpdateProductRequest) (*proto.UpdateProductResponse, error) {
	product := request.GetProduct()

	// convert protobuf Product to ProductItem to add in the database
	productItem := &database.ProductItem{
		ID:    product.GetId(),
		Name:  product.GetName(),
		Type:  product.GetType(),
		Prize: product.GetPrize(),
	}

	err := UpdateProductQuery(productItem)
	if err != nil {
		// return invalid argument grpc error
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not update the product: %v", err),
		)
	}

	product.Name = productItem.Name
	product.Type = productItem.Type
	product.Prize = productItem.Prize

	response := &proto.Response{Status: 200, Message: "updated"}
	return &proto.UpdateProductResponse{Product: product, Response: response}, nil
}

// DeleteProduct service to DeleteProduct
func (p *ProductServiceServer) DeleteProduct(ctx context.Context, request *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	id := request.GetId()

	err := DeleteProductQuery(id)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Id didn't match please check the id: %v", err),
		)
	}

	// delete product using product id
	response := &proto.Response{Status: 200, Message: "deleted"}
	return &proto.DeleteProductResponse{Success: true, Response: response}, nil
}

// GetProduct service to GetProduct
func (p *ProductServiceServer) GetProduct(ctx context.Context, request *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	id := request.GetId()

	productItem, err := ReadProductQuery(id)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Id didn't match please check the id: %v", err),
		)
	}

	product := &proto.Product{
		Id:    productItem.ID,
		Name:  productItem.Name,
		Type:  productItem.Type,
		Prize: productItem.Prize,
	}

	// get product using product id
	response := &proto.Response{Status: 201, Message: "success"}
	return &proto.GetProductResponse{Product: product, Response: response}, nil
}

// GetProducts service to get products
func (p *ProductServiceServer) GetProducts(request *proto.GetProductsRequest, stream proto.ProductService_GetProductsServer) error {
	productItems, err := ReadProductsQuery()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to retreive any products: %v", err),
		)
	}

	for _, productItem := range productItems {
		product := &proto.Product{
			Id:    productItem.ID,
			Name:  productItem.Name,
			Type:  productItem.Type,
			Prize: productItem.Prize,
		}

		response := &proto.Response{Status: 201, Message: "success"}
		stream.Send(&proto.GetProductsResponse{Product: product, Response: response})

	}

	return nil
}
