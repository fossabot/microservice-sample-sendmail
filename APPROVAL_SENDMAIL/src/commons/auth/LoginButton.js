import { useAuth0 } from "@auth0/auth0-react";
import React  from "react";
import { Button } from "react-bootstrap";

const LoginButton = () => {
    const { isAuthenticated, loginWithRedirect } = useAuth0();

    return !isAuthenticated ? (
        <Button onClick={() => loginWithRedirect()} >Log In</Button>
    ) : null;
}

export default LoginButton;