https://roadmap.sh/projects/task-tracker

# Task-Cli
  - This is a project from backend road map website
  - beginner level
  - currently implemented with rewriting the entire file
  - will update in future probably

# Usage Instructions as per project specs in the site

Add task-cli binary to environment path if you want to use at all times

```# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress```
