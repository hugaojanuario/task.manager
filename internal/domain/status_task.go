package domain

type TaskManager string

const (
	StatusPedende     TaskManager = "pendente"
	StatusEmAndamento TaskManager = "em andamento"
	StatusFinalizada  TaskManager = "finalizada"
)
