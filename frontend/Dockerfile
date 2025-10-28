ARG NODE_IMAGE=node:25.0.0-trixie-slim

FROM ${NODE_IMAGE} AS base
WORKDIR /app
COPY package.json package-lock.json ./
RUN --mount=type=cache,target=/root/.npm npm ci

FROM base AS dev
ENV NODE_ENV=development
EXPOSE 3000
CMD ["npm","run","dev"]

FROM base AS builder
COPY . .

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates git curl && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN npm run build  # Nextなら output: 'standalone'


FROM ${NODE_IMAGE} AS prod
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
USER node
EXPOSE 3000
ENV PORT=3000 HOSTNAME=0.0.0.0
CMD ["node","server.js"]