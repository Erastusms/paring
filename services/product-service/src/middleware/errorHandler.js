const ResponseHandler = require('../utils/responseHandler');

const errorHandler = (err, req, res, next) => {
    console.error('Error:', err);

    // Mongoose validation error
    if (err.name === 'ValidationError') {
        const errors = Object.values(err.errors).map(e => e.message);
        return ResponseHandler.badRequest(res, 'Validation failed', errors);
    }

    // Mongoose CastError (invalid ObjectId)
    if (err.name === 'CastError') {
        return ResponseHandler.badRequest(res, 'Invalid ID format');
    }

    // Duplicate key error
    if (err.code === 11000) {
        const field = Object.keys(err.keyPattern)[0];
        return ResponseHandler.badRequest(res, `${field} already exists`);
    }

    // Default error
    return ResponseHandler.error(res, err.message || 'Internal server error', 500);
};

module.exports = errorHandler;