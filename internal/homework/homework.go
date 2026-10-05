package homework

import "fmt"

// BuildGreeting возвращает приветствие для ученика.
//
// Ожидаемый формат:
// "Привет, {name}! Добро пожаловать в Go."
func BuildGreeting(name string) string {
	greeting := fmt.Sprintf("Привет, %s! Добро пожаловать в Go.", name)
	return greeting
}

// BuildCourseWelcome возвращает название курса.
//
// Ожидаемый формат:f x
// "Курс: {courseName}"
func BuildCourseWelcome(courseName string) string {
	course := fmt.Sprintf("Курс: %s", courseName)
	return course
}

// BuildLessonTitle возвращает название первого урока.
//
// Ожидаемый формат:
// "Урок 1: {lessonName}"
func BuildLessonTitle(lessonName string) string {
	lessonTitle := fmt.Sprintf("Урок 1: %s", lessonName)
	return lessonTitle
}

// BuildRepositoryPath возвращает путь до репозитория на GitHub.
//
// Ожидаемый формат:
// "github.com/{owner}/{repo}"
func BuildRepositoryPath(owner string, repo string) string {
	repositoryPath := fmt.Sprintf("github.com/%s/%s", owner, repo)
	return repositoryPath
}

// BuildRunCommand возвращает команду запуска Go-программы.
//
// Ожидаемый формат:
// "go run {packagePath}"
func BuildRunCommand(packagePath string) string {
	runCommand := fmt.Sprintf("go run %s", packagePath)
	return runCommand
}
