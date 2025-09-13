import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080/api", // Go backend
});

// Products
export const getProducts = () => API.get("/products");
export const addProduct = (data) => API.post("/products", data);

// Stock
export const stockIn = (data) => API.post("/stock/in", data);
export const stockOut = (data) => API.post("/stock/out", data);

// Inventory
export const getInventory = () => API.get("/inventory");