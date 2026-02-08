package model

type TaskManager string
const(
	StatusPedende TaskManager = "pendente"
	StatusEmAndamento TaskManager = "em andamento"
	StatusFinalizada TaskManager = "finalizada"
)