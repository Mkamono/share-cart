import type { LoaderFunctionArgs } from "@remix-run/node";
import { $path } from "remix-routes";

import { authenticator } from "~/service/auth.server";

export const loader = ({ request }: LoaderFunctionArgs) => {
    return authenticator.authenticate("auth0", request, {
        successRedirect: $path("/auth/user/register"),
        failureRedirect: $path("/auth/login"),
    });
};
