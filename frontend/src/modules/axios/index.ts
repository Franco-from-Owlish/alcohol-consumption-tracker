import axios, { type CreateAxiosDefaults } from 'axios'

const axiosConfig: CreateAxiosDefaults = {
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 1000,
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  }
};

const api = axios.create(axiosConfig);

export {
  api
}