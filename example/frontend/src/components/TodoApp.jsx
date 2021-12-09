import * as React from "react";
import { v4 } from "uuid";

import { Todo, TodoItem } from "./TodoItem";
import { TodoFooter } from "./Footer";

const ENTER_KEY = 13;

const ALL_TODOS = "all";
const ACTIVE_TODOS = "active";
const COMPLETED_TODOS = "completed";

const reducer = (
  p,
  c
) => ({
  ...p,
  ...(typeof c === "function" ? c(p) : c),
});

export const TodoApp = ({ Pact }) => {
  const [state, setState] = React.useReducer(
    reducer,
    {
      todos: [],
      onChanges: [],
      nowShowing: ALL_TODOS,
      editing: null,
      newTodo: "",
      editText: "",
    },
    (s) => s
  );

  React.useEffect(() => {
    // Pact.genKeyPair().catch((err) => console.log("genKeyPair error", err));
  }, []);

  const getTodos = () => {
    return Pact.getTodos()
      .then((res) => res.result.data)
      .then((todos) => {
        const notDeleted = todos
          .filter((todo) => {
            return todo.deleted === false;
          })
          .sort((a, b) => {
            if (a.title < b.title) return -1;
            if (b.title > a.title) return 1;
            return 0;
          });
        setState({ todos: notDeleted });
      });
  };

  React.useEffect(() => {
    getTodos();
  }, []);

  const add = (title) => {
    const uuid = v4();
    Pact.add(title, uuid)
      .then(() => getTodos())
      .catch((err) => console.log("send add err", err));
  };

  const toggle = (todo) => {
    Pact.toggle(todo.id)
      .then(() => getTodos())
      .catch((err) => console.log("toggle err", err));
  };

  const toggleAll = async () => {
    const activeTodos = state.todos.filter((todo) => !todo.completed);
    const completedTodos = state.todos.filter((todo) => todo.completed);
    const toggleTodos =
      state.todos.length === completedTodos.length
        ? completedTodos
        : activeTodos;
    const ids = toggleTodos.map((todo) => todo.id);
    Pact.toggleAll(ids)
      .then(() => getTodos())
      .catch((err) => console.log("toggleAll err", err));
  };

  const destroy = (todo) => {
    if (typeof todo === "undefined") return;
    Pact.destroy(todo.id)
      .then(() => getTodos())
      .catch((err) => console.log("destroy err", err));
  };

  const clearCompleted = async () => {
    const completedTodos = state.todos.filter((todo) => todo.completed);
    const ids = completedTodos.map((todo) => todo.id);
    Pact.clearCompleted(ids)
      .then(() => getTodos())
      .catch((err) => console.log("clear completed err", err));
  };

  const save = (todo, text) => {
    Pact.save(todo.id, text ?? "")
      .then(() => getTodos())
      .catch((err) => console.log("save err", err));

    setState({ editing: null });
  };

  const edit = (todo) => {
    setState({ editing: todo.id });
  };

  const cancel = () => {
    setState({ editing: null });
  };

  const showActive = () => {
    setState({ nowShowing: ACTIVE_TODOS });
  };

  const showCompleted = () => {
    setState({ nowShowing: COMPLETED_TODOS });
  };

  const showAll = () => {
    setState({ nowShowing: ALL_TODOS });
  };

  const handleChange = (event) => {
    setState({ newTodo: event.target.value });
  };

  const handleNewTodoKeyDown = (
    event
  ) => {
    if (event.keyCode !== ENTER_KEY) {
      return;
    }

    event.preventDefault();

    var val = state.newTodo.trim();

    if (val) {
      add(val);
      setState({ newTodo: "" });
    }
  };

  var footer;
  var main;
  var todos = state.todos;
  var shownTodos = todos.filter(function (todo) {
    switch (state.nowShowing) {
      case ACTIVE_TODOS:
        return !todo.completed;
      case COMPLETED_TODOS:
        return todo.completed;
      default:
        return true;
    }
  }, this);

  var todoItems = shownTodos.map(function (todo) {
    return (
      <TodoItem
        key={todo.id}
        todo={todo}
        onToggle={() => toggle(todo)}
        onDestroy={() => destroy(todo)}
        onEdit={edit}
        onSave={save}
        editing={state.editing}
        onCancel={cancel}
      />
    );
  }, this);

  var activeTodoCount = todos.reduce(function (accum, todo) {
    return todo.completed ? accum : accum + 1;
  }, 0);

  var completedCount = todos.length - activeTodoCount;

  if (activeTodoCount || completedCount) {
    footer = (
      <TodoFooter
        count={activeTodoCount}
        completedCount={completedCount}
        nowShowing={state.nowShowing}
        onClearCompleted={clearCompleted}
        showActive={showActive}
        showCompleted={showCompleted}
        showAll={showAll}
      />
    );
  }

  if (todos.length) {
    main = (
      <section className="main">
        <input
          id="toggle-all"
          className="toggle-all"
          type="checkbox"
          onChange={toggleAll}
          checked={activeTodoCount === 0}
        />
        <label htmlFor="toggle-all" />
        <ul className="todo-list">{todoItems}</ul>
      </section>
    );
  }

  return (
    <div>
      <header className="header">
        <h1>todos</h1>
        <input
          className="new-todo"
          placeholder="What needs to be done?"
          value={state.newTodo}
          onKeyDown={handleNewTodoKeyDown}
          onChange={handleChange}
          autoFocus={true}
        />
      </header>
      {main}
      {footer}
    </div>
  );
};
