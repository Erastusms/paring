import { Redis } from 'ioredis';

export default class EventSubscriber {
  private redis: Redis;

  constructor() {
    this.redis = new Redis({
      host: process.env.REDIS_HOST || 'localhost',
      port: Number(process.env.REDIS_PORT) || 6379,
      retryStrategy: (times) => Math.min(times * 50, 2000), // Retry setiap 50ms, max 2s
    });

    // Error handling sederhana untuk koneksi
    this.redis.on('error', (err) => {
      console.error('Redis Subscriber Error:', err);
    });
  }

  async subscribe<T>(
    channel: string,
    handler: (msg: T) => void
  ): Promise<void> {
    await this.redis.subscribe(channel);
    this.redis.on('message', (ch, message) => {
      if (ch === channel) {
        try {
          const parsed = JSON.parse(message) as T;
          console.log(`Received event from "${channel}":`, message);
          handler(parsed);
        } catch (err) {
          console.error('Failed to parse message:', err);
        }
      }
    });
  }
}
