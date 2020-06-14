import React from 'react';
import { createApolloFetch } from 'apollo-fetch';

const apolloFetch = createApolloFetch({
  uri: 'http://localhost:8080/graphql',
});

function Dashboard(){
  apolloFetch({
    query: `query{
      vegetable(name: "Corn"){
       name
        price
        image
      }
    } `,
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