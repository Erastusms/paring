const Product = require('../models/Product');
const asyncHandler = require('../middleware/asyncHandler');
const ResponseHandler = require('../utils/responseHandler');
const Validators = require('../utils/validators');

class ProductController {
    addProduct = asyncHandler(async (req, res) => {
        const { name, description, price, stock, category, imageUrl, sellerId } = req.body;

        Validators.validatePositiveNumber(price, 'Price');
        const newProduct = new Product({
            name,
            description,
            price,
            stock,
            category,
            imageUrl,
            sellerId
        });

        await newProduct.save();
        return ResponseHandler.created(res, newProduct, 'Product added successfully');
    });

    getProductById = asyncHandler(async (req, res) => {
        const { id } = req.params;
        const product = await Product.findById(id);

        if (!product) {
            return ResponseHandler.notFound(res, 'Product');
        }

        return ResponseHandler.success(res, product, 'Product retrieved successfully');
    });

    listProducts = asyncHandler(async (req, res) => {
        const { category, minPrice, maxPrice, limit = 10, page = 1 } = req.query;
        const query = {};
        if (category) {
            query.category = category;
        }

        if (minPrice || maxPrice) {
            query.price = {};
            if (minPrice) query.price.$gte = Number(minPrice);
            if (maxPrice) query.price.$lte = Number(maxPrice);
        }

        // Pagination
        const skip = (Number(page) - 1) * Number(limit);
        const [products, total] = await Promise.all([
            Product.find(query)
                .skip(skip)
                .limit(Number(limit))
                .sort({ createdAt: -1 }),
            Product.countDocuments(query)
        ]);
        const pagination = {
            total,
            page: Number(page),
            limit: Number(limit),
            totalPages: Math.ceil(total / Number(limit))
        };
        const result = {
            products,
            ...pagination
        }
        return ResponseHandler.success(
            res,
            result,
            'Products retrieved successfully'
        );
    });

    editProduct = asyncHandler(async (req, res) => {
        const { id } = req.params;
        const updates = req.body;

        const allowedUpdates = ['name', 'description', 'price', 'stock', 'category', 'imageUrl'];
        Validators.validateAllowedFields(updates, allowedUpdates);

        if (updates.price) {
            Validators.validatePositiveNumber(updates.price, 'Price');
        }

        const product = await Product.findByIdAndUpdate(
            id,
            updates,
            { new: true, runValidators: true }
        );

        if (!product) {
            return ResponseHandler.notFound(res, 'Product');
        }

        return ResponseHandler.success(res, product, 'Product updated successfully');
    });

    removeProduct = asyncHandler(async (req, res) => {
        const { id } = req.params;
        const product = await Product.findByIdAndDelete(id);

        if (!product) {
            return ResponseHandler.notFound(res, 'Product');
        }

        return ResponseHandler.success(res, product, 'Product removed successfully');
    });

    updateProductStock = asyncHandler(async (req, res) => {
        const { id } = req.params;
        const { stock } = req.body;

        console.log(`[Product Service] Updating stock for ID: ${id} | Delta: ${stock}`);

        const product = await Product.findById(id);

        if (!product) {
            return ResponseHandler.notFound(res, 'Product');
        }

        console.log(`[Product Service] Current stock: ${product.stock}`);
        const newStock = product.stock + Number(stock);
        console.log(`[Product Service] New stock: ${newStock}`);

        if (newStock < 0) {
            console.log(`[Product Service] Stock cannot be negative`);
            return ResponseHandler.badRequest(res, 'Insufficient stock');
        }

        product.stock = newStock;
        await product.save();
        return ResponseHandler.success(res, product, 'Stock updated successfully');
    });
}

module.exports = new ProductController();