import './App.css';
import React from 'react';
import Profile from './profile';
import Tabe from './tabe';
import Mascot from './Mascot';

function App() {
  return (
    <div className="App">
      <div className='box'>
        <div className='box_left'>
          <Profile/>
          <Tabe/>
        </div>
      </div>
      <Mascot />
    </div>
  );
}

export default App;
