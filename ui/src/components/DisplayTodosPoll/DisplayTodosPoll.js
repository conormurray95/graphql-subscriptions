import React from 'react';
import {  useQuery, useMutation } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';


const TODOS = gql`query findTodos {
  todos {
    text
    done
    id
  }
}`;

const UPDATE_TODO = gql`mutation updateTodo($id: ID!, $done: Boolean!, $text: String!) {
  updateTodo(input:{id: $id, done: $done, text: $text}) {
    id
    text
    done
  }
}`


function DisplayTodosPoll(){
  let todos = [];
  const { loading, error, data, refetch } = useQuery(TODOS);
  const [updateTodo ] = useMutation(UPDATE_TODO);

  setInterval(() => refetch(), 5000)

  if(!loading){
    todos = data.todos;
  }

  return (
    <div>
      Your todo list (Updated every 5 seconds)
      {todos.map(todo => (
          <div>
            <input type="checkbox" checked={todo.done} onChange={() => updateTodo({variables: {id: todo.id, done: !todo.done, text: todo.text}})}/>
            {todo.text}
          </div>
      ))}
    </div>
  )
}

export default DisplayTodosPoll;