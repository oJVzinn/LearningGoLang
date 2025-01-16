package tasks

import (
	"fmt"
	"strings"
	"time"
)

func TaskSetup() error {
	continueLoop := true
	for continueLoop {
		fmt.Println(getMenu())
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			return err
		}

		switch option {
		case 0:
			{
				fmt.Println("Digite o nome da tarefa: ")
				var taskName string
				_, err := fmt.Scanln(&taskName)
				if err != nil {
					return err
				}

				_, find := FindByName(taskName)
				if find {
					fmt.Println("Essa tarefa já existe!")
					break
				}

				AddNewTask(Task{Name: taskName, InitialDate: time.Now().Format("02/01/2006 às 15:04:05")})
				fmt.Println("Tarefa salva com sucesso!")
				break
			}

		case 1:
			{
				fmt.Println("Digite o nome da tarefa: ")
				var taskName string
				_, err := fmt.Scanln(&taskName)
				if err != nil {
					return err
				}

				task, find := FindByName(taskName)
				if !find {
					fmt.Println("Tarefa não encontrada!")
					break
				}

				RemoveTask(*task)
				fmt.Println("Tarefa removida com sucesso!")
				break
			}

		case 2:
			{
				fmt.Println("Digite o nome da tarefa: ")
				var taskName string
				_, err := fmt.Scanln(&taskName)
				if err != nil {
					return err
				}

				task, find := FindByName(taskName)
				if !find {
					fmt.Println("Tarefa não encontrada!")
					break
				}

				UpdateTask(*task)
				fmt.Println("Tarefa atualizada com sucesso!")
				break
			}

		case 3:
			{
				builder := strings.Builder{}
				builder.WriteString("Nenhuma tarefa registada!")
				if len(tasks) > 0 {
					builder.Reset()
					for _, task := range tasks {
						completedMessage := func() string {
							if task.Completed {
								return "Sim"
							}

							return "Não"
						}

						builder.WriteString(task.Name + " iniciou em: " + task.InitialDate + ". Concluida: " + completedMessage())
						builder.WriteString("\n")
					}
				}

				fmt.Println("\n", builder.String())
				break
			}

		case 4:
			{
				fmt.Println("Encerrando...")
				continueLoop = false
				break
			}

		default:
			{
				fmt.Println("Opção não encontrada!")
				break
			}
		}
	}

	return nil
}

func getMenu() string {
	return "\nDigite qual opção você quer\nCriar tarefa = 0\nDeletar tarefa = 1\nConcluir tarega = 2\nListar tarefas = 3\nSair = 4\n"
}
