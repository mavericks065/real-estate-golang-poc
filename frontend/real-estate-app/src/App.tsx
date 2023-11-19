import React, {useEffect, useState} from 'react';
import './App.css';


interface Ads {
  ads: {
    id: number;
    title: string;
  }[]
}

function App() {
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
