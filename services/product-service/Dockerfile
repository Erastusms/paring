FROM node:18

# Tentukan working directory
WORKDIR /app

# Copy file package.json dan package-lock.json (jika ada)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy seluruh source code
COPY . .

# Expose port server kamu
EXPOSE 3001

# Jalankan app
CMD ["npm", "start"]
