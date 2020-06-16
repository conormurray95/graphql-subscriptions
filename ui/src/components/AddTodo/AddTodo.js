import React from 'react';
import {  useMutation } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';


const ADD_TODO = gql`mutation createTodo($text: String!) {
  createTodo(input:{text:$text}) {
    id
    text
    done
  }
}`

function AddTodo(){
  let input;
  const [addTodo] = useMutation(ADD_TODO);

  return (
    <div style={{paddingBottom:'50px'}}>
      
      <form onSubmit={e => {
        e.preventDefault();
        addTodo({variables: {text: input.value}});
        input.value = '';
      }}>
        <input ref={node => {input = node}}/>
        <button type="submit">Add Todo</button>
      </form>
    </div>
  )
}

export default AddTodo;