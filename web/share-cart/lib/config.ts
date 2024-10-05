import { cleanEnv, str, url } from "envalid"; // cspell:disable-line

export const config = cleanEnv(process.env, {
	API_HOST: url(),
	AUTH0_SECRET: str(),
	AUTH0_BASE_URL: url(),
	AUTH0_ISSUER_BASE_URL: url(),
	AUTH0_CLIENT_ID: str(),
	AUTH0_CLIENT_SECRET: str(),
	NODE_ENV: str({ choices: ["development", "production", "staging", "test"] }),
});
