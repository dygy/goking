import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const render = (Component) => {
    ReactDOM.render(
        <React.StrictMode>
            <Component />
        </React.StrictMode>,
        document.getElementById('root')
    );
}
if (module.hot) {
    module.hot.accept('./App.js', () => {
        console.log('Accepting the updated printMe module!');
        render(App);
    })
}
render(App)
// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
window.forceClearCache = () => {
    const scripts = document.getElementsByTagName('script');
    const torefreshs = ['myscript.js', 'myscript2.js']; // list of js to be refresh
    const key = 1; // change this key every time you want force a refresh
    for(let i = 0; i < scripts.length; i++){
        for(let j = 0; j < torefreshs.length; j++){
            if(scripts[i].src && (scripts[i].src.indexOf(torefreshs[j]) > -1)){
                scripts[i].src = scripts[i].src.replace(torefreshs[j], `${torefreshs[j]}k=${key}`); // change src in order to refresh js
            }
        }
    }
    window.location.reload(true);
}