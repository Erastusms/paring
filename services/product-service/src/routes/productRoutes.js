const express = require('express');
const { addProduct, getProductById, listProducts, editProduct, removeProduct, updateProductStock } = require('../controllers/productController');
const authenticate = require('../middleware/authenticate');

const router = express.Router();

// router.post('/', authenticate, addProduct);
router.post('/', addProduct);
router.get('/:id', getProductById);
router.get('/', listProducts);
router.put('/:id', authenticate, editProduct);
router.delete('/:id', authenticate, removeProduct);
router.patch('/:id', authenticate, updateProductStock);

module.exports = router;