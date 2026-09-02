# Tutorial de Git: lo más importante para controlar versiones

## 1. ¿Qué es Git?

Git es un sistema de control de versiones distribuido.

Su función principal es llevar registro de los cambios que haces en archivos a lo largo del tiempo.

Esto permite:

- guardar versiones del proyecto
- volver a estados anteriores
- trabajar con varias personas al mismo tiempo
- mantener ramas de desarrollo
- evitar perder trabajo

Git no es lo mismo que GitHub:

- Git = sistema de control de versiones local
- GitHub = plataforma en la nube para alojar repositorios Git

---

## 2. ¿Para qué sirve Git?

Git sirve para:

- controlar cambios en código
- mantener historial de commits
- trabajar en equipo
- crear ramas para features o correcciones
- fusionar cambios sin romper el proyecto
- volver a una versión anterior cuando algo sale mal

Es muy útil para:

- proyectos de software
- documentación
- configuración
- cualquier trabajo donde quieras llevar historial de cambios

---

## 3. ¿Qué es un repositorio?

Un repositorio es un proyecto Git.

Puede estar en tu computadora o en un servidor remoto.

Hay dos tipos principales:

- repositorio local
- repositorio remoto

### Repositorio local

Es el que está en tu máquina.

```bash
git init
```

### Repositorio remoto

Es el que está en GitHub, GitLab o Bitbucket.

```bash
git clone <url>
```

---

## 4. ¿Qué es un commit?

Un commit es un punto de guardado del proyecto.

Cada commit guarda:

- los archivos modificados
- un mensaje descriptivo
- una referencia temporal en el historial

### Ejemplo

```bash
git add .
git commit -m "agrega validación de login"
```

Ese mensaje explica qué cambio hiciste.

### Regla útil

Los commits deben ser pequeños y claros.

No conviene hacer un commit gigante sin describir bien el cambio.

---

## 5. Estado del repositorio

Git tiene un flujo básico de estado:

- working directory: archivos en tu carpeta actual
- staging area: archivos preparados para commit
- repository: commits guardados en Git

### Ver el estado

```bash
git status
```

Esto te muestra:

- archivos modificados
- archivos agregados
- archivos en staging
- ramas actuales

---

## 6. Agregar archivos al commit

```bash
git add nombre_archivo
```

O todos los cambios:

```bash
git add .
```

### Importante

`git add` no guarda todavía el cambio. Solo lo prepara para commit.

---

## 7. Hacer un commit

```bash
git commit -m "mensaje del cambio"
```

Ejemplo:

```bash
git commit -m "corrige error en login"
```

### Buenas prácticas

- usa mensajes claros
- no pongas mensajes vagos como “fix”
- intenta ser descriptivo

---

## 8. Historial de commits

Para ver qué commits se han hecho:

```bash
git log
```

También puedes ver una versión más corta:

```bash
git log --oneline
```

Esto te muestra:

- hash del commit
- autor
- fecha
- mensaje

---

## 9. ¿Qué es una rama?

Una rama es una línea de desarrollo separada del proyecto principal.

Por ejemplo:

- `main` o `master` = rama principal
- `feature/login` = rama para una nueva funcionalidad
- `fix/error` = rama para corregir un bug

### Ver ramas

```bash
git branch
```

### Crear rama

```bash
git checkout -b feature/login
```

O una forma moderna:

```bash
git switch -c feature/login
```

---

## 10. Cambiar de rama

```bash
git switch main
```

O en versiones antiguas:

```bash
git checkout main
```

---

## 11. ¿Qué es merge?

Merge significa fusionar cambios de una rama a otra.

Ejemplo:

```bash
git switch main
git merge feature/login
```

Esto incorpora los cambios de la rama `feature/login` dentro de `main`.

### Cuando conviene usar merge

- cuando querés incorporar una funcionalidad terminada
- cuando la rama ya está lista

---

## 12. ¿Qué es rebase?

Rebase también mueve o reaplica cambios de una rama sobre otra.

Se usa para dejar un historial más limpio.

Ejemplo:

```bash
git rebase main
```

### Diferencia simple

- merge = fusiona ramas conservando historial
- rebase = reescribe el historial de la rama para que quede lineal

### Cuidado

Rebase puede ser más complicado si la rama ya fue compartida con otras personas.

---

## 13. ¿Qué es un remoto?

Un remoto es una referencia a un repositorio en la nube o en otra máquina.

Ejemplo:

```bash
git remote -v
```

Esto muestra las URLs remotas.

### Agregar remoto

```bash
git remote add origin <url-del-repo>
```

### Cambiar remoto

```bash
git remote set-url origin <nueva-url>
```

---

## 14. ¿Qué es git push?

`git push` sube tus cambios al repositorio remoto.

```bash
git push origin main
```

Ejemplo:

```bash
git push -u origin feature/login
```

- `-u` = establece la rama remota por defecto para ese branch

---

## 15. ¿Qué es git pull?

`git pull` baja los cambios del repositorio remoto y los actualiza en tu rama local.

```bash
git pull origin main
```

### Importante

Si tienes cambios locales sin guardar, `git pull` puede dar conflictos.

---

## 16. ¿Qué es git clone?

Clona un repositorio remoto a tu máquina local.

