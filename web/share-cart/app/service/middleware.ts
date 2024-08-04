import { Middleware } from "openapi-fetch";

export const jwtMiddleware = (accessToken: string): Middleware => ({
  async onRequest({ request }) {
    request.headers.set("Authorization", accessToken);
    return request;
  },
});
