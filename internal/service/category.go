package service

import(
	"github.com/devfullcycle/AulaGRPC/internal/database"
	"github.com/devfullcycle/AulaGRPC/internal/pb"
)

type CategoryService struct{

	pb.mustEmbedUnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(CategoryDB database.Category) * CategoryService{
	return &CategoryService{
		CategoryDB: CategoryDB,
	}
}
func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}