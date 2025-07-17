# frontend/Dockerfile
FROM docker.io/node:24-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 3000
EXPOSE 3001

CMD ["sh", "-c", "npm run dev & npm run healthz"]
