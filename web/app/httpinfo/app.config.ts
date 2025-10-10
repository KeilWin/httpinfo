import { defineConfig } from "@solidjs/start/config";

export default defineConfig({
    vite: {
        css: {
            modules: {
            localsConvention: "camelCaseOnly",
            },
        },
    },
});
