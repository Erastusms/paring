FROM node:18

# Tentukan working directory di dalam container
WORKDIR /app

# Copy file package.json dan package-lock.json (kalau ada)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy semua source code ke container
COPY . .

# Expose port sesuai server kamu
EXPOSE 3001

# Jalankan perintah untuk start app
CMD ["npm", "start"]
