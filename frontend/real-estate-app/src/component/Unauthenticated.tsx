import React, {useEffect, useState} from "react";
import {UserManager} from "oidc-client-ts";



export default function Unauthenticated(props: { userManager: UserManager }) {
    const { userManager } = props;
    const login = () => {
        userManager.signinRedirect()
    }

    return <button onClick={login} type="button">
        login
    </button>
}
