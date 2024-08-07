import {
	ActionFunction,
	type ActionFunctionArgs,
	redirect,
} from "@remix-run/node";
import { authenticator } from "~/service/auth.server";

export const loader = () => redirect("/");

export const action: ActionFunction = ({ request }: ActionFunctionArgs) => {
	return authenticator.authenticate("auth0", request);
};
