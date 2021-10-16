import './App.css';
import { hot } from 'react-hot-loader/root';
function App() {
    try {
        console.log(window.getOS());
    } catch (e) {
        //
    }
  return (
    <div className="App">
      <header className="App-header">
        <div className={'btn-row'}>
          <button className={'btn'}>1</button>
          <button className={'btn'}>2</button>
        </div>
      </header>
    </div>
  );
}

export default hot(App);
