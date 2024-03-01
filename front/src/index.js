import React from 'react';
import App from './App';
import ReactDOM from 'react-dom';
import { BrowserRouter } from "react-router-dom";
import "./static/css/styleIndex.css"
ReactDOM.render(
    <React.StrictMode>
        <BrowserRouter>
        <App />
        </BrowserRouter>,
    </React.StrictMode>,
  document.getElementById('root')
);