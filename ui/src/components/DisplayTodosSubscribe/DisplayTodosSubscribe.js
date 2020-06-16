import React from 'react';
import {  useQuery, useMutation, useSubscription } from '@apollo/react-hooks';
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

const TODO_SUBSCRIPTION = gql`
  subscription{
    todoAdded{
      id
      text
      done
    }
}
`;

function DisplayTodosSubscribe(){
  let todos = [];
  const { loading, error, data, refetch } = useQuery(TODOS);
  const [updateTodo ] = useMutation(UPDATE_TODO);
  const subscriptionData = useSubscription(
    TODO_SUBSCRIPTION,
  );

  if(!subscriptionData.loading){
    if(subscriptionData.data.todoAdded)
    {
      if(!data.todos.find(todo => todo.id === subscriptionData.data.todoAdded.id)){
        data.todos.push(subscriptionData.data.todoAdded)
      }
    }
  }

  if(!loading){
    todos = data.todos;
  }

  return (
    <div>
      Your todo list (Updated with subscription results)
      {todos.map(todo => (
          <div>
            <input type="checkbox" checked={todo.done} onChange={() => updateTodo({variables: {id: todo.id, done: !todo.done, text: todo.text}})}/>
            {todo.text}
          </div>
      ))}
    </div>
  )
}

export default DisplayTodosSubscribe;