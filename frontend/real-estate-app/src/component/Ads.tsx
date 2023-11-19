import React, {useEffect, useState} from "react";


interface Ads {
    ads: {
        id: number;
        title: string;
        description: string;
    }[]
}

export default function Ads() {
    const [data, setData] = useState<Ads | undefined>();

    useEffect(() => {
        fetch("http://localhost:9000/api/ads").then(async response => {
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
