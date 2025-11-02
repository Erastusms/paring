import { Router } from 'express';
import { healthCheck, ping } from './controllers/health.controller';

export const healthRouter = Router();

healthRouter.get('/health', healthCheck);
healthRouter.get('/ping', ping);

// Dummy endpoint untuk trigger event ke RabbitMQ
healthRouter.post('/order', async (req, res) => {
  const order = { id: Date.now(), item: 'Book' };
  res.json({ message: 'Order created event published', order });
});
