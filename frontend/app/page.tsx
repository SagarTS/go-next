'use client';

import { useEffect, useState } from 'react';
import { API } from '../services/api';

interface todoData {
  todo: string;
}

export default function Home() {
  const [inputText, setInputText] = useState('');
  const [data, setData] = useState<todoData[] | null>(null);
  const [loading, setLoading] = useState(false);

  const handleTodoPost = async () => {
    await API.post('/todos', { todo: inputText });
    setInputText('');
    fetchData();
  };

  async function fetchData() {
    setLoading(true);
    try {
      const response = await fetch('http://localhost:8000/todos');
      const jsonData = await response.json();
      setData(jsonData);
    } catch (err) {
      console.error('Error', err);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <main className='flex min-h-screen flex-col items-center p-24'>
      {loading ? (
        <p>LOADING...</p>
      ) : (
        <>
          <h1>Todo</h1>
          <div className='flex gap-5 p-5'>
            <input
              onChange={(e) => {
                setInputText(e.target.value);
              }}
              placeholder='Todo Task'
            />
            <button onClick={handleTodoPost}>Tick</button>
          </div>

          <div>
            {data &&
              data?.map((d, i) => (
                <div key={i}>
                  <p>{d.todo}</p>
                </div>
              ))}
          </div>
        </>
      )}
    </main>
  );
}
