const Product = require('../models/Product');

exports.addProduct = async (req, res) => {
    try {
        const { name, description, price, stock, category, imageUrl, sellerId } = req.body;

        if (price <= 0) {
            return res.status(400).json({ error: 'Price must be positive' });
        }

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
        res.status(201).json({ message: 'Product added successfully', product: newProduct });
    } catch (error) {
        console.error('Error adding product:', error);
        res.status(500).json({ error: 'Server error while adding product' });
    }
};

exports.listProducts = async (req, res) => {
    try {
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

        // Pagination: skip = (page-1) * limit
        const skip = (Number(page) - 1) * Number(limit);

        const products = await Product.find(query)
            .skip(skip)
            .limit(Number(limit))
            .sort({ createdAt: -1 });  // Sort by newest first

        const total = await Product.countDocuments(query);  // Untuk metadata pagination

        res.status(200).json({
            message: 'Products listed successfully',
            products,
            pagination: {
                total,
                page: Number(page),
                limit: Number(limit),
                totalPages: Math.ceil(total / limit)
            }
        });
    } catch (error) {
        console.error('Error listing products:', error);
        res.status(500).json({ error: 'Server error while listing products' });
    }
};