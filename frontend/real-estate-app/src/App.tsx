import React, {useEffect, useState} from 'react';
import './App.css';
import { UserManager } from 'oidc-client-ts';


interface Ads {
  ads: {
    id: number;
    title: string;
  }[]
}

function App() {
  const userManager = new UserManager({
    authority: "http://localhost:8080",
    client_id: "",
    client_secret: "",
    redirect_uri: "/callback",
  });

  useEffect(() => {
    userManager.getUser().then(user => {
      console.log(user);
    })
  }, [])

  const [data, setData] = useState<Ads | undefined>();

  useEffect(() => {
    fetch("http://localhost:9000/api/ads").then(async response => {
      const results = await response.json();
      setData(results);
    })
  }, []);

  return (
    <div className="App">
      {data !== undefined ?
          <ul>
            {data.ads.map(x => {
              return <li key={x.id}>{x.title}</li>
            })}
          </ul>
          : ""}
    </div>
  );
}

export default App;
