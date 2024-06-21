# Using a Node.js base image as the builder stage here
FROM node:14 as builder

# Setting up working directory
WORKDIR /app

# Copying package.json and package-lock.json (if exists)
COPY package*.json ./

# Installing dependencies
RUN npm install

# Copying the rest of the sample application code
COPY . .

# Build the Next.js application
RUN npm run build

# Production stage
FROM node:14-slim

# Setting working directory
WORKDIR /app

# Copy package.json and package-lock.json (if exists)
COPY package*.json ./

# Install only production dependencies
RUN npm install --only=prod

# Copy built files from the builder stage
COPY --from=builder /app/.next ./.next


# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
