'use client';

import { useState } from 'react';
import { API } from '../services/api';

export default function Home() {
  const [inputText, setInputText] = useState('');

  const handleTodoPost = async () => {
    await API.post('/todos', inputText);
  };

  return (
    <main className='flex min-h-screen flex-col items-center p-24'>
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
    </main>
  );
}
