# Проект по изучению возможностей серверной части Golang #

---

### Информация для разворачивания проекта

В рамках этого проекта будет рализована базовая функциональность фреймворка на `Go`

Для сборки проекта локально достаточно будет в корне проекта `/web` испольнить следующую команду

```makefile
    make init
```

Для сборки проекта запустить в корне проекта `/web`

```makefile
    make init
```

Для запуска проекта запустить в корне проекта `/web`

```makefile
    make up
```

### Возможные проблемы
1. Проверить установлены ли все зависимости из `go.mod`
2. Проверить, что в файла `go.mod` название `module` созвпадает с названием папки, в котором находится проект `module project`
3. Проверить не заняты ли порты, которые указаны в контейнерах в папке `docker`