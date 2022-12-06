import React, { useState, useEffect } from "react";
import axios from "axios";
import "./App.css";
import TodoCard from "./Atoms/ItemCard";
import TodoForm from "./Atoms/Form";
import Footer from "./Atoms/Footer";

axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";

const getTodos = (endpoint) => {
  endpoint = process.env.REACT_APP_API_ENDPOINT + endpoint;
  return axios.get(endpoint);
};

const App = () => {
  const [todos, setTodos] = useState();

  useEffect(() => {
    getTodos("/todos").then(function (response) {
      setTodos(response.data);
    });
  }, []);

  console.log("[+] End point:", process.env.REACT_APP_API_ENDPOINT);

  return (
    <div className="App">
      <header className="App-header">
        <div>
          <TodoForm></TodoForm>
        </div>
        <div>
          {todos ? (
            todos.length !== 0 ? (
              todos.map((todo, index) => <TodoCard data={todo} key={index} />)
            ) : (
              <div>Not found</div>
            )
          ) : (
            <div>Not found</div>
          )}
        </div>
      </header>
      <footer>
        <div>
          <Footer />
        </div>
      </footer>
    </div>
  );
};

export default App;
