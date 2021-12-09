import * as React from "react";

type Props = {
  count: number;
  completedCount: number;
  onClearCompleted: () => void;
  nowShowing: "all" | "active" | "completed";
  showActive: () => void;
  showCompleted: () => void;
  showAll: () => void;
};

export const TodoFooter = (props: Props) => {
  const pluralize = (count: number, word: string) => {
    return count === 1 ? word : word + "s";
  };

  var activeTodoWord = pluralize(props.count, "item");
  var clearButton = null;

  if (props.completedCount > 0) {
    clearButton = (
      <button className="clear-completed" onClick={props.onClearCompleted}>
        Clear completed
      </button>
    );
  }

  var nowShowing = props.nowShowing;
  return (
    <footer className="footer">
      <span className="todo-count">
        <strong>{props.count}</strong> {activeTodoWord} left
      </span>
      <ul className="filters">
        <li>
          <a
            onClick={props.showAll}
            href="#/"
            className={nowShowing === "all" ? "selected" : "not-selected"}
          >
            All
          </a>
        </li>{" "}
        <li>
          <a
            onClick={props.showActive}
            href="#/active"
            className={nowShowing === "active" ? "selected" : "not-selected"}
          >
            Active
          </a>
        </li>{" "}
        <li>
          <a
            onClick={props.showCompleted}
            href="#/completed"
            className={nowShowing === "completed" ? "selected" : "not-selected"}
          >
            Completed
          </a>
        </li>
      </ul>
      {clearButton}
    </footer>
  );
};
