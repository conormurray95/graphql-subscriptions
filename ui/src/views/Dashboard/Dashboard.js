import React from 'react';
import { createApolloFetch } from 'apollo-fetch';
import { useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';


const TODOS = gql`query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}`;

const addTodo = gql`mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}`

function Dashboard(){
  const { loading, error, data } = useQuery(TODOS);
  console.log({data, loading, error});

  return (
    <div>
      Test data
    </div>
  )
}

export default Dashboard;