```bash
git clone <url>
```

Ejemplo:

```bash
git clone git@github.com:usuario/proyecto.git
```

Esto crea una copia local del repositorio.

---

## 17. ¿Qué es git fetch?

`git fetch` descarga cambios del remoto sin mezclar nada con tu branch actual.

```bash
git fetch origin
```

Luego puedes revisar diferencias sin alterar tu trabajo actual.

---

## 18. ¿Qué es git diff?

Muestra las diferencias entre versiones.

```bash
git diff
```

También puedes comparar un archivo o una rama específica:

```bash
git diff archivo.txt
```

```bash
git diff main..feature/login
```

---

## 19. ¿Qué es git restore?

Sirve para deshacer cambios en archivos.

### Descartar cambios locales de un archivo

```bash
git restore archivo.txt
```

### Deshacer cambios en staging

```bash
git restore --staged archivo.txt
```

Esto es útil si te equivocaste y no quieres guardar ese cambio.

---

## 20. ¿Qué es git reset?

`git reset` sirve para mover el puntero HEAD a otro commit.

Ejemplo:

```bash
git reset --soft HEAD~1
```

Esto hace que el último commit vuelva al área de staging.

Otro ejemplo:

```bash
git reset --hard HEAD~1
```

### Cuidado

`--hard` elimina cambios del historial y del working directory.

Se usa con cuidado.

---

## 21. ¿Qué es git revert?

`git revert` crea un commit nuevo que deshace un cambio anterior.

```bash
git revert <hash-del-commit>
```

Es más seguro que `reset`, porque no borra historial, sino que crea un commit de corrección.

---

## 22. ¿Qué es .gitignore?

Es un archivo que indica a Git qué archivos ignorar.

Ejemplo:

```gitignore
node_modules/
.env
*.log
```

Esto evita subir cosas que no deberían ir al repositorio, como:

- archivos de entorno
- dependencias locales
- logs
- archivos temporales

---

## 23. ¿Qué es un conflicto de merge?

Un conflicto aparece cuando dos personas tocaron la misma parte de un archivo y Git no sabe cuál versión conservar.

Ejemplo:

```bash
git merge feature/login
```

Si se genera conflicto, Git marca el archivo con marcadores:

```text
<<<<<<< HEAD
...
=======
...
>>>>>>> feature/login
```

Tú debes resolverlo manualmente y luego:

```bash
git add archivo
git commit
```

---

## 24. Flujo de trabajo recomendado

### Flujo simple

1. crear rama
2. hacer cambios
3. hacer commit
4. hacer push
5. abrir pull request
6. merge a main

Ejemplo:

```bash
git switch -c feature/nueva-funcion
# haces cambios
git add .
git commit -m "agrega nueva función"
git push -u origin feature/nueva-funcion
```

---

## 25. Git en equipo

Cuando varios desarrolladores trabajan en el mismo proyecto:

- cada uno trabaja en su rama
- suben sus cambios al remoto
- se revisan los cambios
- se fusionan a la rama principal

Esto reduce conflicto y facilita organizar el trabajo.

---

## 26. ¿Qué es un pull request o merge request?

Es una solicitud para revisar cambios antes de fusionarlos a la rama principal.

Los equipos usan esto para:

- revisar código
- detectar errores
- discutir cambios
- aprobar cambios antes del merge

---

## 27. ¿Qué es GitHub?

GitHub es una plataforma web para alojar repositorios Git.

Permite:

- compartir código
- colaborar con otros
- revisar PRs
- documentar proyectos
- deploys y CI/CD

---

## 28. Comandos básicos más usados

```bash
git init
```

```bash
git clone <url>
```

```bash
git status
```

```bash
git add .
```

```bash
git commit -m "mensaje"
```

```bash
git push origin main
```

```bash
git pull origin main
```

```bash
git branch
```

```bash
git switch -c nombre-rama
```

```bash
git merge nombre-rama
```

```bash
git log --oneline
```

```bash
git diff
```

---

## 29. Recomendaciones prácticas

- haz commits pequeños y descriptivos
- usa ramas para cada tarea
- no subas archivos sensibles
- revisa `git status` antes de hacer commit
- usa `.gitignore` para ignorar archivos innecesarios
- no reescribas historial de forma casual si ya compartiste tu rama

---

## 30. Resumen rápido

Git es una herramienta que permite:

- guardar versiones del proyecto
- deshacer errores
- colaborar con otros
- mantener ramas de trabajo
- controlar cambios de forma ordenada

Los conceptos clave son:

- repositorio
- commit
- rama
- push
- pull
- merge
- conflict
- .gitignore

---

## 31. Flujo de trabajo típico

```bash
git switch -c feature/cambio
# editas archivos
# pruebas tu código

git add .
git commit -m "agrega cambio"
git push -u origin feature/cambio
```

Luego en GitHub:

- crea pull request
- revisa el cambio
- aprueba
- merge a main

---

## 32. Regla mental final

Git no solo guarda archivos. Guarda la historia del proyecto.

Eso es lo que permite:

- volver en el tiempo
- comparar versiones
- trabajar en equipo
- mantener el proyecto ordenado

Y cuando lo combinas con SSH y llaves, ya tenés el flujo real de trabajo profesional con GitHub o GitLab.
