import express from 'express';
import { checkBackendHealth } from './healthz.js';

const app = express();
const port = process.env.PORT || 3000;

app.get('/healthz', async (req, res) => {
  const healthy = await checkBackendHealth();
  if (healthy) {
    res.status(200).json({ status: 'ok' });
  } else {
    res.status(500).json({ status: 'backend not ready' });
  }
});

app.listen(port, () => {
  console.log(`Health server listening on port ${port}`);
});
