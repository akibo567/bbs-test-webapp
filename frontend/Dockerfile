ARG NODE_IMAGE=node:20-alpine

FROM ${NODE_IMAGE} AS base
WORKDIR /app
COPY package*.json ./
RUN --mount=type=cache,target=/root/.npm npm ci

FROM base AS dev
ENV NODE_ENV=development
EXPOSE 3000
CMD ["npm","run","dev"]

FROM base AS builder
COPY . .
RUN npm run build  # Nextなら output: 'standalone'


FROM ${NODE_IMAGE} AS runner
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
USER node
EXPOSE 3000
CMD ["node","server.js"]