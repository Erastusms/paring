# Gunakan Node.js versi LTS yang ringan
FROM node:18-alpine AS builder

# Tentukan working directory di dalam container
WORKDIR /app

# Copy package.json & package-lock.json (kalau ada)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy seluruh source code
COPY . .

# Build TypeScript → JavaScript
RUN npm run build

# =========================
# Stage 2: Production image
# =========================
FROM node:18-alpine

WORKDIR /app

# Copy hanya hasil build dan node_modules (lebih ringan)
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY package*.json ./

# Expose port default
EXPOSE 3000

# Jalankan service
CMD ["node", "dist/index.js"]
