import { cleanEnv, str, url } from "envalid"; // cspell:disable-line

export const config = cleanEnv(process.env, {
	SESSION_SECRET: str(),
	AUTH0_CALLBACK_URL: url(),
	AUTH0_CLIENT_ID: str(),
	AUTH0_CLIENT_SECRET: str(),
	AUTH0_DOMAIN: str(),
	AUTH0_LOGOUT_URL: url(),
	AUTH0_RETURN_TO_URL: url(),
	AUTH0_AUDIENCE: str(),
	NODE_ENV: str({ choices: ["development", "production", "staging", "test"] }),
});
