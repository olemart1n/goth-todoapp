const filterOutTodoFromDom = (id) => {
  const todoListSection = document.querySelector("#todo-list");
  const todoList = todoListSection.children; // Note: children is a property, not a method, so no parentheses

  // Find the todo element with the matching ID and remove it
  const todoToDelete = Array.from(todoList).find(
    (todo) => todo.dataset.id === id.toString(),
  );

  if (todoToDelete) {
    todoToDelete.remove();
  }
};
