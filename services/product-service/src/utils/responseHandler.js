class ResponseHandler {
    static success(res, data = null, message = 'Success', statusCode = 200) {
        const response = {
            success: true,
            message,
            ...(data && { data })
        };
        return res.status(statusCode).json(response);
    }

    static error(res, message = 'An error occurred', statusCode = 500, errors = null) {
        const response = {
            success: false,
            message,
            ...(errors && { errors })
        };
        return res.status(statusCode).json(response);
    }

    static notFound(res, message = 'Resource') {
        return this.error(res, `${message} not found`, 404);
    }

    static badRequest(res, message = 'Bad request', errors = null) {
        return this.error(res, message, 400, errors);
    }

    static created(res, data = null, message = 'Resource created successfully') {
        return this.success(res, data, message, 201);
    }
}

module.exports = ResponseHandler;