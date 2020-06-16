import React from 'react';
import AddTodo from '../../components/AddTodo';
import DisplayTodosInitialLoad from '../../components/DisplayTodosInitialLoad';
import DisplayTodosPoll from '../../components/DisplayTodosPoll';
import DisplayTodosNotification from '../../components/DisplayTodosNotification';
import DisplayTodosSubscribe from '../../components/DisplayTodosSubscribe';


function Dashboard(){

  return (
    <div>
      <AddTodo/>
      <DisplayTodosSubscribe/>
    </div>
  )
}

export default Dashboard;