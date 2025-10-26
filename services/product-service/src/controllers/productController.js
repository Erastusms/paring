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
        res.status(500).json({ error: error.message });
    }
};

exports.getProductById = async (req, res) => {
    try {
        const { id } = req.params;
        const product = await Product.findById(id);
        if (!product) {
            return res.status(404).json({ error: 'Product not found' });
        }
        res.status(200).json({ product });
    } catch (error) {
        console.error('Error fetching product:', error);
        res.status(500).json({ error: 'Server error while fetching product' });
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

exports.editProduct = async (req, res) => {
    try {
        const { id } = req.params;
        const updates = req.body;

        const allowedUpdates = ['name', 'description', 'price', 'stock', 'category', 'imageUrl'];
        const isValidOperation = Object.keys(updates).every((update) => allowedUpdates.includes(update));
        if (!isValidOperation) {
            return res.status(400).json({ error: 'Invalid updates' });
        }

        const product = await Product.findByIdAndUpdate(id, updates, { new: true, runValidators: true });
        if (!product) {
            return res.status(404).json({ error: 'Product not found' });
        }

        res.status(200).json({ message: 'Product updated successfully', product });
    } catch (error) {
        console.error('Error editing product:', error);
        res.status(500).json({ error: 'Server error while editing product' });
    }
};

exports.removeProduct = async (req, res) => {
    try {
        const { id } = req.params;

        const product = await Product.findByIdAndDelete(id);
        if (!product) {
            return res.status(404).json({ error: 'Product not found' });
        }

        res.status(200).json({ message: 'Product removed successfully' });
    } catch (error) {
        console.error('Error removing product:', error);
        res.status(500).json({ error: 'Server error while removing product' });
    }
};

exports.updateProductStock = async (req, res) => {
    try {
        const { id } = req.params;
        const { stock } = req.body; // 'stock' ini adalah delta, misal -1

        console.log(`[Product Service] Received update for ID: ${id} | Delta: ${stock}`);

        const product = await Product.findById(id);
        if (!product) {
            console.log(`[Product Service] FAILED: Product not found (ID: ${id})`);
            return res.status(404).json({ error: 'Product not found' });
        }

        console.log(`[Product Service] Found product. Old stock: ${product.stock}`);

        // Logika ini sudah benar
        product.stock += stock;

        console.log(`[Product Service] New stock calculated: ${product.stock}`);

        if (product.stock < 0) {
            console.log(`[Product Service] FAILED: Stock cannot be negative`);
            return res.status(400).json({ error: 'Stock cannot be negative' });
        }

        await product.save(); // Kemungkinan besar gagal di sini
        console.log(`[Product Service] SUCCESS: Stock saved for ID: ${id}`);

        res.status(200).json({ message: 'Stock updated', product });
    } catch (error) {
        // INI YANG KEMUNGKINAN BESAR TERJADI
        console.error('[Product Service] FATAL ERROR:', error.message);
        res.status(500).json({ error: 'Server error', message: error.message });
    }
};