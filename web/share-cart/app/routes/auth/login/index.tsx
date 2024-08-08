import { Form } from "@remix-run/react";
import { $path } from "remix-routes";

export default function Login() {
    return (
        <Form action={$path("/auth/auth0/login")} method="post">
            <button type="submit">Login with Auth0</button>
        </Form>
    );
}
