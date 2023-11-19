import React, {useEffect, useState} from 'react';
import logo from './logo.svg';
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
    fetch("http://localhost:8080/api/ads").then(async response => {
      const results = await response.json();
      setData(results);
    })
  }, []);


  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
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
