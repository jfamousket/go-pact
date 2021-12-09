import * as React from "react";
import ReactDOM from "react-dom";

const ESCAPE_KEY = 27;
const ENTER_KEY = 13;

const isInputElement = (
  node
)=> {
  return node instanceof HTMLInputElement;
};

export const TodoItem = (props) => {
  const [state, setState] = React.useState({
    editText: props.todo.title,
  });
  const ref = React.useRef(null);

  React.useEffect(() => {
    if (!props.editing && props.editing) {
      const node = ReactDOM.findDOMNode(ref.current);
      if (isInputElement(node)) {
        node.focus();
        node.setSelectionRange(node.value.length, node.value.length);
      }
    }
  }, []);

  const handleSubmit = (
    event
  ) => {
    var val = state.editText.trim();
    if (val) {
      props.onSave(props.todo, val);
    } else {
      props.onDestroy();
    }
  };

  const handleEdit = () => {
    props.onEdit(props.todo);
    setState({ editText: props.todo.title });
  };

  const handleKeyDown = (event) => {
    if (event.which === ESCAPE_KEY) {
      setState({ editText: props.todo.title });
      props.onCancel(event);
    } else if (event.which === ENTER_KEY) {
      handleSubmit(event);
    }
  };

  const handleChange = (event) => {
    if (props.editing === props.todo.id) {
      setState({ editText: event.target.value });
    }
  };

  let status =
    props.editing === props.todo.id
      ? "editing"
      : props.todo.completed
      ? "completed"
      : "static";

  return (
    <li className={status}>
      <div className="view">
        <input
          className="toggle"
          type="checkbox"
          checked={props.todo.completed}
          onChange={props.onToggle}
        />
        <label onDoubleClick={handleEdit}>{props.todo.title}</label>
        <button className="destroy" onClick={props.onDestroy} />
      </div>
      <input
        ref={ref}
        className="edit"
        value={state.editText}
        onBlur={handleSubmit}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
      />
    </li>
  );
};
