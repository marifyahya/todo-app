FROM node:20-alpine

WORKDIR /app

# Copy package.json and package-lock.json (if exists)
COPY package*.json ./

RUN npm install

# Copy project files
COPY . .

# Expose port
EXPOSE 5173

# Run in dev mode with host flag for Docker access
CMD ["npm", "run", "dev", "--", "--host"]
