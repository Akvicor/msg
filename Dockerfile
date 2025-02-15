# 构建应用
FROM akvicor/builder:v0.0.9-node20-go22 AS builder
WORKDIR /app
COPY . .
RUN cd frontend && make build
RUN cp -r frontend/build backend/cmd/app/server/web/build
RUN cd backend && make build

# 最小化镜像
FROM debian:12.8-slim
WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/backend/build/msg ./msg
COPY --from=builder /app/prod.sh ./prod.sh

RUN ln -sf /usr/share/zoneinfo/Etc/GMT-8 /etc/localtime && \
    mkdir /data && \
    chmod +x ./prod.sh

EXPOSE 3000
CMD ["./prod.sh"]

