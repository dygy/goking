import './App.css';
import { hot } from 'react-hot-loader/root';
function App() {
    try {
        console.log(window?.getOS());
    } catch (e) {
        //
    }
  return (
    <div className="App">
      <header className="App-header">
        <div className={'btn-row'}>
            1
            <button className={'btn'}>{Math.random()*10}</button>
            <button className={'btn'}>{Math.random()*10}</button>
            <button className={'btn'}>{Math.random()*10}</button>
        </div>
      </header>
    </div>
  );
}

export default hot(App);
