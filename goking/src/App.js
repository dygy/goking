import './App.css';
import { hot } from 'react-hot-loader/root';
function App() {
    try {
        // console.log(window.getOS());
    } catch (e) {
        //
    }
  return (
    <div className="App">
        2222
      <header className="App-header">
        <div className={'btn-row'}>
            <button onClick={()=>window.location?.reload()} className={'btn'}>sk,/SASDSDdss ls</button>
            <button className={'btn'}>{Math.random()*10}</button>
            <button className={'btn'}>{Math.random()*10}</button>
        </div>
      </header>
    </div>
  );
}

export default hot(App);
