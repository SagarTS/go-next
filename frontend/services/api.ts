import axios from 'axios';

axios.defaults.responseType = 'json';

const API = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export { API };
