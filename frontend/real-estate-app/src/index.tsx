import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {createBrowserRouter, redirect, RouterProvider} from "react-router-dom";
import WelcomeUser from "./component/WelcomeUser";
import Ads from "./component/Ads";
import Unauthenticated from "./component/Unauthenticated";
import {InMemoryWebStorage, UserManager} from "oidc-client-ts";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const store = new InMemoryWebStorage();

const userManager = new UserManager({
    authority: "http://localhost:8080/realms/Real-Estate/",
    client_id: "webapp",
    client_secret: "2NCzklvScW6er0PVoN6zhqah9jffBzHg",
    redirect_uri: "http://localhost:3000/authenticated",
    response_mode: "query",
    response_type: "code",
    // stateStore: new InMemoryWebStorage(),
    // @ts-ignore
    // stateStore: {
    //     get(key: string): Promise<string | null> {
    //         return Promise.resolve(store.getItem(key))
    //     },
    //     set(key: string,  value: string): Promise<void> {
    //         store.setItem(key,  value);
    //         return Promise.resolve()
    //     },
    //     remove(key: string): Promise<string | null> {
    //         store.removeItem(key)
    //         return Promise.resolve(null);
    //     },
    //     getAllKeys(): Promise<string[]> {
    //         throw new Error("Not implemented")
    //     },
    // },
});

async function loadUser() {
    const user = await userManager.getUser();
    if (user === null) {
        redirect("/");
    }
    return user;
}

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        children: [
            {
                path: "/",
                element: <Unauthenticated userManager={userManager} />,
            },
            {
                path: "authenticated",
                element: <WelcomeUser userManager={userManager} />,
                // loader: loadUser,
                children: [
                    {
                        path: "ads",
                        element: <Ads />,
                        loader: loadUser,
                    }
                ]
            },
        ]
    },
]);

root.render(
  // <React.StrictMode>
    <RouterProvider router={router}/>
  // </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
