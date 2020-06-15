import React from 'react';
import { createApolloFetch } from 'apollo-fetch';

const apolloFetch = createApolloFetch({
  uri: 'http://localhost:8080/query',
});

function Dashboard(){
  apolloFetch({
    query: `query findTodos {
      todos {
        text
        done
        user {
          name
        }
      }
  }`,
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