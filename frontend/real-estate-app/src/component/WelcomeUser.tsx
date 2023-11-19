import {User, UserManager} from "oidc-client-ts";
import {useEffect, useState} from "react";


export default function WelcomeUser(props: { userManager: UserManager }) {
    const {userManager} = props;

    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        userManager.getUser().then(userInfo => {
            setUser(userInfo as User)
            console.log(userInfo);
        }).catch(err => {
            console.error(err)
        })
    }, [userManager])

    return (
        <div>
            Hello {user?.profile.preferred_username ?? "Loading"}
            <br/>
            <a href="/ads">Go see the ads</a>
        </div>
    )
}
