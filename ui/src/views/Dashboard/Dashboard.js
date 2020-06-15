import React from 'react';
import { createApolloFetch } from 'apollo-fetch';

const queryTodos = `query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}`;

const addTodo = `mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}`

const apolloFetch = createApolloFetch({
  uri: 'http://localhost:8080/query',
});

function Dashboard(){
  apolloFetch({
    query: queryTodos,
  }).then(res => {
    console.log(res);
  }).catch(err => {
    console.log(err)
  });

  return (
    <div>
      Test data
    </div>
  )
}

export default Dashboard;