task tracker cli
===============

runnen:
  go run .
  go build -o task-cli.exe

gebruik:
  add "tekst"
  update <id> "tekst"
  delete <id>
  mark-in-progress <id>
  mark-done <id>
  list
  list done / list todo / list in-progress


hoe ik het gemaakt heb
----------------------

Go met standard library. Geen extras.

main.go - os.Args voor commands. check args, strconv.Atoi voor id's.

task.go - Task struct (id, description, status, timestamps). load, do, save. 
[x] done [~] bezig [ ] todo.

storage.go - loadTasks leest json, saveTasks schrijft. file vanzelf bij eerste add.
