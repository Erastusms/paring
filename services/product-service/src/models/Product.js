const mongoose = require('mongoose');

const productSchema = new mongoose.Schema({
    name: {
        type: String,
        required: [true, 'Product name is required'],
        trim: true,
        maxlength: [100, 'Name cannot exceed 100 characters']
    },
    description: {
        type: String,
        required: [true, 'Description is required'],
        trim: true
    },
    price: {
        type: Number,
        required: [true, 'Price is required'],
        min: [0, 'Price cannot be negative']
    },
    stock: {
        type: Number,
        required: [true, 'Stock is required'],
        min: [0, 'Stock cannot be negative']
    },
    category: {
        type: String,
        required: [true, 'Category is required'],
        enum: ['electronics', 'clothing', 'books', 'food', 'other']
    },
    imageUrl: {
        type: String,
        default: null
    },
    sellerId: {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'User',  // Reference ke User Service nanti; untuk sekarang, optional
        default: null
    }
}, {
    timestamps: true
});

module.exports = mongoose.model('Product', productSchema);