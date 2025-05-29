FROM node:22-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:alpine

# ✅ Serve from Nginx's default root directory
COPY --from=builder /app/dist /usr/share/nginx/html

# ✅ Optional: Use custom Nginx config (for React Router support)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Optional: inject env vars if needed
COPY env.sh /docker-entrypoint.d/env.sh
RUN dos2unix /docker-entrypoint.d/env.sh && chmod +x /docker-entrypoint.d/env.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]