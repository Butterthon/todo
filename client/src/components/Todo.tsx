import React, { useEffect } from 'react'
import { GetTodo } from 'api/todo';

interface TodoProps {
  items: {
    id: string,
    text: string
  }[]
};

const Todo: React.FC<TodoProps> = props => {
  useEffect(() => {
    const todo = GetTodo();
  })

  return (
    <ul>
      {
        props.items.map(item => (
          <li key={item.id}>{item.text}</li>
        ))
      }
    </ul>
  )
};

export default Todo;
