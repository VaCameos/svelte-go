import { fileURLToPath, URL } from "url";

import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    host: "localhost",
    port: 8080,
    open: true,
    https: false,
    proxy: {
      "/api": "http://localhost:8081/",
    },
  },
});
