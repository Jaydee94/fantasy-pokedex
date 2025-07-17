import axios from 'axios';
import { createServer } from 'http';

const port = process.env.HEALTHZ_PORT || 3001;
const backendUrl = process.env.VITE_API_URL || 'http://localhost:8080';

const server = createServer(async (req, res) => {
  if (req.url === '/healthz') {
    try {
      const response = await axios.get(`${backendUrl}/healthz`);
      if (response.status === 200) {
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ status: 'ok' }));
      } else {
        res.writeHead(500, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ status: 'backend not ready' }));
      }
    } catch (err) {
      res.writeHead(500, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'backend unreachable', error: err.message }));
    }
  } else {
    res.writeHead(404);
    res.end();
  }
});

server.listen(port, () => {
  console.log(`Vite healthz server listening on port ${port}`);
});
