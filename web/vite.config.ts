import { defineConfig } from "vite";
import path from "path";

import vue from "@vitejs/plugin-vue";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import webfontDownload from 'vite-plugin-webfont-dl';

export default defineConfig({
  base: "/",
  server: {
    proxy: {
      "/api": {
        target: "http://0.0.0.0:8088",
      },
      "/_/api": {
        target: "http://0.0.0.0:8088",
      },
    },
  },
  plugins: [
    vue(),
    createSvgIconsPlugin({
      iconDirs: [path.resolve(process.cwd(), "./src/assets/icons")],
      symbolId: "icon-[dir]-[name]",
    }),
    webfontDownload(),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
    extensions: [".js", ".ts", ".json", ".vue"],
  },
});
