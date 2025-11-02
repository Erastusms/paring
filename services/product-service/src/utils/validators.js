class Validators {
    static validatePositiveNumber(value, fieldName = 'Value') {
        if (value <= 0) {
            throw new Error(`${fieldName} must be positive`);
        }
        return true;
    }

    static validateAllowedFields(updates, allowedFields) {
        const updateKeys = Object.keys(updates);
        const isValid = updateKeys.every(key => allowedFields.includes(key));
        if (!isValid) {
            const invalidFields = updateKeys.filter(key => !allowedFields.includes(key));
            throw new Error(`Invalid fields: ${invalidFields.join(', ')}`);
        }
        return true;
    }

    static validateRequiredFields(data, requiredFields) {
        const missingFields = requiredFields.filter(field => !data[field]);
        if (missingFields.length > 0) {
            throw new Error(`Missing required fields: ${missingFields.join(', ')}`);
        }
        return true;
    }
}

module.exports = Validators;
