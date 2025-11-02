const express = require('express');
const router = express.Router();
const ProductController = require('../controllers/productController');
const authenticate = require('../middleware/authenticate');

router.post('/', authenticate, ProductController.addProduct);
router.get('/:id', ProductController.getProductById);
router.get('/', ProductController.listProducts);
router.put('/:id', authenticate, ProductController.editProduct);
router.delete('/:id', authenticate, ProductController.removeProduct);
router.patch('/:id', authenticate, ProductController.updateProductStock);

module.exports = router;