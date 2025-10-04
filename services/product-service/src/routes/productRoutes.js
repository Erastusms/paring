const express = require('express');
const { addProduct, listProducts } = require('../controllers/productController');
// const authenticate = require('../middleware/authenticate');  // Uncomment jika butuh auth

const router = express.Router();

router.post('/', addProduct);
// router.get('/', authenticate, listProducts);  // Dengan auth
router.get('/', listProducts);

module.exports = router;