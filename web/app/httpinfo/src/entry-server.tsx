// @refresh reload
import { createHandler, StartServer } from "@solidjs/start/server";

import { setServerData } from "./store/ServerStore";

export default createHandler((pageEvent) => {
  if (new URL(pageEvent.request.url).pathname.startsWith("/.well-known/")) {
    return new Response("", { status: 404 });
  }
  setServerData({ headers: Array.from(pageEvent.request.headers.entries()) });
    return (
      <StartServer
        document={({ assets, children, scripts }) => (
          <html lang="en">
            <head>
              <meta charset="utf-8" />
              <meta name="viewport" content="width=device-width, initial-scale=1" />
              <link rel="icon" href="/favicon.svg" />
              {assets}
            </head>
            <body>
              <div id="app">{children}</div>
              {scripts}
            </body>
          </html>
        )}
      />
  )
});
