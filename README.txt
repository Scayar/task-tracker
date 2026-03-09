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

Gewoon Go met standard library, niks extras.

main.go - hier kijk ik naar os.Args. eerste arg = command, rest zijn parameters. 
check of er genoeg args zijn anders error. id's moet ik van string naar int dus 
strconv.Atoi. als dat faalt is het geen goed id. usage() laat alle commands zien.

task.go - Task struct met id, description, status, createdAt, updatedAt. elke 
functie laadt tasks, doet zn ding, slaat op. bij add zoek ik max id + 1. bij 
update/delete/mark loop ik tot ik de juiste id vind. list filtered op status 
als je dat meegeeft. [x] done [~] bezig [ ] nog te doen.

storage.go - loadTasks leest tasks.json, als file niet bestaat of json klopt 
niet dan lege slice. saveTasks schrijft weer weg met MarshalIndent. bestand 
wordt vanzelf aangemaakt bij eerste add.
