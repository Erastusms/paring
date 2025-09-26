import express from 'express';
import pool from "./db";
import { healthRouter } from './routes';
// import { startConsumer } from './events/consumer';

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());
app.use('/', healthRouter);

app.get("/users", async (_req, res) => {
  try {
    const result = await pool.query("SELECT userid, name, email, createdat FROM users");
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).send("Database error");
  }
});

app.listen(PORT, () => {
  console.log(`✅ Auth Service running on port ${PORT}`);
  // startConsumer(); // start RabbitMQ consumer
});
