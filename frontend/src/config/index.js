const config = {
  apiUrl: import.meta.env.VITE_API_URL || "http://localhost:8081/api",
  appName: "Todo App",
  env: import.meta.env.MODE, // 'development' or 'production'
};

export default config;
