const mongoose = require("mongoose");
require("dotenv").config();

const connectDB = async () => {
  try {
    const mongoUri = process.env.MONGO_URI || 'mongodb://admin:pass@localhost:27018/paring_product?authSource=admin';
    await mongoose.connect(mongoUri);
    // await mongoose.connect('mongodb://admin:pass@localhost:27018/paring_product?authSource=admin');
    console.log("MongoDB connected");
  } catch (err) {
    console.error(err.message);
    process.exit(1);
  }
};

module.exports = connectDB;
