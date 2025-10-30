const express = require('express');
const dotenv = require('dotenv');
const { confirmPayment } = require('./paymentController');

dotenv.config();

const app = express();
app.use(express.json());

app.post('/api/payment/confirm', confirmPayment);

const PORT = process.env.PORT || 3002;
app.listen(PORT, () => console.log(`Mock Payment Service running on port ${PORT}`));