const express = require('express');
const { addProduct, listProducts, editProduct, removeProduct } = require('../controllers/productController');
const authenticate = require('../middleware/authenticate');

const router = express.Router();

router.post('/', authenticate, addProduct);
router.get('/', listProducts);
router.put('/:id', authenticate, editProduct);
router.delete('/:id', authenticate, removeProduct);

module.exports = router;