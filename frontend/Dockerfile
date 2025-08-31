# Build stage
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
# Next.js 13+ / 15系なら standalone 出力を推奨
# next.config.js で output: 'standalone' にしておく
RUN npm run build

# Run stage
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
# standalone で最小ファイル群だけコピー
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
# Next が期待するPORT
ENV PORT=3000
EXPOSE 3000
CMD ["node", "server.js"]
