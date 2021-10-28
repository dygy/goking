import './App.css';
import { hot } from 'react-hot-loader/root';
import {useState} from "react";
function App() {
    const [resp, setResp] = useState(null)
    window.getStr()?.then(setResp)
    return (
        <div className="App">
            <p style={{
                color: 'white',
                fontSize: '11px',
                width: '100vw',
                height: '300px',
                position: 'absolute',
                top: 15,
                left: 15,
                display: 'flex',
                flexWrap: 'wrap',
            }}>
                {Object.keys(window)
                    .map((e) => <span style={{margin: 3}}>{e}</span>)
                }
            </p>
            <header className="App-header">
                <div className={'btn-row'}>
                    {resp && <button className={'btn'}>{resp}</button>}
                    <button className={'btn'}>{Math.random()*10}</button>
                    <button className={'btn'}>{Math.random()*10}</button>
                    <button className={'btn'}>{Math.random()*10}</button>
                </div>
            </header>
        </div>
  );
}

export default hot(App);
