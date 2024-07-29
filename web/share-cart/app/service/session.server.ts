import { createCookieSessionStorage } from "@remix-run/node";
import { config } from "./config.server";

export const sessionStorage = createCookieSessionStorage({
	cookie: {
		name: "_session", // use any name you want here
		sameSite: "lax", // this helps with CSRF
		path: "/", // remember to add this so the cookie will work in all routes
		httpOnly: true, // for security reasons, make this cookie http only
		secrets: [config.SESSION_SECRET], // replace this with an actual secret
		secure: config.isProduction, // enable this in prod only
		maxAge: 60 * 60 * 24 * 30, // 30 days
	},
});

export const { getSession, commitSession, destroySession } = sessionStorage;
