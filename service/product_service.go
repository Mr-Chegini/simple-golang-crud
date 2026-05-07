package service

import (
	"errors"

	"github.com/Mr-Chegini/simple-golang-crud/models"
	"github.com/Mr-Chegini/simple-golang-crud/repository"
	"github.com/Mr-Chegini/simple-golang-crud/utils"
)

// ProductServiceInterface defines business operations for products.
type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	ListProductsByUser(userID uint) ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

// ProductService implements ProductServiceInterface.
type ProductService struct {
	productRepo repository.ProductRepositoryInterface
	userRepo    repository.UserRepositoryInterface
}

// NewProductService creates a new ProductService.
func NewProductService(productRepo repository.ProductRepositoryInterface, userRepo repository.UserRepositoryInterface) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

// CreateProduct validates and creates a new product.
func (s *ProductService) CreateProduct(product *models.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	if !utils.IsNonEmpty(product.Name) {
		return errors.New("product name is required")
	}

	if !utils.IsPositivePrice(product.Price) {
		return errors.New("product price must be greater than zero")
	}

	if product.UserID != 0 {
		if _, err := s.userRepo.FindByID(product.UserID); err != nil {
			return errors.New("product owner user not found")
		}
	}

	return s.productRepo.Create(product)
}

// GetProductByID returns a single product.
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("invalid product id")
	}
	return s.productRepo.FindByID(id)
}

// ListProducts returns all products.
func (s *ProductService) ListProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

// ListProductsByUser returns products for a specific user.
func (s *ProductService) ListProductsByUser(userID uint) ([]models.Product, error) {
	if userID == 0 {
		return nil, errors.New("invalid user id")
	}
	return s.productRepo.FindByUserID(userID)
}

// UpdateProduct validates and updates a product.
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	if product.ID == 0 {
		return errors.New("invalid product id")
	}

	if !utils.IsNonEmpty(product.Name) {
		return errors.New("product name is required")
	}

	if !utils.IsPositivePrice(product.Price) {
		return errors.New("product price must be greater than zero")
	}

	if product.UserID != 0 {
		if _, err := s.userRepo.FindByID(product.UserID); err != nil {
			return errors.New("product owner user not found")
		}
	}

	return s.productRepo.Update(product)
}

// DeleteProduct removes a product by ID.
func (s *ProductService) DeleteProduct(id uint) error {
	if id == 0 {
		return errors.New("invalid product id")
	}
	return s.productRepo.Delete(id)
}
