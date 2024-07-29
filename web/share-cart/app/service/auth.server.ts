import { sessionStorage } from "./session.server";

import { Authenticator } from "remix-auth";
import {
	type Auth0ExtraParams,
	type Auth0Profile,
	Auth0Strategy,
} from "remix-auth-auth0";
import { config } from "./config.server";

export const authenticator = new Authenticator<{
	accessToken: string;
	refreshToken: string | undefined;
	extraParams: Auth0ExtraParams;
	profile: Auth0Profile;
}>(sessionStorage);

const auth0Strategy = new Auth0Strategy(
	{
		callbackURL: config.AUTH0_CALLBACK_URL,
		clientID: config.AUTH0_CLIENT_ID,
		clientSecret: config.AUTH0_CLIENT_SECRET,
		domain: config.AUTH0_DOMAIN,
		audience: config.AUTH0_AUDIENCE,
	},
	async ({ accessToken, refreshToken, extraParams, profile }) => {
		return {
			accessToken,
			refreshToken,
			extraParams,
			profile,
		};
	},
);

authenticator.use(auth0Strategy);
