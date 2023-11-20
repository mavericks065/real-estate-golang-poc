import React, {useEffect, useState} from "react";
import {User} from "oidc-client-ts";
import {useLoaderData} from "react-router-dom";


interface Ads {
    ads: {
        id: number;
        title: string;
        description: string;
    }[]
}

export default function Ads() {
    const [data, setData] = useState<Ads | undefined>();
    const user = useLoaderData() as User;
    console.log("Loaded data", user);

    useEffect(() => {
        fetch("http://localhost:9000/api/ads", {
            headers: {
                Authorization: user.access_token,
            }
        }).then(async response => {
            const results = await response.json();
            setData(results);
        })
    }, []);

    return (
        <div>
            {data !== undefined ?
                <ul>
                    {data.ads.map(x => {
                        return <li key={x.id}>{x.title}</li>
                    })}
                </ul>
                : ""}
        </div>
    )
}
