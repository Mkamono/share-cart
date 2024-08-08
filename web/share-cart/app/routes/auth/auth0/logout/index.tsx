import {
    ActionFunction,
    type ActionFunctionArgs,
    redirect,
} from "@remix-run/node";
import { config } from "~/service/config.server";

import { destroySession, getSession } from "~/service/session.server";

export const action: ActionFunction = async ({
    request,
}: ActionFunctionArgs) => {
    const session = await getSession(request.headers.get("Cookie"));
    const logoutURL = new URL(config.AUTH0_LOGOUT_URL); // i.e https://YOUR_TENANT.us.auth0.com/v2/logout

    logoutURL.searchParams.set("client_id", config.AUTH0_CLIENT_ID);
    logoutURL.searchParams.set("returnTo", config.AUTH0_RETURN_TO_URL);

    return redirect(logoutURL.toString(), {
        headers: {
            "Set-Cookie": await destroySession(session),
        },
    });
};
