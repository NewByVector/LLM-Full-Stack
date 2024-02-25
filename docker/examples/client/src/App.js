import './App.css';
import { useEffect, useState } from 'react';

async function timeout(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

function App() {
  const [input, setInput] = useState('');
  const [allValues, setAllValues] = useState('');
  const [currentValue, setCurrentValue] = useState('');

  const getAll = async () => {
    const response = await fetch('/api/values/all');
    const data = await response.json();
    setAllValues(JSON.stringify(data));
  }

  const getCurrent = async () => {
    const response = await fetch('/api/values/current');
    const data = await response.json();
    setCurrentValue(JSON.stringify(data));
  }

  useEffect(() => {
    getAll();
    getCurrent();
  }, []);


  const onSubmit = async () => {
    await fetch('/api/values', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ index: input })
    });
    await timeout(1000);
    getAll();
    getCurrent();
  }

  return (
    <div className="App">
      <input value={input} onChange={e => setInput(e.target.value)} />
      <button onClick={onSubmit}>Submit</button>
      <code>
        {allValues}
      </code>
      <code>
        {currentValue}
      </code>
    </div>
  );
}

export default App;
