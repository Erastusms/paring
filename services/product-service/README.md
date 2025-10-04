## Product Service Endpoints
- POST /api/products: Add new product
  - Body: { name, description, price, stock, category, imageUrl, sellerId }
  - Response: 201 { message, product }
- GET /api/products: List products with filtering
  - Query Params: category (string), minPrice (number), maxPrice (number), page (number, default 1), limit (number, default 10)
  - Response: 200 { message, products, pagination }