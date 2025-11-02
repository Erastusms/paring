exports.confirmPayment = (req, res) => {
    const { orderId, amount } = req.body;
    if (!orderId || !amount) {
        return res.status(400).json({ error: 'Missing orderId or amount' });
    }

    // Simulate payment: 70% success rate
    const isSuccess = Math.random() > 0.3;
    const transactionId = `txn_${Date.now()}`;

    if (isSuccess) {
        res.status(200).json({
            status: 'SUCCESS',
            transactionId,
            message: 'Payment confirmed'
        });
    } else {
        res.status(400).json({
            status: 'FAILED',
            message: 'Payment failed (simulated)'
        });
    }
};