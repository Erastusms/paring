import { Redis } from 'ioredis';

export default class EventPublisher {
  private redis: Redis;

  constructor() {
    this.redis = new Redis({
      host: process.env.REDIS_HOST || 'localhost',
      port: Number(process.env.REDIS_PORT) || 6379,
      retryStrategy: (times) => Math.min(times * 50, 2000), // Retry setiap 50ms, max 2s
    });
    
    // Error handling sederhana untuk koneksi
    this.redis.on('error', (err) => {
      console.error('Redis Publisher Error:', err);
    });
  }

  async publish(channel: string, message: object): Promise<void> {
    const payload = JSON.stringify(message);
    await this.redis.publish(channel, payload);
    console.log(`Published event to channel "${channel}":`, payload);
  }
}
