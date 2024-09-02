import {
	ActionFunction,
	type ActionFunctionArgs,
	redirect,
} from "@remix-run/node";
import { $path } from "remix-routes";
import { authenticator } from "~/service/auth.server";

export const loader = () => redirect($path("/"));

export const action: ActionFunction = ({ request }: ActionFunctionArgs) => {
	return authenticator.authenticate("auth0", request);
};
