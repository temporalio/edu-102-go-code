package translation

const TaskQueueName = "durable-exec-tasks"

type TranslationWorkflowInput struct {
	Name         string
	LanguageCode string
}

type TranslationWorkflowOutput struct {
	HelloMessage   string
	GoodbyeMessage string
}

type TranslationActivityInput struct {
	Term         string
	LanguageCode string
}

type TranslationActivityOutput struct {
	Translation string
}
