import { Request, Response } from 'express';

export const healthCheck = (req: Request, res: Response) => {
  res.json({ status: 'ok', service: 'auth-service' });
};

export const ping = (req: Request, res: Response) => {
  res.send('pong');
};
