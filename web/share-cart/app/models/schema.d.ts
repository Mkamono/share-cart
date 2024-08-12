/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
	"/test": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		get: {
			parameters: {
				query?: never;
				header?: never;
				path?: never;
				cookie?: never;
			};
			requestBody?: never;
			responses: {
				200: components["responses"]["200OK"];
				default: components["responses"]["500InternalServerError"];
			};
		};
		put?: never;
		post?: never;
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/sign-up": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		get?: never;
		put?: never;
		post: {
			parameters: {
				query?: never;
				header?: never;
				path?: never;
				cookie?: never;
			};
			requestBody: components["requestBodies"]["SignUp"];
			responses: {
				201: components["responses"]["201Created"];
				401: components["responses"]["401Unauthorized"];
				default: components["responses"]["500InternalServerError"];
			};
		};
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/market": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		/** Get all markets */
		get: {
			parameters: {
				query?: never;
				header?: never;
				path?: never;
				cookie?: never;
			};
			requestBody?: never;
			responses: {
				200: components["responses"]["getMarket"];
				default: components["responses"]["500InternalServerError"];
			};
		};
		put?: never;
		post?: never;
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
}
export type webhooks = Record<string, never>;
export interface components {
	schemas: {
		market: {
			/** @example 1 */
			id: number;
			/** @example market_name */
			name: string;
			/** @example market_description */
			description: string;
		};
	};
	responses: {
		/** @description Internal Server Error */
		"500InternalServerError": {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": {
					/** @example Internal Server Error */
					message: string;
				};
			};
		};
		/** @description OK */
		"200OK": {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": {
					/** @example OK */
					message: string;
				};
			};
		};
		/** @description Created */
		"201Created": {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": {
					/** @example Created */
					message: string;
				};
			};
		};
		/** @description Unauthorized */
		"401Unauthorized": {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": {
					/** @example Unauthorized */
					message: string;
				};
			};
		};
		/** @description A list of markets. */
		getMarket: {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": components["schemas"]["market"][];
			};
		};
	};
	parameters: never;
	requestBodies: {
		SignUp: {
			content: {
				"application/json": {
					/** @example your_name */
					name?: string;
				};
			};
		};
	};
	headers: never;
	pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
