# Use a Node.js base image as the builder stage
FROM node:14 as builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if exists)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Next.js application
RUN npm run build

# Production stage
FROM node:14-slim

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if exists)
COPY package*.json ./

# Install only production dependencies
RUN npm install --only=prod

# Copy built files from the builder stage
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/../public ./public

# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
